# JSON-like Comment Outs

**XMLドキュメントコメントより、書きやすく・読みやすい構造化コメントを目指す実験的なコメント方式です。**

JSON-like Comment Outs は、クラス・関数・メソッドの宣言直前に JSON風の構造を持つコメントを書きます。

重要なのは固定された英語キーではなく、**Key が持つ共通の意味 (Semantic Key)** です。

そのため、コメントはコードを読むチームが最も読みやすい言語で書くことを推奨します。

## 日本語の例

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

## English example

```text
{
  Responsibility: [
    GetUser: Get user information
  ]
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

どちらも、責務・処理・入力・出力という同じ意味を持つ構造として扱います。

> **構造は共有する。言語はチームに合わせる。**

## Why

XML Documentation は IDE やドキュメント生成との連携に優れていますが、ソースコード上ではタグやマークアップの記述量が大きくなることがあります。

JSON-like Comment Outs では、マークアップよりも**コードの意図を書くこと**へ集中できる形式を目指します。

厳密な JSON であることや、XML Documentation との完全互換は目的としていません。

## Scope

JSON-like Comment Out は主に宣言レベルで使用します。

- Class
- Function
- Method

実装内部の処理単位・条件分岐・ローカル変数などは、通常コメントを使用します。

```ts
// Repositoryから取得した未整形のユーザー情報
const user = repository.find(userId);
```

内部コメントまで JSON-like にする必要はありません。

ツール解析をしやすくするための書き方は**推奨**として定義しますが、解析都合のために人間向けコメントを厳しく縛ることはしません。

## Documentation

### 日本語 — Canonical

日本語ドキュメントを**正本 (canonical)** とします。  
日本語キーと英語キーの両方を扱います。

- [日本語ドキュメント](./docs/jp/README.md)
- [仕様](./docs/jp/SPECIFICATION.md)
- [テンプレート](./docs/jp/TEMPLATE.md)

### English

English documentation is a translation and uses **English JSON-like keys only**.

- [English documentation](./docs/en/README.md)
- [Specification](./docs/en/SPECIFICATION.md)
- [Templates](./docs/en/TEMPLATE.md)

## Tools

サポートツールは [tools](./tools/README.md) 以下で設計・管理します。

最初の候補は [責務表生成ツール](./tools/responsibility-table/README.md) です。

## License

MIT License. See [LICENSE](./LICENSE).
