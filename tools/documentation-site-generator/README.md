# Documentation Site Generator / ドキュメントサイト生成ツール

**Status: Planned — まだ実装されていません。**

コメントや、他のToolsが生成した Markdown から HTML ドキュメントサイトを生成するツールです。

## やりたいこと

- クラス説明書をWebで閲覧する
- 責務表をトップページや索引として使う
- クラス / 関数間をリンクする
- Namespace / Package / Module ごとにナビゲーションする

## 方針

まずは Markdown を中間形式として利用する想定です。

```text
Source Comments
      ↓
Class Description / Responsibility Table
      ↓
Markdown
      ↓
Documentation Site
```

## 今後決めること

- 独自HTML生成か既存静的サイトジェネレータ連携か
- リンク解決方法
- テーマ
- 検索
- 公開方法
