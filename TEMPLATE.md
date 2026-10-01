# JSON-like Comment Outs Templates

コピーして利用するための基本テンプレートです。

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

## Function / Method - Full

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

## TypeScript / JavaScript Example

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
  // ...
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
  // ...
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
    pass
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

## Minimal Class

Action を書く必要がない単純なクラスでは、次の形まで縮められます。

```text
{
  Responsibility: 
  Fields: [
    fieldName: 
  ]
}
```

## Notes

- 空欄の任意項目は削除してください。
- Action は実装の逐語訳ではなく、処理の目的と流れを書いてください。
- テンプレートを埋めること自体を目的にせず、コード理解に必要な情報だけを残してください。
