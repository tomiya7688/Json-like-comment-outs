# JSON-like Comment Outs — 日本語ドキュメント

> **この日本語版を正本 (canonical) とします。**

JSON-like Comment Outs は、XMLドキュメントコメントのような宣言レベルの構造化コメントを、より書きやすく・読みやすくすることを目指すコメント方式です。

この方式では、固定された英単語ではなく**項目が持つ意味（Semantic Key）**を重視します。

日本語チームでは日本語キーを使えます。

```text
{
  責務: [
    GetUser: ユーザー情報を取得する
  ]
  処理: [
    1: IDを検証する
    2: Repositoryから取得する
  ]
  引数: [
    userId: 対象ユーザーID
  ]
  戻り値: [
    user: ユーザー情報
  ]
}
```

英語キーを使うこともできます。

```text
{
  Responsibility: [
    GetUser: ユーザー情報を取得する
  ]
  Action: [
    1: IDを検証する
    2: Repositoryから取得する
  ]
  Param: [
    userId: 対象ユーザーID
  ]
  Return: [
    user: ユーザー情報
  ]
}
```

重要なのはキーの綴りではなく、**同じ意味の情報がチーム内で一貫して表現されていること**です。

## Documents

- [Specification](./SPECIFICATION.md)
- [Templates](./TEMPLATE.md)
- [English documentation](../en/README.md)
