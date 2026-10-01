# JSON-like Comment Outs Templates

JSON-like Comment Out は**クラス・関数・メソッドなどの宣言前**に使用します。

実装内部の処理単位や変数説明には通常コメントを使用します。

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
    3: 取得結果をAPI返却形式へ変換する
    4: 呼び出し元へ返す
  ]
  Param: [
    userId: 取得対象ユーザーのID
  ]
  Return: [
    user: API返却形式のユーザー情報
  ]
  Error: [
    NotFound: 対象ユーザーが存在しない
  ]
}
*/
function getUser(userId) {
  // userIdの妥当性を確認する
  validateUserId(userId);

  // Repositoryから取得した未整形のユーザー情報
  const user = repository.find(userId);

  // API返却形式へ変換する
  const response = toUserResponse(user);

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

```py
# {
#   Responsibility: ユーザーIDからユーザー情報を取得する
#   Action: [
#     1: user_idの妥当性を確認する
#     2: Repositoryから対象ユーザーを取得する
#     3: 呼び出し元へ返す
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

    # Repositoryから取得したユーザー情報
    user = repository.find(user_id)

    return user
```

## Internal Comment Examples

内部コメントには特別な形式を要求しません。

### Processing unit

```ts
// キャッシュに存在しない場合のみRepositoryへ問い合わせる
if (!cachedUser) {
  user = repository.find(userId);
}
```

### Variable

```ts
// 税込み・割引適用後の最終請求額
const total = calculateTotal(order);
```

### No comment needed

```ts
for (let i = 0; i < items.length; i++) {
  // ...
}
```

## Empty Param / Return

```text
Param: []
```

```text
Return: []
```

## Guideline

```text
Declaration
└─ JSON-like Comment Out
   └─ 責務・処理・入出力・状態を構造化して説明する

Implementation
└─ Normal Comments
   └─ 処理単位・理由・非自明な変数を自然文で説明する
```

JSON-like にすること自体を目的にせず、**読みやすく、書きやすく、保守しやすいこと**を優先してください。
