# JSON-like Comment Outs

**XMLドキュメントコメントより、書きやすく・読みやすい構造化コメントを目指す実験的なコメント方式です。**

クラス・関数・メソッドの宣言直前に、JSON風の構造を持つコメントを書きます。

> この記法は JSON そのものではありません。  
> また、XMLドキュメントコメントと機能的に互換であることを目的としていません。

## Why

XMLドキュメントコメントは、IDEやドキュメント生成との連携に優れています。

一方で、ソースコード上ではタグが多くなりやすく、説明そのものよりマークアップの比率が大きくなる場合があります。

例えば:

```csharp
/// <summary>
/// ユーザーIDからユーザー情報を取得する。
/// </summary>
/// <param name="userId">取得対象ユーザーのID。</param>
/// <returns>取得したユーザー情報。</returns>
User GetUser(string userId)
```

JSON-like Comment Outs では、同じ種類の情報を次のように表現します。

```text
{
  Responsibility: ユーザーIDからユーザー情報を取得する
  Action: [
    1: userIdを検証する
    2: Repositoryから対象ユーザーを取得する
    3: 取得結果を返す
  ]
  Param: [
    userId: 取得対象ユーザーのID
  ]
  Return: [
    user: 取得したユーザー情報
  ]
}
```

このリポジトリでは、**マークアップを書くことより、コードの意図を書くことへ集中できる形式**を目指します。

## Goals

- XML風のタグ構造より、ソース上で素早く読めること
- コメントを書く際の記述量を抑えること
- クラス・関数の責務、処理、入出力を一定の形で表現できること
- 人間だけでなく、AIやツールからも構造を認識しやすいこと
- コメント規約がコード本体より重くならないこと

## Non-goals

- 正式な JSON として解釈できること
- XMLドキュメントコメントとの完全互換
- IDEの IntelliSense やドキュメント生成機能を、そのまま置き換えること
- ソースコード中のすべてのコメントを JSON風にすること

必要であれば、将来的にこの形式から既存のドキュメント形式へ変換するツールを作る余地はあります。

## Scope

JSON-like Comment Outs の対象は、主に**宣言レベルのドキュメントコメント**です。

- クラス
- 関数
- メソッド

関数内部の処理説明やローカル変数の説明は、通常のコメントを使います。

```ts
function getUser(userId) {
  // userIdの形式を検証する
  validateUserId(userId);

  // Repositoryから取得した未整形のユーザー情報
  const user = repository.find(userId);

  // API返却用の形式へ変換する
  return toUserResponse(user);
}
```

**JSON-like = 宣言前の構造化ドキュメント**  
**通常コメント = 実装内部の説明**

という役割分担を基本とします。

## Rule Levels

ルールは次の3段階に分けます。

- **必須 (MUST)**: JSON-like Comment Outs として満たす基本ルール
- **推奨 (SHOULD)**: 特別な理由がなければ採用するルール
- **任意 (MAY)**: 必要な場合だけ利用する項目・書き方

## Class

```text
{
  Responsibility: このクラスが担当する責務
  Fields: [
    fieldName: フィールドの意味・保持する状態
  ]
  Action: [
    1: このクラスが提供する主な振る舞い
    2: 別の主な振る舞い
  ]
}
```

### 必須

- `Responsibility`
- `Fields`

### 推奨

- `Action`

## Function / Method

```text
{
  Responsibility: この関数が担当する処理
  Action: [
    1: 入力を検証する
    2: 必要なデータを取得する
    3: 結果を組み立てる
    4: 呼び出し元へ返す
  ]
  Param: [
    name: 引数の意味
  ]
  Return: [
    name: 戻り値の意味
  ]
}
```

### 必須

- `Responsibility`
- `Action`
- `Param`
- `Return`

## Internal Comments

関数・メソッド内部では、**意味のある処理単位に通常コメントを付けることを推奨**します。

また、変数名だけでは意味・由来・状態が分かりにくいローカル変数についても、必要に応じて通常コメントを付けることを推奨します。

```ts
// Repositoryから取得した、まだAPI形式へ変換していないユーザー情報
const user = repository.find(userId);
```

次のような自明な変数まで機械的にコメントする必要はありません。

```ts
for (let i = 0; i < items.length; i++) {
  // ...
}
```

内部コメントには特別な JSON-like 形式を要求しません。

## Action

Action はコードを逐語的に日本語へ置き換えるためのものではありません。

避けたい例:

```text
Action: [
  1: iを0にする
  2: for文を回す
  3: resultへpushする
]
```

推奨:

```text
Action: [
  1: 対象データを順番に検証する
  2: 条件を満たすデータを抽出する
  3: 抽出結果を返却形式へ変換する
]
```

**実装方法より、処理の目的・意味・流れを書く**ことを重視します。

## Optional Keys

必要に応じて追加できます。

```text
SideEffect: [
  DBの状態を更新する
]

Error: [
  NotFound: 対象が存在しない
]

Note: [
  補足事項
]
```

## Documents

- [Specification](./SPECIFICATION.md)
- [Templates](./TEMPLATE.md)

## License

MIT License. See [LICENSE](./LICENSE).
