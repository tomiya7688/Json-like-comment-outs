# Responsibility Table Tool Specification

Status: **Design Draft**

> This English edition is a translation. The Japanese version is canonical.  
> The tool is not implemented yet.

## 1. Purpose

Extract responsibilities from JSON-like Comment Outs and generate Markdown responsibility tables for a codebase.

Basic output:

```markdown
| Class | Responsibility |
| --- | --- |
| UserService | Manage user information |
| AuthManager | Manage authentication state and operations |
```

The tool does not infer responsibilities. It reads Semantic Keys from comments.

The tool is designed to analyze **comments only**. Declaration names are also read from the Responsibility entry.

Recommended form:

```text
Responsibility: [
  UserService: Manage user information
]
```

This allows the tool to generate the table without parsing the code declaration itself.

## 2. Interfaces

### GUI

The GUI is expected to provide:

- Project / directory selection
- Output destination
- Output grouping mode
- SemanticKeys settings
- Generation

Keyboard operation using **Alt-based accelerators** is part of the design.

### CUI

A CUI version will also be provided.

GUI and CUI will share the same parser, generator, and SemanticKeys configuration.

Exact commands and flags are not decided yet.

## 3. Output Grouping

Users can choose between:

- One Markdown file for all targets
- Multiple Markdown files grouped by Namespace / Package / Module or another language-appropriate unit

## 4. Language Priority

### First priority

1. Go
2. Python
3. C#

### Next priority

4. Ruby
5. Rust
6. C++
7. Java

## 5. Declaration Units

Class-based languages will generally use class-like declarations.

Go has no classes, so its responsibility unit is still undecided. Candidates include structs, interfaces, and packages.

Rust declaration units such as structs, enums, and traits will also be decided before implementation.

## 6. SemanticKeys

The tool resolves comment keys using a simple **canonical key → aliases** JSON mapping.

```json
{
  "SemanticKeys": {
    "Responsibility": ["Role", "Responsible"],
    "Action": ["Actions", "Steps"],
    "Fields": ["State"],
    "Param": ["Parameter", "Parameters", "Input"],
    "Return": ["Returns", "Output"],
    "SideEffect": ["SideEffects"],
    "Error": ["Errors"],
    "Note": ["Notes"]
  }
}
```

Rules:

- Canonical keys and aliases are **case-insensitive**
- The canonical key is recognized automatically
- Aliases do not need to be strict synonyms
- The same alias must not map to multiple canonical keys
- The tool performs no fuzzy matching or synonym inference for unmapped keys

## 7. GUI SemanticKeys Settings

The GUI will provide a Settings action for editing SemanticKeys aliases.

Expected operations:

- View canonical keys
- Add aliases
- Edit aliases
- Remove aliases
- Detect duplicate aliases

GUI changes will be shared with the CUI configuration.

Configuration file name, location, and Global / Project precedence are not decided yet.

## 8. Responsibility Extraction

The tool uses values resolved to the canonical `Responsibility` key.

Inside Responsibility, the tool reads `name: responsibility` pairs.

```text
Responsibility: [
  UserService: Manage user information
]
```

`UserService` becomes the name column and the value becomes the responsibility column.

By default, it should not:

- Infer missing responsibilities from code
- Generate responsibilities with AI
- Treat unknown keys as Responsibility

Handling declarations without a responsibility comment is still undecided.

## 9. Markdown

Minimal form:

```markdown
| Class | Responsibility |
| --- | --- |
| UserService | Manage user information |
```

Column naming may be localized or adjusted for languages where "Class" is not appropriate.

## 10. Open Questions

Before creating an implementation issue, decide:

- Go responsibility unit
- Rust responsibility unit
- Handling declarations without comments
- Output file naming
- Namespace / Package / Module grouping details
- SemanticKeys configuration file name and location
- Exact GUI Alt accelerators
- CUI command / option names
- Markdown column localization
