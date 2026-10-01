# JSON-like Comment Outs Specification

Status: **Draft v0.1**

この文書は JSON-like Comment Outs の基本ルールを定義します。

## 1. Scope

JSON-like Comment Outs は、クラス・関数・メソッドなどの宣言直前に置く、構造化された説明コメントです。

目的は次の2点です。

1. コードを読む前に、責務・状態・処理・入出力を把握できること
2. 人間とAIの双方が、一定の構造として読み取りやすいこと

この記法は有効な JSON である必要はありません。

## 2. Placement

コメントブロックは、原則として対象となるクラス・関数・メソッドの**直前**に配置します。

```text
<JSON-like Comment Out>
class / function / method ...
```

コメント構文そのものは使用言語に従います。

## 3. Basic Syntax

基本形:

```text
{
  Key: value
  Key: [
    name: description
  ]
}
```

### 3.1 Syntax rules

- ブロックは `{` で始まり `}` で終える
- 項目名の後ろには `:` を置く
- 複数項目は改行して記述する
- リストは `[` と `]` で囲む
- カンマ、ダブルクォートは必須ではない
- 有効な JSON に変換することを目的にしない
- 読みやすさを損なう過度なネストは避ける

## 4. Common Keys

### Responsibility

対象が「何を担当するか」を記述します。

実装方法ではなく、存在理由・責務を簡潔に表現します。

```text
Responsibility: ユーザー認証の状態と認証処理を管理する
```

### Action

対象が行う処理・振る舞いを、意味のある単位で順序付きに記述します。

```text
Action: [
  1: 入力値を検証する
  2: Repositoryから対象データを取得する
  3: 返却形式へ変換する
]
```

Action は実装コードの逐語訳にしません。

避ける:

```text
1: iを0にする
2: for文を回す
3: resultへpushする
```

推奨:

```text
1: 対象データを順番に検証する
2: 条件を満たすデータを抽出する
3: 抽出結果を返却形式へ変換する
```

### Fields

クラスが保持する状態を記述します。

```text
Fields: [
  token: 現在の認証トークン
  user: ログイン中のユーザー情報
  isAuthenticated: 現在の認証状態
]
```

単なる型名だけでなく、**そのフィールドが何のために存在するか**が分かる説明を推奨します。

### Param

関数・メソッドが受け取る値を記述します。

```text
Param: [
  userId: 取得対象ユーザーのID
]
```

引数がない場合は次のように空配列で表現できます。

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

## 5. Class Rules

クラスでは、**責務と保持する状態を詳しく記述すること**を重視します。

### Required

- `Responsibility`
- `Fields`

### Recommended

- `Action`

### Optional

- `SideEffect`
- `Error`
- `Note`

基本形:

```text
{
  Responsibility: このクラスが担当する責務
  Fields: [
    fieldName: フィールドが保持する状態・目的
  ]
  Action: [
    1: このクラスが提供する主な振る舞い
    2: 別の主な振る舞い
  ]
}
```

クラスの Action は、内部実装の細かい手順よりも、**クラスとして提供する主要な振る舞い**を記述します。

## 6. Function / Method Rules

関数・メソッドでは、**処理の流れを詳しく記述すること**を重視します。

### Required

- `Responsibility`
- `Action`
- `Param`
- `Return`

### Optional

- `SideEffect`
- `Error`
- `Note`

基本形:

```text
{
  Responsibility: この関数が担当する処理
  Action: [
    1: 最初に行う意味のある処理
    2: 次に行う意味のある処理
    3: 結果を作成する
  ]
  Param: [
    name: 引数の意味
  ]
  Return: [
    name: 戻り値の意味
  ]
}
```

### Action ordering

関数・メソッドの Action は、原則として実際の処理順に並べます。

番号は `1` から始め、連番にすることを推奨します。

## 7. Optional Keys

### SideEffect

対象の外部に状態変化を発生させる場合に使用します。

```text
SideEffect: [
  DBの注文情報を更新する
  在庫数を減算する
]
```

### Error

重要な失敗条件や、呼び出し元が意識すべきエラーを記述します。

```text
Error: [
  NotFound: 対象データが存在しない
  PaymentError: 決済処理に失敗した
]
```

### Note

他の項目に適さない重要な補足事項に使用します。

```text
Note: [
  この処理は冪等である
]
```

## 8. Writing Guidelines

### SHOULD

- 「なぜ存在するか」「何をするか」が分かる言葉を使う
- Action と実際の処理順を対応させる
- Fields ではフィールドの役割・保持する状態を書く
- Param / Return では値の意味を書く
- コード変更時にコメントも更新する
- コードだけでは読み取りにくい意図を優先して書く

### SHOULD NOT

- コードをそのまま日本語へ置き換える
- 自明な構文を逐一説明する
- 実装と一致しない古いコメントを残す
- 必要以上に巨大なコメントを作る
- すべての任意項目を機械的に追加する

## 9. Maintenance Rule

JSON-like Comment Out は対象コードの一部として扱います。

責務・処理・フィールド・入出力の意味が変わった場合、コードと同じ変更単位でコメントも更新してください。

古い構造化コメントは、コメントが存在しない場合より誤解を招く可能性があります。

## 10. Extensibility

プロジェクト固有の項目を追加しても構いません。

ただし、共通項目の意味を変更することは推奨しません。

例:

```text
Permission: AdminOnly
Cache: 5min
Transaction: Required
```

独自項目を多数利用する場合は、プロジェクト側で意味を定義してください。
