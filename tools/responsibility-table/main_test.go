package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunRejectsOutputHardLinkedToGoSource(t *testing.T) {
	sourceRoot := t.TempDir()
	sourcePath := filepath.Join(sourceRoot, "sample.go")
	source := []byte("package sample\n// {\"Responsibility\":{\"Sample\":\"Handles input\"}}\n")
	if err := os.WriteFile(sourcePath, source, 0o644); err != nil {
		t.Fatal(err)
	}

	outputPath := filepath.Join(sourceRoot, "responsibility-table.md")
	if err := os.Link(sourcePath, outputPath); err != nil {
		t.Skipf("hard links unavailable: %v", err)
	}

	var stdout, stderr bytes.Buffer
	if code := run([]string{"--source", sourceRoot, "--output", outputPath}, &stdout, &stderr); code != 2 {
		t.Fatalf("run() returned %d, want 2; stderr: %s", code, stderr.String())
	}
	if !strings.Contains(stderr.String(), "must not overwrite a Go source file") {
		t.Fatalf("missing overwrite diagnostic: %s", stderr.String())
	}
	actual, err := os.ReadFile(sourcePath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(actual, source) {
		t.Fatal("Go source file was changed")
	}
}

func TestExtractResponsibilityRejectsNullValue(t *testing.T) {
	lookup, err := buildSemanticKeyLookup(defaultSemanticKeys())
	if err != nil {
		t.Fatal(err)
	}
	rows, candidate, err := extractResponsibility(`{"Responsibility":{"UserService":null}}`, lookup)
	if !candidate || err == nil || !strings.Contains(err.Error(), "must be strings") {
		t.Fatalf("null responsibility: rows=%v, candidate=%v, err=%v", rows, candidate, err)
	}
	if len(rows) != 0 {
		t.Fatalf("unexpected rows for null responsibility: %v", rows)
	}
}

func TestDecodeSemanticKeysRejectsNullMapping(t *testing.T) {
	_, err := decodeSemanticKeys([]byte(`{"SemanticKeys":null}`))
	if err == nil || !strings.Contains(err.Error(), "SemanticKeys must map") {
		t.Fatalf("expected SemanticKeys mapping error, got %v", err)
	}
}

func TestExtractResponsibilityMalformedCommentCandidate(t *testing.T) {
	lookup, err := buildSemanticKeyLookup(defaultSemanticKeys())
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name      string
		comment   string
		candidate bool
	}{
		{name: "unknown substring", comment: `{"NotResponsibility":}`, candidate: false},
		{name: "nested key", comment: `{"Unknown":{"Responsibility":}}`, candidate: false},
		{name: "known key", comment: `{"Responsibility":}`, candidate: true},
		{name: "known alias", comment: `{"役割":}`, candidate: true},
		{name: "escaped key", comment: `{"Respons\u0069bility":}`, candidate: true},
		{name: "braces in value", comment: `{"Unknown":"{","Responsibility":}`, candidate: true},
		{name: "value is not key", comment: `{"Unknown":"Responsibility":}`, candidate: false},
		{name: "array item is not key", comment: `{"Unknown":[0,"Responsibility":]}`, candidate: false},
		{name: "quoted substring", comment: `{"Not\"Responsibility":}`, candidate: false},
		{name: "space before colon", comment: `{"Responsibility" :}`, candidate: true},
		{name: "different case", comment: `{"RESPONSIBILITY":}`, candidate: true},
		{name: "no object start", comment: `"Responsibility":}`, candidate: false},
		{name: "extra closing brace", comment: `{} } {"Unknown":{"Responsibility":}}`, candidate: false},
		{name: "second object", comment: `{"Unknown":1} {"Responsibility":}`, candidate: false},
		{name: "nested then known", comment: `{"Unknown":{"X":1},"Responsibility":}`, candidate: true},
		{name: "array then known", comment: `{"Unknown":[1,2],"Responsibility":}`, candidate: true},
		{name: "missing comma", comment: `{"Unknown":1 "Responsibility":}`, candidate: false},
		{name: "array before key", comment: `{[]"Responsibility":}`, candidate: false},
		{name: "number before key", comment: `{0"Responsibility":}`, candidate: false},
		{name: "array after comma", comment: `{"Unknown":1,[]"Responsibility":}`, candidate: false},
		{name: "unquoted key", comment: `{Responsibility:}`, candidate: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rows, candidate, err := extractResponsibility(test.comment, lookup)
			if candidate != test.candidate || (err != nil) != test.candidate || len(rows) != 0 {
				t.Fatalf("rows=%v candidate=%v err=%v, want candidate=%v", rows, candidate, err, test.candidate)
			}
		})
	}
}

