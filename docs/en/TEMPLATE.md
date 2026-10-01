# JSON-like Comment Outs Templates

This English edition uses **English keys only**.

## Class

```text
{
  Responsibility:
  Fields: [
    fieldName:
  ]
  Action: [
    1:
  ]
}
```

## Function / Method

```text
{
  Responsibility:
  Action: [
    1:
    2:
  ]
  Param: [
    name:
  ]
  Return: [
    name:
  ]
}
```

## Full

```text
{
  Responsibility:
  Action: [
    1:
    2:
  ]
  Param: [
    name:
  ]
  Return: [
    name:
  ]
  SideEffect: []
  Error: []
  Note: []
}
```

## TypeScript / JavaScript Example

```ts
/*
{
  Responsibility: Get user information by user ID
  Action: [
    1: Validate userId
    2: Fetch the user from the repository
    3: Convert the result to the API response format
    4: Return the result to the caller
  ]
  Param: [
    userId: ID of the target user
  ]
  Return: [
    user: User information in API response format
  ]
  Error: [
    NotFound: The target user does not exist
  ]
}
*/
function getUser(userId) {
  // Validate the user ID
  validateUserId(userId);

  // Raw user data returned by the repository
  const user = repository.find(userId);

  // Convert to the API response format
  const response = toUserResponse(user);

  return response;
}
```

## Implementation Comments

Implementation details use ordinary comments.

```ts
// Final amount after tax and discounts
const total = calculateTotal(order);
```
