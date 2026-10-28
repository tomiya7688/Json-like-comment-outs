package main

import (
	"fmt"
	"html"
	"strings"
)

func renderMarkdown(rows []tableRow, locale string) []byte {
	nameHeader := "名前"
	responsibilityHeader := "責務"
	if locale == "en" {
		nameHeader = "Name"
		responsibilityHeader = "Responsibility"
	}

	var builder strings.Builder
	fmt.Fprintf(&builder, "| %s | %s |\n", nameHeader, responsibilityHeader)
	builder.WriteString("| --- | --- |\n")
	for _, row := range rows {
		fmt.Fprintf(&builder, "| %s | %s |\n", escapeMarkdownCell(row.name), escapeMarkdownCell(row.responsibility))
	}
	return []byte(builder.String())
}

func escapeMarkdownCell(value string) string {
	value = html.EscapeString(value)
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "|", "\\|")
	value = strings.ReplaceAll(value, "`", "\\`")
	value = strings.ReplaceAll(value, "\r\n", " ")
	value = strings.ReplaceAll(value, "\r", " ")
	value = strings.ReplaceAll(value, "\n", " ")
	return value
}
