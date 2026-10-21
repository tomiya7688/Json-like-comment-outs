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
