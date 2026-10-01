# Specification Diff / 仕様差分ツール

**まだできてないよ...**
**Status: Planned — まだ実装されていません。**

2つのリビジョンやブランチのコメントを比較し、**仕様として何が変わったか**を表示するツールです。

## やりたいこと

Git の行差分ではなく、Semantic Key 単位で差分を見せます。

例えば:

```text
UserService

Responsibility
- ユーザー情報を取得する
+ ユーザー情報の取得・更新を担当する

Added Function
+ DeleteUser

Changed Param
UpdateUser
+ force: 強制更新するか
```

## 方針

- コメントの構造を比較する
- 実装コードの差分そのものは対象外
- SemanticKeys で正規化して比較する
- 表記やフォーマットだけの違いは、できる限り仕様差分として扱わない

## 今後決めること

- Git commit / branch / directory のどれを比較対象にするか
- 通常コメントの差分も含めるか
- Markdown / CUI / GUI の出力形式
