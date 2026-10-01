# JSON-like Comment Outs Templates

JSON-like Comment Out は**クラス・関数・メソッドなどの宣言前**に使用します。

実装内部の処理単位や変数説明には通常コメントを使用します。

> 項目名・説明文は、そのコードを読むチームが最も読みやすい言語で書くことを推奨します。

## 日本語テンプレート

### Class

```text
{
  責務:
  フィールド: [
    fieldName:
  ]
  処理: [
    1:
  ]
}
```

### Function / Method

```text
{
  責務:
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

### Full

```text
{
  責務:
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
  副作用: [

  ]
  エラー: [
    ErrorName:
  ]
  補足: [

  ]
}
```

## English Template

英語を主に使用するチームでは、例えば次のように書けます。

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

英語キーは必須ではありません。

## TypeScript / JavaScript Example

```ts
/*
{
  責務: ユーザーIDからユーザー情報を取得する
  処理: [
    1: userIdの妥当性を確認する
    2: Repositoryから対象ユーザーを取得する
    3: 取得結果をAPI返却形式へ変換する
    4: 呼び出し元へ返す
  ]
  引数: [
    userId: 取得対象ユーザーのID
  ]
  戻り値: [
    user: API返却形式のユーザー情報
  ]
  エラー: [
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
  責務: ユーザー認証の状態と認証処理を管理する
  フィールド: [
    token: 現在利用している認証トークン
    user: ログイン中のユーザー情報
    isAuthenticated: 現在認証済みかどうか
  ]
  処理: [
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
#   責務: ユーザーIDからユーザー情報を取得する
#   処理: [
#     1: user_idの妥当性を確認する
#     2: Repositoryから対象ユーザーを取得する
#     3: 呼び出し元へ返す
#   ]
#   引数: [
#     user_id: 取得対象ユーザーのID
#   ]
#   戻り値: [
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

日本語:

```text
引数: []
戻り値: []
```

英語:

```text
Param: []
Return: []
```

## Guideline

```text
Declaration
└─ JSON-like Comment Out
   └─ チームが読める言語で責務・処理・入出力・状態を構造化する

Implementation
└─ Normal Comments
   └─ 処理単位・理由・非自明な変数を自然文で説明する
```

JSON-like にすることや英語キーを使うこと自体を目的にせず、**読みやすく、書きやすく、保守しやすいこと**を優先してください。
