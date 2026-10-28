# Responsibility Table Tool Specification

Status: **MVP Implemented (Go CUI)**

> This English edition is a translation. The Japanese version is canonical.  
> This specification records the implemented MVP scope.

The MVP scope is defined in Section 10. GUI, additional languages, and split output are deferred to later design and implementation issues.
## 1. Purpose

Extract responsibilities written in JSON-like Comment Outs and generate Markdown tables that provide an overview of a codebase.

Basic output:

```markdown
| Name | Responsibility |
| --- | --- |
| UserService | Manage user information |
| AuthManager | Manage authentication state and operations |
```

The tool reads Semantic Keys from comments instead of inferring responsibilities.

This tool analyzes **comments only**. Class and type names are also read from names written in Responsibility entries.

Recommended form:

```json
{
  "Responsibility": {
    "UserService": "Manage user information"
  }
}
```


## 2. Interfaces

### GUI (Post-MVP)

The GUI is a post-MVP extension. Its expected operations are:

- Select the target project / directory
- Select the output destination
- Select how to group the output
- Configure SemanticKeys
- Generate the table

Alt-based keyboard operation is planned for the GUI. Specific accelerator keys will be decided in a separate GUI design issue.

### CUI (MVP)

The MVP provides a CUI for Go source projects. It extracts responsibility comments and generates Markdown.

The command and options are defined in Section 10.

## 3. Output Modes

The tool will let users choose between:

### MVP: Combine into one file

Write all targets to the same Markdown file.

The default output path is responsibility-table.md under the input directory. Use --output to select another path.

### Later: Split by unit

Split the Markdown output by units such as Namespace / Package / Module, according to each language's structure. Each language adapter defines which unit to use.

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

The implementation order may change.

The MVP supports Go only. Add other languages through separate issues.

## 5. Declaration Units

The MVP does not parse source declarations. Each name-to-responsibility pair in a Responsibility object becomes one row.

Go struct, interface, and package units are not used in the MVP. Declaration units for Rust and other languages are deferred to a later design that handles source structure.

## 6. SemanticKeys

SemanticKeys resolve the meaning of keys in comments. The configuration is a simple JSON mapping of **canonical key → aliases**.

```json
{
  "SemanticKeys": {
    "Responsibility": ["Role", "Responsible"],
    "Action": ["Actions", "Steps"],
    "Fields": ["State"],
    "Param": ["Parameter", "Parameters", "Input"],
    "Returns": ["Output"],
    "SideEffect": ["SideEffects"],
    "Error": ["Errors"],
    "Note": ["Notes"]
  }
}
```

Rules:

- Canonical keys and aliases are **case-insensitive**.
- A canonical key is valid even when it is not listed in its own alias array.
- An alias does not need to be a translation or strict synonym; teams may choose any key they agree has the same meaning.
- The same alias must not be assigned to multiple canonical keys.
- The tool does not infer synonyms or use fuzzy matching for unregistered keys.

With the example above, the following keys all resolve to `Responsibility`:

```text
Responsibility
responsibility
RESPONSIBILITY
Role
Responsible
```

## 7. GUI SemanticKeys Settings (Post-MVP)

The GUI includes a **Settings** action for editing SemanticKeys.

Example:

```text
Semantic Keys

Responsibility
  Role
  Responsible

Action
  Actions
  Steps

[ + Alias ]
```

Expected operations:

- View canonical keys
- Add aliases
- Edit aliases
- Remove aliases
- Detect duplicate aliases

For the MVP, the default configuration is .json-like-comment-outs/semantic-keys.json under the input directory. The --semantic-keys option can select another file. Global settings and GUI editing are out of scope.

## 8. Responsibility Extraction

When generating the table, the tool uses values whose keys resolve to the canonical `Responsibility` through SemanticKeys.

Within `Responsibility`, it reads an object mapping each name to its responsibility:

```json
{
  "Responsibility": {
    "UserService": "Manage user information"
  }
}
```

`UserService` becomes the name column, and the text on the right becomes the responsibility column.

By default, the tool:

- Does not infer responsibilities absent from comments by inspecting code.
- Does not automatically generate responsibilities with AI.
- Does not treat unknown keys as `Responsibility`.

Because the MVP scans comments only, it cannot enumerate declarations that have no Responsibility entry. It emits neither an empty row nor a warning for such declarations.

## 9. Markdown

Use --locale ja|en, defaulting to ja. The headers are Name and Responsibility in English, or 名前 and 責務 in Japanese. Responsibility text is not translated.
Minimal form:

```markdown
| Name | Responsibility |
| --- | --- |
| UserService | Manage user information |
```

For languages where “Class” is not appropriate, the design may use a label such as “Type” or “Declaration”.

## 10. MVP Decisions

- The MVP is a Go-only CUI. It recursively scans Go source files under the input directory.
- The command is responsibility-table with required --source <dir> and optional --output <file>, --semantic-keys <file>, and --locale ja|en options.
- The tool does not parse source declarations or ASTs. Each name-to-responsibility pair in a Responsibility object becomes one row.
- Declarations without a Responsibility entry cannot be enumerated, so they produce neither a row nor a warning.
- Output is combined into one file, defaulting to responsibility-table.md under the input directory. Package-based and other split output is deferred.
- The default SemanticKeys configuration is .json-like-comment-outs/semantic-keys.json under the input directory. An explicit path takes precedence. If no file exists, use the default aliases in this specification.
- The default locale is ja. GUI, Alt accelerators, other languages, and Rust declaration units are deferred to later issues.
