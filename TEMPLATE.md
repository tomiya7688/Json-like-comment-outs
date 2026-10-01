# JSON-like Comment Outs Templates

コピーして利用するための基本テンプレートです。

ルールは **必須 / 推奨 / 任意** に分かれています。詳細は [SPECIFICATION.md](./SPECIFICATION.md) を参照してください。

## Class

### 必須

```text
{
  Responsibility: 
  Fields: [
    fieldName: 
  ]
}
```

### 推奨を含む基本形

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

### 必須基本形

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

### 任意項目を含む Full

副作用やエラーも重要な場合:

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
  SideEffect: [
    
  ]
  Error: [
    ErrorName: 
  ]
  Note: [
    
  ]
}
```

## Variable Comment

意味のある状態・中間結果を保持するローカル変数では、次の形式を推奨します。

```text
{ Variable: variableName, Meaning: 変数が保持する値の意味・役割 }
```

```ts
// { Variable: user, Meaning: Repositoryから取得した未整形のユーザー情報 }
const user = repository.find(userId);
```

## TypeScript / JavaScript Example

関数・メソッド内部では、Action に対応する処理単位へ通常コメントを付けることを推奨します。

```ts
/*
{
  Responsibility: ユーザーIDからユーザー情報を取得する
  Action: [
    1: userIdの妥当性を確認する
    2: Repositoryから対象ユーザーを取得する
    3: 取得結果を返却形式へ変換する
    4: 呼び出し元へユーザー情報を返す
  ]
  Param: [
    userId: 取得対象ユーザーのID
  ]
  Return: [
    user: 取得したユーザー情報
  ]
  Error: [
    NotFound: 対象ユーザーが存在しない
  ]
}
*/
function getUser(userId) {
  // userIdの妥当性を確認する
  validateUserId(userId);

  // Repositoryから対象ユーザーを取得する
  // { Variable: user, Meaning: Repositoryから取得した未整形のユーザー情報 }
  const user = repository.find(userId);

  // 取得結果を返却形式へ変換する
  // { Variable: response, Meaning: APIへ返却するために整形済みのユーザー情報 }
  const response = toUserResponse(user);

  // 呼び出し元へユーザー情報を返す
  return response;
}
```

## Class Example

```ts
/*
{
  Responsibility: ユーザー認証の状態と認証処理を管理する
  Fields: [
    token: 現在利用している認証トークン
    user: ログイン中のユーザー情報
    isAuthenticated: 現在認証済みかどうか
  ]
  Action: [
    1: 認証情報を使ってログインする
    2: 認証結果とユーザー情報を保持する
    3: ログアウト時に認証状態を破棄する
  ]
}
*/
class AuthManager {
  login(credentials) {
    // 認証APIへログイン情報を送信する
    const result = authenticate(credentials);

    // 認証成功時の状態を保持する
    this.token = result.token;
    this.user = result.user;
    this.isAuthenticated = true;
  }

  logout() {
    // 保持している認証状態を破棄する
    this.token = null;
    this.user = null;
    this.isAuthenticated = false;
  }
}
```

## Python Example

Python では行コメントとして記述できます。

```py
# {
#   Responsibility: ユーザーIDからユーザー情報を取得する
#   Action: [
#     1: user_idの妥当性を確認する
#     2: Repositoryから対象ユーザーを取得する
#     3: 呼び出し元へユーザー情報を返す
#   ]
#   Param: [
#     user_id: 取得対象ユーザーのID
#   ]
#   Return: [
#     user: 取得したユーザー情報
#   ]
# }
def get_user(user_id):
    # user_idの妥当性を確認する
    validate_user_id(user_id)

    # Repositoryから対象ユーザーを取得する
    user = repository.find(user_id)

    # 呼び出し元へユーザー情報を返す
    return user
```

## Empty Param / Return

引数がない場合:

```text
Param: []
```

戻り値がない場合:

```text
Return: []
```

## Comment Layers

推奨する考え方:

```text
JSON-like Comment Out
  └─ クラス・関数全体の責務や処理フロー

通常コメント
  └─ 関数・メソッド内部の具体的な処理単位
```

JSON-like Comment Out だけですべての実装詳細を説明しようとせず、通常コメントと役割を分担してください。

## Notes

- 必須項目は原則として省略しません。
- 推奨項目・推奨ルールは、合理的な理由がなければ採用します。
- 任意項目は必要な場合だけ追加します。
- Action は実装の逐語訳ではなく、処理の目的と流れを書いてください。
- 関数・メソッド内部では、意味のある処理単位に通常コメントを付けることを推奨します。
- 意味のある状態・中間結果を保持するローカル変数には、`{ Variable: ..., Meaning: ... }` 形式のコメントを付けることを推奨します。
- 明白なループカウンタや自明な一時変数へのコメントは省略できます。
- テンプレートを埋めること自体を目的にせず、コード理解に必要な情報を残してください。
