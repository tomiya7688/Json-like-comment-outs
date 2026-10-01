# JSON-like Comment Outs — English Documentation

> This English documentation is a translation. The Japanese documentation is canonical.

JSON-like Comment Outs is an experimental structured comment style designed to make declaration-level documentation easier to write and read than XML-style documentation comments.

This edition uses **English keys only**.

```text
{
  Responsibility: Get user information
  Action: [
    1: Validate the user ID
    2: Fetch the user from the repository
  ]
  Param: [
    userId: ID of the target user
  ]
  Return: [
    user: User information
  ]
}
```

The important part is not a specific spelling of a key, but the semantic role represented by that key. In this English edition, examples use a consistent English vocabulary.

## Documents

- [Specification](./SPECIFICATION.md)
- [Templates](./TEMPLATE.md)
- [Canonical Japanese documentation](../jp/README.md)
