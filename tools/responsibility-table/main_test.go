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
