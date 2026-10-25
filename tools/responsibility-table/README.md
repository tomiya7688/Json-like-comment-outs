# Responsibility Table

JSON-like Comment Outs から責務を抽出し、Markdown の責務表を生成するサポートツールです。

**状態: MVP実装済み（Go CUI）**

## 使い方

Go 1.22 以降が必要です。次のコマンドは対象ディレクトリ以下のGoソースを再帰的に読み取ります。

```text
go run . --source <対象ディレクトリ> [--output <出力ファイル>] [--semantic-keys <設定ファイル>] [--locale ja|en]
```

既定では入力ディレクトリ直下の responsibility-table.md に出力します。SemanticKeys設定は同ディレクトリの .json-like-comment-outs/semantic-keys.json を使い、ファイルがなければ仕様の既定別名を使います。列見出しは日本語が既定です。

コメント内のResponsibilityオブジェクトだけを抽出します。宣言型や所属Packageは解析しません。GUI、分割出力、Go以外の言語は未対応です。

診断は標準エラーへ出力します。引数の不足・不正、--source の型違い、出力先に既存Goソースと同じファイルを指定した場合は終了コード2、ファイル走査・設定読込・出力I/Oのエラーは終了コード1です。コメント内の不正なResponsibility記述やGoソース中の不正な文字は警告として出力し、処理は続行します。

## 開発時の整形

標準の `gofmt` を使います。開発時のGoは `go.mod` の `toolchain go1.27.1` に合わせます。Goファイルの改行は `.gitattributes` でLFに固定します。PowerShellでこのディレクトリから、変更したGoファイルだけを指定してください。引数を省くと、このモジュール直下のGoファイルだけが対象です。vendor・生成物・他ツールのファイルは対象外です。

```powershell
.\format.ps1 semantic_keys.go main_test.go
.\format.ps1 -Check semantic_keys.go main_test.go
```

`-Check` はファイルを書き換えず、未整形なら失敗します。CIもGo 1.27.1でこのモジュールの整形状態を確認します。

## Documentation

### 日本語 — Canonical

- [概要・仕様](./docs/jp/SPECIFICATION.md)

### English

- [Specification](./docs/en/SPECIFICATION.md)
