# JSON-like Comment Outs Specification

Status: **Draft v0.4**

この文書は JSON-like Comment Outs の基本ルールを定義します。

## 1. Purpose

JSON-like Comment Outs は、XMLドキュメントコメントのような**宣言レベルの構造化コメントを、より書きやすく・読みやすい形式で表現すること**を目的とします。

特に次を重視します。

1. タグやマークアップより、説明そのものへ集中できること
2. ソースコードを開いたときに、責務・処理・入出力を素早く把握できること
3. 一定の構造を保ち、人間・AI・将来のツールが解釈しやすいこと
4. コメント規約そのものが大きな記述負担にならないこと

### Non-goals

本仕様は次を目的としません。

- 有効な JSON として解釈できること
- XMLドキュメントコメントとの完全互換
- IDEや公式ドキュメント生成機能を直接置き換えること
- すべての通常コメントを JSON風に統一すること

## 2. Scope

JSON-like Comment Out の対象は、原則として次の宣言です。

- クラス
- 関数
- メソッド

対象の**直前**に配置し、その宣言全体の説明を行います。

関数・メソッド内部の処理説明、ローカル変数の説明、条件分岐の意図などは通常コメントの領域とし、JSON-like 形式を要求しません。

## 3. Rule Levels

### 必須 (MUST)

JSON-like Comment Outs として基本的に満たすルールです。

### 推奨 (SHOULD)

合理的な理由がない限り採用することを推奨するルールです。

### 任意 (MAY)

必要な場合だけ利用するルール・項目です。

## 4. Placement

### 必須

JSON-like Comment Out は対象となるクラス・関数・メソッドの**直前**に配置します。

```text
<JSON-like Comment Out>
class / function / method ...
```

コメント構文自体は使用言語に従います。

### 推奨

関数・メソッド内部では、意味のある処理単位に通常コメントを配置します。

```ts
function getUser(userId) {
  // 入力値を検証する
  validateUserId(userId);

  // Repositoryから対象ユーザーを取得する
  const user = repository.find(userId);

  // API返却形式へ変換する
  return toUserResponse(user);
}
```

JSON-like Comment Out は**宣言全体の設計図**、通常コメントは**実装を追うための説明**として役割を分けます。

## 5. Internal Comments

内部コメントは JSON-like Comment Outs の構文対象外です。

### 5.1 処理単位コメント

意味のある処理のまとまりについて、通常コメントを付けることを推奨します。

```ts
// キャッシュに存在しない場合のみRepositoryへ問い合わせる
if (!cachedUser) {
  user = repository.find(userId);
}
```

コードをそのまま言い換えるのではなく、処理の目的・理由を優先します。

### 5.2 Variable Comments

変数名だけでは**意味・由来・状態・単位**などが十分に伝わらないローカル変数には、通常コメントを付けることを推奨します。

```ts
// Repositoryから取得した、まだAPI形式へ変換していないユーザー情報
const user = repository.find(userId);

// 税込み・割引適用後の最終請求額
const total = calculateTotal(order);
```

特別な JSON-like 構文は使用しません。

次のような場合は省略できます。

- `i`, `j` など用途が明白なループカウンタ
- 名前と代入式だけで意味が明確な一時変数
- 説明を付けてもコードの読みやすさがほとんど向上しない場合

クラスのフィールドについては、宣言前コメントの `Fields` が主要な説明場所です。必要に応じて実フィールドにも通常コメントを追加できます。

## 6. Basic Syntax

基本形:

```text
{
  Key: value
  Key: [
    name: description
  ]
}
```

### 必須

- ブロックは `{` で始まり `}` で終える
- 項目名の後ろには `:` を置く
- 複数項目は改行して記述する
- リストは `[` と `]` で囲む

### 推奨

- 過度なネストを避ける
- 説明より記号が目立つ構造にしない
- 一つの項目へ複数の異なる責務を詰め込まない
- ソースコード上で流し読みしやすい整形を維持する

### 任意

- カンマ
- ダブルクォート

有効な JSON に変換できることは要件ではありません。

## 7. Common Keys

### Responsibility

対象が何を担当するか、なぜ存在するかを記述します。

```text
Responsibility: ユーザー認証の状態と認証処理を管理する
```

### Action

対象が行う処理・振る舞いを意味のある単位で記述します。

```text
Action: [
  1: 入力値を検証する
  2: Repositoryから対象データを取得する
  3: 返却形式へ変換する
]
```

