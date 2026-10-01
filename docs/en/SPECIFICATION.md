# JSON-like Comment Outs Specification

Status: **Draft v0.9**

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
  Responsibility: [
    AuthManager: Manage authentication state and authentication operations
  ]
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
  Responsibility: [
    GetUser: Get user information
  ]
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

## 9. Recommendations for Tool Compatibility

The rules in this section are **SHOULD-level recommendations** intended to make comments easier for parsers, validators, generators, and other tools to handle consistently.

**A comment is not invalid JSON-like Comment Outs merely because it does not follow these recommendations.**

Human readability may take priority over tooling convenience.

### Association with declarations

Place a JSON-like Comment Out immediately before its target declaration when practical.

If decorators, attributes, or annotations belong to that declaration, tools may treat the comment as belonging to the declaration across those constructs.

### Include the declaration name in Responsibility

To let comment-only tools obtain both the target name and its responsibility, the following form is recommended.

Class:

```text
Responsibility: [
  UserService: Manage user information
]
```

Function / method:

```text
Responsibility: [
  GetUser: Get user information by user ID
]
```

With this form, tools such as responsibility-table and class-description generators can work from comments without parsing the declaration itself.

This is a **SHOULD-level recommendation**, not a requirement.

### Param / Fields names

Use actual implementation names for parameter and field entries when practical.

```text
Param: [
  userId: ID of the target user
]
```

This makes mechanical comparison between comments and declarations possible.

### Multiple return values

For languages with multiple return values, use names when available.

When names are unavailable, numeric positions may be used.

```text
Return: [
  1: User information
  2: Error
]
```

List order may then represent return-value order.

### Duplicate Semantic Keys

Avoid writing multiple keys in one block that resolve to the same Semantic Key.

For example, if `Role` and `Responsibility` both resolve to the same meaning, avoid:

```text
{
  Role: AAA
  Responsibility: BBB
}
```

Tools otherwise cannot know which value should take precedence.

### Empty templates

Incomplete templates generated by tooling may contain empty values.

```text
{
  Responsibility: [
    GetUser:
  ]
  Action: []
  Param: [
    userId:
  ]
  Return: []
}
```

Such comments may still be parseable even though their documentation is incomplete.

Validators may warn about missing descriptions separately.

### Unknown keys

Project-specific keys are allowed.

```text
{
  Responsibility: ...
  Permission: AdminOnly
  Cache: 5min
}
```

Tools should not guess that unknown keys are equivalent to known Semantic Keys.

### Key order

Top-level key order should not carry semantic meaning.

However, order inside lists may carry meaning, for example:

- Action execution order
- Multiple return-value order

### Action and implementation comments

Treat Action as a declaration-level processing overview.

Treat ordinary comments inside a function or method as implementation-level processing detail.

A one-to-one match between them is not required.

### Language-specific declarations

Languages without classes, or with different type systems, may apply the same ideas to language-appropriate declaration units.

Examples include:

- Go: struct / interface / package
- Rust: struct / enum / trait

Each tool defines which declarations it supports.

## 10. Optional Fields

### SideEffect

Use for externally visible state changes.

### Error

Use for important failure conditions that callers should understand.

### Note

Use for important supplemental information.

## 11. Implementation Comments

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

## 12. Maintenance

### MUST

- Keep documentation consistent with the implementation
- Update comments when responsibility, actions, fields, inputs, or outputs change

### SHOULD

- Prefer readability over rigid formatting
- Do not fill optional fields mechanically
- Avoid comments that merely translate obvious code into prose

## 13. Tooling

Tool-specific configuration and behavior are defined in each tool's own documentation rather than in the core specification.

Potential future tooling includes:

- Validation
- XML Documentation conversion
- Markdown / HTML generation
- AI-oriented code-context extraction

See the [Responsibility Table tool](../../tools/responsibility-table/README.md) for the current support-tool design.
