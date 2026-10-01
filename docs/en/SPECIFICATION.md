# JSON-like Comment Outs Specification

Status: **Draft v0.7**

> This English edition is a translation. The Japanese specification is canonical.

## 1. Purpose

JSON-like Comment Outs aims to express declaration-level structured documentation, similar in purpose to XML documentation comments, in a form that is easier to write and read directly in source code.

The specification prioritizes:

1. Focusing on the explanation rather than markup
2. Quickly understanding responsibilities, actions, inputs, and outputs
3. Keeping a recognizable structure for humans, AI, and future tooling
4. Avoiding a documentation format that becomes heavier than the code itself

### Non-goals

This specification does not require:

- Valid JSON
- Full compatibility with XML documentation comments
- Replacing IDE or official documentation-generation features
- Converting every implementation comment into JSON-like syntax

## 2. Scope

JSON-like Comment Outs are primarily placed immediately before:

- Classes
- Functions
- Methods

Implementation details inside functions or methods, local variables, and branch intent should use ordinary comments.

## 3. Rule Levels

### MUST

Required for basic conformance.

### SHOULD

Recommended unless there is a reasonable reason not to follow the rule.

### MAY

Optional and used when useful.

## 4. Language

This English edition uses English keys in all examples.

Comments should still be written in the natural language that the team understands best.

## 5. Semantic Keys

The specification defines the **semantic role** of each key.

For this English edition, the recommended vocabulary is:

| Semantic role | Key |
| --- | --- |
| Responsibility of the target | Responsibility |
| Main actions or behavior | Action |
| State held by a class | Fields |
| Input values | Param / Parameters |
| Return value | Return |
| External state changes | SideEffect |
| Failure conditions | Error |
| Additional information | Note |

### SHOULD

- Use the same vocabulary consistently within a project
- Prefer terms the team understands immediately
- Fix the project vocabulary when tooling such as linting or documentation generation depends on stable keys

## 6. Basic Syntax

```text
{
  Key: value
  Key: [
    name: description
  ]
}
```

### MUST

- Start the block with `{` and end it with `}`
- Place `:` after a key
- Put separate fields on separate lines
- Use `[` and `]` for lists

### SHOULD

- Avoid excessive nesting
- Keep markup lighter than the explanation
- Avoid mixing unrelated responsibilities into one field
- Keep formatting easy to scan in source code

### MAY

- Commas
- Double quotes

Being parseable as strict JSON is not a requirement.

## 7. Classes

Classes emphasize responsibility and state.

### MUST

- `Responsibility`
- `Fields`
- Responsibility should explain the scope of the class
- Fields should explain the meaning, role, or state of important fields

### SHOULD

- `Action`
- Describe major class-level behavior
- Avoid listing low-level implementation steps

Example:

```text
{
  Responsibility: Manage authentication state and authentication operations
  Fields: [
    token: Current authentication token
    user: Currently authenticated user
  ]
  Action: [
    1: Log in with credentials
    2: Store authentication state
    3: Clear authentication state on logout
  ]
}
```

## 8. Functions / Methods

### MUST

- `Responsibility`
- `Action`
- `Param`
- `Return`
- Actions should normally follow execution order
- Actions should describe meaningful processing units

### SHOULD

- Number actions starting from 1
- Add ordinary comments for meaningful processing units inside the implementation
- Keep the declaration-level Action list roughly aligned with implementation sections
- Add ordinary comments for non-obvious local variables

Example:

```text
{
  Responsibility: Get user information
  Action: [
    1: Validate userId
    2: Fetch the user from the repository
    3: Convert the result to the API response format
  ]
  Param: [
    userId: ID of the target user
  ]
  Return: [
    user: User information in API response format
  ]
}
```

## 9. Optional Fields

### SideEffect

Use for externally visible state changes.

### Error

Use for important failure conditions that callers should understand.

### Note

Use for important supplemental information.

## 10. Implementation Comments

Implementation comments are outside the JSON-like syntax.

### Processing unit

```ts
// Query the repository only when the user is not cached
if (!cachedUser) {
  user = repository.find(userId);
}
```

### Variable

```ts
// Final amount after tax and discounts
const total = calculateTotal(order);
```

Prefer explaining purpose, reason, or meaning instead of simply restating the code.

## 11. Maintenance

### MUST

- Keep documentation consistent with the implementation
- Update comments when responsibility, actions, fields, inputs, or outputs change

### SHOULD

- Prefer readability over rigid formatting
- Do not fill optional fields mechanically
- Avoid comments that merely translate obvious code into prose

## 12. Tooling

Tool-specific configuration and behavior are defined in each tool's own documentation rather than in the core specification.

Potential future tooling includes:

- Validation
- XML Documentation conversion
- Markdown / HTML generation
- AI-oriented code-context extraction

See the [Responsibility Table tool](../../tools/responsibility-table/README.md) for the current support-tool design.