実装コードの逐語訳は避けます。

### Fields

クラスが保持する主要な状態を記述します。

```text
Fields: [
  token: 現在の認証トークン
  user: ログイン中のユーザー情報
  isAuthenticated: 現在の認証状態
]
```

型だけでなく、そのフィールドの意味・役割・状態を記述します。

### Param

関数・メソッドが受け取る値を記述します。

```text
Param: [
  userId: 取得対象ユーザーのID
]
```

引数がない場合:

```text
Param: []
```

### Return

戻り値の意味を記述します。

```text
Return: [
  user: 取得したユーザー情報
]
```

戻り値がない場合:

```text
Return: []
```

## 8. Class Rules

クラスでは、責務と保持する状態を特に重視します。

### 必須

- `Responsibility`
- `Fields`
- Responsibility はクラスの担当範囲が分かる内容にする
- Fields は主要フィールドの意味・役割・状態が分かる内容にする

### 推奨

- `Action`
- Action にはクラスが提供する主要な振る舞いを書く
- 内部実装の細かい手順ではなく、クラス単位の振る舞いを書く

### 任意

- `SideEffect`
- `Error`
- `Note`
- プロジェクト固有の追加キー

## 9. Function / Method Rules

関数・メソッドでは、処理の流れを特に重視します。

### 必須

- `Responsibility`
- `Action`
- `Param`
- `Return`
- Action は原則として実際の処理順に並べる
- Action は意味のある処理単位で記述する

### 推奨

- Action の番号は `1` から始めて連番にする
- 実装内部では意味のある処理単位に通常コメントを付ける
- Action と実装内部の処理区切りがおおむね対応するようにする
- 通常コメントでは、コードだけでは分かりにくい「なぜ」も補足する
- 非自明なローカル変数には通常コメントで意味を補足する

### 任意

- `SideEffect`
- `Error`
- `Note`
- プロジェクト固有の追加キー

## 10. Optional Keys

### SideEffect

外部状態を変更する場合に使用します。

```text
SideEffect: [
  DBの注文情報を更新する
  在庫数を減算する
]
```

### Error

呼び出し元が意識すべき重要な失敗条件を記述します。

```text
Error: [
  NotFound: 対象データが存在しない
  PaymentError: 決済処理に失敗した
]
```

### Note

その他の重要な補足事項に使用します。

```text
Note: [
  この処理は冪等である
]
```

## 11. Writing Guidelines

### 必須

- コメントの内容を実装と一致させる
- 責務・処理・フィールド・入出力の意味が変わった場合はコメントも更新する

### 推奨

- マークアップより説明そのものを優先する
- 「何をするか」「なぜ存在するか」が分かる言葉を使う
- Action は処理の目的・意味・流れを書く
- Fields はフィールドの役割・保持する状態を書く
- Param / Return は値の意味を書く
- 内部コメントは構文説明より、処理のまとまり・理由・変数の意味を書く

### 避ける

- コードをそのまま自然言語へ置き換える
- 自明な構文を逐一説明する
- 実装と一致しない古いコメントを残す
- コメントの形式を守るためだけに冗長な説明を書く
- すべての内部コメントを JSON-like にする

## 12. Comment Model

本仕様ではコメントを二種類に分けます。

### Declaration Documentation

JSON-like Comment Out を使用します。

対象:

- クラス
- 関数
- メソッド

役割:

- 責務
- 主な処理
- フィールド
- 入出力
- 副作用
- エラー

### Implementation Comments

通常コメントを使用します。

対象:

- 処理単位
- 条件分岐
- ループ
- 非自明なローカル変数
- 実装上の理由・注意点

この分離により、JSON-like Comment Outs の構造化という利点を残しながら、実装内部まで形式で縛りすぎないことを目指します。

## 13. Tooling

JSON-like Comment Outs は現時点ではソースコード上の可読性を第一目的とします。

XMLドキュメントコメントのようなIDE連携やドキュメント生成は仕様上保証しません。

将来的には、必要に応じて次のようなツールを実装できます。

- JSON-like Comment Out の検証
- XML Documentation への変換
- Markdown / HTML ドキュメント生成
- AI向けコードコンテキスト抽出

## 14. Extensibility

プロジェクト固有の項目を追加して構いません。

ただし、共通キーの意味を変更することは推奨しません。

```text
Permission: AdminOnly
Cache: 5min
Transaction: Required
```
