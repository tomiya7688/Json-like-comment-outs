# 責務表生成ツール 仕様

Status: **Design Draft**

> この日本語版を正本 (canonical) とします。  
> 現時点では設計段階であり、実装は行いません。

## 1. 目的

JSON-like Comment Outs に記述された責務を抽出し、コードベース全体を俯瞰できる Markdown の責務表を生成します。

基本出力:

```markdown
| クラス名 | 責務 |
| --- | --- |
| UserService | ユーザー情報を管理する |
| AuthManager | 認証状態と認証処理を管理する |
```

ツールは責務を推測せず、コメントに記述された Semantic Key を読み取ります。

このツールは**コメントだけを解析対象**とし、クラス名・型名なども Responsibility / 責務の中に書かれた名前から取得します。

推奨形:

```text
責務: [
  UserService: ユーザー情報を管理する
]
```

このため、責務表生成のためにソースコードの宣言を解析する必要はありません。

## 2. 対応形態

### GUI

GUI版を用意します。

想定する主な操作:

- 対象プロジェクト / ディレクトリの選択
- 出力先の選択
- 出力のまとめ方の選択
- SemanticKeys 設定
- 生成

マウス操作だけでなく、**Alt 起点のキーボード操作**に対応します。

具体的なアクセラレータキーは実装設計時に決定します。

### CUI

CUI版も用意します。

GUI と CUI は同じ解析・生成ロジックと同じ SemanticKeys 設定を使用します。

具体的なコマンド名・オプション名は未確定です。

## 3. 出力方式

次の方式を選択できるようにします。

### 1ファイルに集約

対象すべてを同じ Markdown ファイルへ出力します。

### 単位ごとに分割

言語の構造に応じて、Namespace / Package / Module などの単位ごとに Markdown を分割します。

どの単位を採用するかは各言語アダプタで定義します。

## 4. 対応言語の優先度

### 優先

1. Go
2. Python
3. C#

### 次点

4. Ruby
5. Rust
6. C++
7. Java

実装順は変更される可能性があります。

## 5. 宣言単位

クラスを持つ言語では、基本的にクラス相当の宣言を責務表の対象とします。

### Go

Go はクラスを持たないため、対象単位はまだ確定していません。

候補:

- struct
- interface
- package
- その他、Goで責務単位として自然な宣言

Goを無理に「クラス相当」へ当てはめず、Goとして読みやすい責務表を設計します。

### Rust

Rust についても struct / enum / trait など、責務表へ含める宣言単位を実装前に確定します。

## 6. SemanticKeys

このツールがコメント内の Key の意味を解決するための設定です。

設定は単純な **canonical key → aliases** の JSON とします。

```json
{
  "SemanticKeys": {
    "Responsibility": ["責務", "役割", "Responsible"],
    "Action": ["処理", "手順"],
    "Fields": ["フィールド", "状態"],
    "Param": ["引数", "入力"],
    "Return": ["戻り値", "出力"],
    "SideEffect": ["副作用"],
    "Error": ["エラー"],
    "Note": ["補足"]
  }
}
```

### ルール

- canonical key と alias は**大文字小文字を区別しない**
- canonical key 自身は alias 配列に書かなくても有効
- alias は翻訳語・厳密な同義語である必要はない
- チームが同じ意味として扱う任意の Key を alias にできる
- 同じ alias を複数の canonical key へ割り当ててはならない
- 未登録 Key に対して類義語推測・曖昧一致は行わない

例えば上記設定では、次をすべて `Responsibility` として扱います。

```text
Responsibility
responsibility
RESPONSIBILITY
責務
役割
Responsible
```

## 7. GUI の SemanticKeys 設定

GUI には**設定ボタン**を用意し、SemanticKeys を編集できるようにします。

イメージ:

```text
Semantic Keys

Responsibility
  責務
  役割
  Responsible

Action
  処理
  手順

[ + Alias ]
```

想定操作:

- canonical key の確認
- alias の追加
- alias の編集
- alias の削除
- 重複 alias の検出

GUIで変更した設定は CUI と共有します。

設定ファイルのファイル名・配置場所・Global / Project の優先規則はまだ確定していません。

## 8. 責務の抽出

責務表生成時は、SemanticKeys により `Responsibility` に正規化された Key の値を使用します。

`Responsibility` の中では、`名前: 責務` の組を読み取ります。

```text
Responsibility: [
  UserService: ユーザー情報を管理する
]
```

この `UserService` を表の名前列、右側の説明を責務列へ出力します。

ツールは原則として:

- コメントにない責務をコードから推測しない
- AIで責務を自動生成しない
- 未知の Key を勝手に Responsibility とみなさない

コメントがない宣言を無視するか、空欄として出力するかは未確定です。

## 9. Markdown

最小形式:

```markdown
| クラス名 | 責務 |
| --- | --- |
| UserService | ユーザー情報を管理する |
```

言語によって「クラス名」の列名が不自然な場合は、`型名`、`宣言名` などへ変更できる設計を検討します。

## 10. 未確定事項

Issue 化の前に、少なくとも次を決めます。

- Go の責務単位
- Rust の責務単位
- コメントがない宣言の扱い
- 出力ファイル命名規則
- Namespace / Package / Module 分割の詳細
- SemanticKeys 設定ファイルの名前と配置場所
- GUI の具体的な Alt アクセラレータ
- CUI のコマンド / オプション名
- Markdown の列名ローカライズ
