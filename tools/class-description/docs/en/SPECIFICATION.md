# Class Description Generator Specification

Status: **Design Draft**

> This English edition is a translation. The Japanese version is canonical.  
> The tool is not implemented yet.

## 1. Purpose

Generate per-class Markdown documentation by analyzing comments in source code.

Primary sources:

- JSON-like Comment Outs before classes
- JSON-like Comment Outs before functions / methods
- Ordinary comments inside functions / methods

The tool restructures information already written in comments.

A central idea is **comment-driven documentation**: if comments are written carefully, they can be turned directly into specification or documentation.

For specification-first development, teams may write declaration comments before implementation, generate the Markdown specification, and then implement according to those comments.

## 2. No Implementation Validation

This tool does not verify:

- Whether comments match the implementation
- Whether the implementation is correct
- Whether responsibilities are appropriate
- Whether parameter or return descriptions match actual behavior
- Whether internal comments correctly describe the code

For this tool, **comments are the source of truth**.

## 3. Basic Output

```markdown
# UserService

## Responsibility

Manage user information.

## Fields

| Field | Description |
| --- | --- |
| repository | Repository used to access user information |
| cache | Cache for previously retrieved users |

## Functions

| Function | Responsibility |
| --- | --- |
| GetUser | Get user information by ID |
| UpdateUser | Update user information |

## Function Details

### GetUser

#### Responsibility

Get user information by ID.

#### Parameters

| Name | Description |
| --- | --- |
| userId | ID of the target user |

#### Return

| Name | Description |
| --- | --- |
| user | Retrieved user information |

#### Processing

1. Validate the input
2. Fetch the user from the repository
3. Convert to the API response format
```

## 4. Class Fields

Fields from the class-level JSON-like Comment Out are included in the generated class document.

```markdown
## Fields

| Field | Description |
| --- | --- |
| repository | Repository used to access user information |
| cache | Cache for previously retrieved users |
```

The tool does not validate whether these comments match actual field declarations.

## 5. Internal Comments

Ordinary comments inside a function are collected in source order.

The tool does not summarize, infer, or add missing processing steps.

## 6. JSON-like Action vs Internal Comments

Declaration-level Action and ordinary implementation comments are separate information sources.

Whether the first version displays only implementation comments or both sources is still undecided.

## 7. SemanticKeys

JSON-like declaration comments are resolved using SemanticKeys.

Ordinary implementation comments are plain text and do not use SemanticKeys.

## 8. Output Units

The primary design is one Markdown document per class-like unit.

Alternative grouping modes may be supported.

## 9. GUI / CUI

Both GUI and CUI versions are planned and should share the same parser and generator.

## 10. Language Priority

First priority:

1. Go
2. Python
3. C#

Next:

4. Ruby
5. Rust
6. C++
7. Java

For languages without classes, such as Go, the documentation unit must be defined separately.

## 11. Open Questions

Before implementation issues are created, decide:

- Final tool name
- Action vs internal-comment output
- Default file grouping
- Missing-comment behavior
- Visibility filtering
- Constructors and properties
- Overloads
- Go / Rust documentation units
- GUI / CUI details
