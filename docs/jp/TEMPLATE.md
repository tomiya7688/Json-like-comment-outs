# JSON-like Comment Outs Templates

> 日本語版は正本です。日本語キーと英語キーの両方を例示します。

## 日本語 — Class

```text
{
  責務: [
    ClassName:
  ]
  フィールド: [
    fieldName:
  ]
  処理: [
    1:
  ]
}
```

## 日本語 — Function / Method

```text
{
  責務: [
    FunctionName:
  ]
  処理: [
    1:
    2:
  ]
  引数: [
    name:
  ]
  戻り値: [
    name:
  ]
}
```

## 日本語 — Full

```text
{
  責務: [
    FunctionName:
  ]
  処理: [
    1:
    2:
  ]
  引数: [
    name:
  ]
  戻り値: [
    name:
  ]
  副作用: []
  エラー: []
  補足: []
}
```

## English Keys — Class

```text
{
  Responsibility: [
    ClassName:
  ]
  Fields: [
    fieldName:
  ]
  Action: [
    1:
  ]
}
```

## English Keys — Function / Method

```text
{
  Responsibility: [
    FunctionName:
  ]
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

## 実装内部

処理単位や変数説明には通常コメントを使用します。

```ts
// Repositoryから取得した未整形のユーザー情報
const user = repository.find(userId);
```
