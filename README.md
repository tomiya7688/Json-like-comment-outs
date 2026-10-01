# JSON-like Comment Outs

コードの直前に **JSON風の構造化コメント** を置き、クラスや関数の責務・処理・入出力を、人間にもAIにも読み取りやすくするための記法です。

> この記法は JSON そのものではありません。  
> 「項目名」「配列」「名前と説明の対応」を JSON 風に表現する、コメント用の軽量なドキュメント規約です。

## 目的

通常の文章コメントは自由度が高い一方で、情報の位置や粒度がばらつきやすくなります。

JSON-like Comment Outs では、コメントの形をある程度固定し、次の情報をコードを読む前に把握できることを目指します。

- そのクラス・関数が何を担当するか
- どのような処理を行うか
- どのような状態やフィールドを持つか
- 何を受け取り、何を返すか
- 必要に応じて、副作用やエラーが何か

また、JSON-like Comment Out は通常のコード内コメントを置き換えるものではありません。

**宣言直前の構造化コメントで全体像を示し、関数・メソッド内部では処理単位ごとに通常コメントを付ける**ことを推奨します。

また、意味のある状態や中間結果を保持するローカル変数には、変数の役割を示すコメントを付けることを推奨します。

## ルールレベル

ルールは次の3段階に分けます。

- **必須 (MUST)**: JSON-like Comment Outs として基本的に満たす項目
- **推奨 (SHOULD)**: 特別な理由がなければ採用する書き方
- **任意 (MAY)**: 必要な場合に追加する項目・書き方

詳細は [Specification](./SPECIFICATION.md) を参照してください。

## 基本形

### Class

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

クラスでは **Responsibility と Fields が必須**です。  
Action は、そのクラスが外部に対して提供する主要な振る舞いとして記述することを推奨します。

### Function / Method

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

関数では **Responsibility / Action / Param / Return が必須**です。

さらに、実装内部では通常コメントを処理単位に付けることを推奨します。

```ts
function getUser(userId) {
  // 入力値を検証する
  validateUserId(userId);

  // Repositoryから対象ユーザーを取得する
  const user = repository.find(userId);

  // 取得結果を返却形式へ変換する
  return toUserResponse(user);
}
```

構造化コメントは関数全体の設計図、通常コメントは実装を追うための道標として扱います。

## 変数コメント

意味のある状態・中間結果を保持するローカル変数には、宣言直前に次の形式でコメントを付けることを推奨します。

```ts
// { Variable: user, Meaning: Repositoryから取得した未整形のユーザー情報 }
const user = repository.find(userId);

// { Variable: response, Meaning: APIへ返却するために整形済みのユーザー情報 }
const response = toUserResponse(user);
```

基本形式:

```text
{ Variable: variableName, Meaning: 変数が保持する値の意味・役割 }
```

ループカウンタや、変数名と代入式だけで意味が十分に明白な一時変数まで機械的にコメントする必要はありません。

クラスフィールドはクラスの `Fields` で説明することを基本とし、同じ説明を各フィールド宣言へ重複して書くことは必須としません。

## 重要な考え方

Action はコードを日本語へ逐語訳するためのものではありません。

避けたい例:

```text
Action: [
  1: iを0にする
  2: for文を回す
  3: resultへpushする
]
```

推奨する例:

```text
Action: [
  1: 対象データを順番に検証する
  2: 条件を満たすデータだけ抽出する
  3: 抽出結果を返却形式へ変換する
]
```

実装方法ではなく、**処理の目的・意味・流れ** を記述します。

## 任意項目

必要な場合のみ、次の項目を追加できます。

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

すべての項目を毎回書く必要はありません。  
コメントがコード本体より重くならないことも重要です。

## ルールとテンプレート

- [Specification](./SPECIFICATION.md)
- [Templates](./TEMPLATE.md)

## License

MIT License. See [LICENSE](./LICENSE).
