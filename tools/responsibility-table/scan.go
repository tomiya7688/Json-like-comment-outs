package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"os"
	"strings"
)

type commentBlock struct {
	parts       []string
	lineComment bool
	endOffset   int
	line        int
}

func readResponsibilityRows(path, relativePath string, semanticLookup map[string]string) ([]tableRow, []string, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	fileSet := token.NewFileSet()
	tokenFile := fileSet.AddFile(path, -1, len(source))
	var scanErrors []string
	var sourceScanner scanner.Scanner
	sourceScanner.Init(tokenFile, source, func(position token.Position, message string) {
		scanErrors = append(scanErrors, fmt.Sprintf("%s:%d:%d: %s", relativePath, position.Line, position.Column, message))
	}, scanner.ScanComments)

	var blocks []commentBlock
	for {
		position, tokenKind, literal := sourceScanner.Scan()
		if tokenKind == token.EOF {
			break
		}
		if tokenKind != token.COMMENT {
			continue
		}

		offset := tokenFile.Offset(position)
		line := tokenFile.Position(position).Line
		if strings.HasPrefix(literal, "//") {
			if len(blocks) > 0 {
				last := &blocks[len(blocks)-1]
				if last.lineComment && line == last.line+len(last.parts) && isWhitespace(source[last.endOffset:offset]) {
					last.parts = append(last.parts, literal)
					last.endOffset = offset + len(literal)
					continue
				}
			}
			blocks = append(blocks, commentBlock{parts: []string{literal}, lineComment: true, endOffset: offset + len(literal), line: line})
			continue
		}
		blocks = append(blocks, commentBlock{parts: []string{literal}, line: line})
	}

	rows := make([]tableRow, 0)
	for _, block := range blocks {
		entries, candidate, parseErr := extractResponsibility(block.text(), semanticLookup)
		if parseErr != nil {
			if candidate {
				scanErrors = append(scanErrors, fmt.Sprintf("%s:%d: %v", relativePath, block.line, parseErr))
			}
			continue
		}
		for _, entry := range entries {
			entry.source = relativePath
			rows = append(rows, entry)
		}
	}
	return rows, scanErrors, nil
}

func isWhitespace(value []byte) bool {
	return strings.TrimSpace(string(value)) == ""
}

func (block commentBlock) text() string {
	if !block.lineComment {
		text := strings.TrimPrefix(block.parts[0], "/*")
		text = strings.TrimSuffix(text, "*/")
		return strings.TrimSpace(text)
	}

	var builder strings.Builder
	for index, part := range block.parts {
		if index > 0 {
			builder.WriteByte(10)
		}
		line := strings.TrimPrefix(part, "//")
		if strings.HasPrefix(line, " ") {
			line = line[1:]
		}
		builder.WriteString(line)
	}
	return strings.TrimSpace(builder.String())
}
