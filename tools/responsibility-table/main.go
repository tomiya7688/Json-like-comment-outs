package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type tableRow struct {
	source         string
	name           string
	responsibility string
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("responsibility-table", flag.ContinueOnError)
	flags.SetOutput(stderr)
	sourceArg := flags.String("source", "", "Go project directory to scan")
	outputArg := flags.String("output", "", "Markdown output path")
	semanticKeysArg := flags.String("semantic-keys", "", "SemanticKeys JSON configuration path")
	localeArg := flags.String("locale", "ja", "table header locale: ja or en")
	flags.Usage = func() {
		fmt.Fprintln(stderr, "Usage: responsibility-table --source <dir> [--output <file>] [--semantic-keys <file>] [--locale ja|en]")
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if flags.NArg() != 0 {
		fmt.Fprintln(stderr, "responsibility-table: unexpected positional arguments")
		flags.Usage()
		return 2
	}
	if strings.TrimSpace(*sourceArg) == "" {
		fmt.Fprintln(stderr, "responsibility-table: --source is required")
		flags.Usage()
		return 2
	}
	if *localeArg != "ja" && *localeArg != "en" {
		fmt.Fprintln(stderr, "responsibility-table: --locale must be ja or en")
		return 2
	}

	sourceRoot, err := filepath.Abs(*sourceArg)
	if err != nil {
		fmt.Fprintf(stderr, "responsibility-table: resolve source path: %v\n", err)
		return 1
	}
	sourceInfo, err := os.Lstat(sourceRoot)
	if err != nil {
		fmt.Fprintf(stderr, "responsibility-table: inspect source path: %v\n", err)
		return 1
	}
	if sourceInfo.Mode()&os.ModeSymlink != 0 || !sourceInfo.IsDir() {
		fmt.Fprintln(stderr, "responsibility-table: --source must be a directory and must not be a symbolic link")
		return 2
	}

	paths, err := collectGoFiles(sourceRoot)
	if err != nil {
		fmt.Fprintf(stderr, "responsibility-table: scan source directory: %v\n", err)
		return 1
	}
	semanticLookup, err := loadSemanticKeyLookup(sourceRoot, *semanticKeysArg)
	if err != nil {
		fmt.Fprintf(stderr, "responsibility-table: load SemanticKeys: %v\n", err)
		return 1
	}

	rows := make([]tableRow, 0)
	for _, path := range paths {
		relativePath, err := filepath.Rel(sourceRoot, path)
		if err != nil {
			fmt.Fprintf(stderr, "responsibility-table: resolve relative path for %s: %v\n", path, err)
			return 1
		}
		fileRows, diagnostics, err := readResponsibilityRows(path, filepath.ToSlash(relativePath), semanticLookup)
		for _, diagnostic := range diagnostics {
			fmt.Fprintf(stderr, "responsibility-table: warning: %s\n", diagnostic)
		}
		if err != nil {
			fmt.Fprintf(stderr, "responsibility-table: read %s: %v\n", path, err)
			return 1
		}
		rows = append(rows, fileRows...)
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].source != rows[j].source {
			return rows[i].source < rows[j].source
		}
		if rows[i].name != rows[j].name {
			return rows[i].name < rows[j].name
		}
		return rows[i].responsibility < rows[j].responsibility
	})

	outputPath := *outputArg
	if outputPath == "" {
		outputPath = filepath.Join(sourceRoot, "responsibility-table.md")
	} else {
		outputPath, err = filepath.Abs(outputPath)
		if err != nil {
			fmt.Fprintf(stderr, "responsibility-table: resolve output path: %v\n", err)
			return 1
		}
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		fmt.Fprintf(stderr, "responsibility-table: create output directory: %v\n", err)
		return 1
	}
	output, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		fmt.Fprintf(stderr, "responsibility-table: write output: %v\n", err)
		return 1
	}
	defer output.Close()
	outputInfo, err := output.Stat()
	if err != nil {
		fmt.Fprintf(stderr, "responsibility-table: inspect output path: %v\n", err)
		return 1
	}
	overwritesSource, err := outputOverwritesGoFile(paths, outputInfo)
	if err != nil {
		fmt.Fprintf(stderr, "responsibility-table: inspect output path: %v\n", err)
		return 1
	}
	if overwritesSource {
		fmt.Fprintln(stderr, "responsibility-table: output path must not overwrite a Go source file")
		return 2
	}
	if err := output.Truncate(0); err != nil {
		fmt.Fprintf(stderr, "responsibility-table: write output: %v\n", err)
		return 1
	}
	if _, err := output.Write(renderMarkdown(rows, *localeArg)); err != nil {
		fmt.Fprintf(stderr, "responsibility-table: write output: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "Wrote %d row(s) to %s\n", len(rows), outputPath)
	return 0
}

func outputOverwritesGoFile(sourcePaths []string, outputInfo os.FileInfo) (bool, error) {
	for _, sourcePath := range sourcePaths {
		sourceInfo, err := os.Stat(sourcePath)
		if err != nil {
			return false, fmt.Errorf("inspect source file %s: %w", sourcePath, err)
		}
		if os.SameFile(outputInfo, sourceInfo) {
			return true, nil
		}
	}
	return false, nil
}

func collectGoFiles(root string) ([]string, error) {
	var paths []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path != root && entry.IsDir() && entry.Name() == ".git" {
			return filepath.SkipDir
		}
		if entry.Type()&os.ModeSymlink != 0 {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !entry.IsDir() && filepath.Ext(path) == ".go" {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)
	return paths, nil
}