func TestRunRejectsSymlinkOutputWithoutChangingTarget(t *testing.T) {
	sourceRoot := t.TempDir()
	target := filepath.Join(t.TempDir(), "notes.md")
	before := []byte("keep this file\n")
	if err := os.WriteFile(target, before, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(target, filepath.Join(sourceRoot, "responsibility-table.md")); err != nil {
		t.Skipf("symbolic links unavailable: %v", err)
	}
	var stdout, stderr bytes.Buffer
	if code := run([]string{"--source", sourceRoot}, &stdout, &stderr); code != 2 {
		t.Fatalf("run() returned %d, want 2; stderr: %s", code, stderr.String())
	}
	actual, err := os.ReadFile(target)
	if err != nil || !bytes.Equal(actual, before) {
		t.Fatalf("symlink target changed: %q, %v", actual, err)
	}
	if err := os.Remove(filepath.Join(sourceRoot, "responsibility-table.md")); err != nil {
		t.Fatal(err)
	}
	missingTarget := filepath.Join(t.TempDir(), "missing.md")
	if err := os.Symlink(missingTarget, filepath.Join(sourceRoot, "responsibility-table.md")); err != nil {
		t.Fatal(err)
	}
	stdout.Reset()
	stderr.Reset()
	if code := run([]string{"--source", sourceRoot}, &stdout, &stderr); code != 2 {
		t.Fatalf("dangling symlink: run() returned %d, want 2; stderr: %s", code, stderr.String())
	}
	if _, err := os.Lstat(missingTarget); !os.IsNotExist(err) {
		t.Fatalf("dangling symlink target was created: %v", err)
	}
}

func TestRunRejectsOutputOverSemanticKeys(t *testing.T) {
	sourceRoot := t.TempDir()
	configPath := filepath.Join(sourceRoot, "semantic-keys.json")
	before := []byte(`{"SemanticKeys":{"Responsibility":[]}}`)
	if err := os.WriteFile(configPath, before, 0o644); err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer
	if code := run([]string{"--source", sourceRoot, "--semantic-keys", configPath, "--output", configPath}, &stdout, &stderr); code != 2 {
		t.Fatalf("run() returned %d, want 2; stderr: %s", code, stderr.String())
	}
	actual, err := os.ReadFile(configPath)
	if err != nil || !bytes.Equal(actual, before) {
		t.Fatalf("SemanticKeys configuration changed: %q, %v", actual, err)
	}
	defaultDir := filepath.Join(sourceRoot, ".json-like-comment-outs")
	if err := os.MkdirAll(defaultDir, 0o755); err != nil {
		t.Fatal(err)
	}
	defaultPath := filepath.Join(defaultDir, "semantic-keys.json")
	if err := os.WriteFile(defaultPath, before, 0o644); err != nil {
		t.Fatal(err)
	}
	stdout.Reset()
	stderr.Reset()
	if code := run([]string{"--source", sourceRoot, "--output", defaultPath}, &stdout, &stderr); code != 2 {
		t.Fatalf("default SemanticKeys: run() returned %d, want 2; stderr: %s", code, stderr.String())
	}
	actual, err = os.ReadFile(defaultPath)
	if err != nil || !bytes.Equal(actual, before) {
		t.Fatalf("default SemanticKeys configuration changed: %q, %v", actual, err)
	}
}

func TestRunReplacesExistingTable(t *testing.T) {
	sourceRoot := t.TempDir()
	outputPath := filepath.Join(sourceRoot, "responsibility-table.md")
	if err := os.WriteFile(outputPath, []byte("old table\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer
	if code := run([]string{"--source", sourceRoot}, &stdout, &stderr); code != 0 {
		t.Fatalf("run() returned %d, want 0; stderr: %s", code, stderr.String())
	}
	actual, err := os.ReadFile(outputPath)
	if err != nil || bytes.Contains(actual, []byte("old table")) || !bytes.Contains(actual, []byte("名前")) {
		t.Fatalf("table was not replaced: %q, %v", actual, err)
	}
}

func TestRunPreservesReadOnlyExistingTable(t *testing.T) {
	sourceRoot := t.TempDir()
	outputPath := filepath.Join(sourceRoot, "responsibility-table.md")
	before := []byte("old table\n")
	if err := os.WriteFile(outputPath, before, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(outputPath, 0o444); err != nil {
		t.Skipf("read-only files unavailable: %v", err)
	}
	defer os.Chmod(outputPath, 0o644)
	probe, err := os.OpenFile(outputPath, os.O_WRONLY, 0)
	if err == nil {
		probe.Close()
		t.Skip("write-only open is permitted for this user")
	}
	var stdout, stderr bytes.Buffer
	if code := run([]string{"--source", sourceRoot}, &stdout, &stderr); code != 1 {
		t.Fatalf("run() returned %d, want 1; stderr: %s", code, stderr.String())
	}
	actual, err := os.ReadFile(outputPath)
	if err != nil || !bytes.Equal(actual, before) {
		t.Fatalf("read-only table changed: %q, %v", actual, err)
	}
}
