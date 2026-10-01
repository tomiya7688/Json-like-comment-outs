# Comment Search & Index / コメント検索・索引ツール

**Status: Planned — まだ実装されていません。**

JSON-like Comment Outs と通常コメントを検索し、コードベース内の責務や機能を探しやすくするツールです。

## やりたいこと

例えば Responsibility を検索して、

```text
認証
  AuthService
  TokenManager

ユーザー
  UserService
  UserRepository
```

のような一覧を作ったり、キーワードから該当するクラス・関数・コメントを検索できるようにします。

Markdown の索引生成も想定します。

## 方針

- SemanticKeys を利用して構造化コメントを検索する
- 通常コメントも全文検索対象にできるようにする
- 実装内容をAIで推測するのではなく、コメントに書かれた内容を検索する

## 今後決めること

- CUI / GUI
- 全文検索の方式
- Markdown Index の形式
- Namespace / Package / Module ごとの分類
