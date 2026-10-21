# 責務表生成ツール 仕様

Status: **MVP実装済み（Go CUI）**

> この日本語版を正本 (canonical) とします。  
> MVPの仕様と実装範囲を記録します。

MVPの範囲は§10で確定しています。GUI、他言語、分割出力は後続の設計・実装Issueで扱います。
## 1. 目的

JSON-like Comment Outs に記述された責務を抽出し、コードベース全体を俯瞰できる Markdown の責務表を生成します。

基本出力:

```markdown
| 名前 | 責務 |
| --- | --- |
| UserService | ユーザー情報を管理する |
| AuthManager | 認証状態と認証処理を管理する |
```

ツールは責務を推測せず、コメントに記述された Semantic Key を読み取ります。

このツールは**コメントだけを解析対象**とし、クラス名・型名なども Responsibility / 責務の中に書かれた名前から取得します。

推奨形:

```json
{
  "責務": {
    "UserService": "ユーザー情報を管理する"
  }
}
```


## 2. 対応形態

### GUI（MVP後）

GUI版はMVP後の拡張とします。

想定する主な操作:

- 対象プロジェクト / ディレクトリの選択
- 出力先の選択
- 出力のまとめ方の選択
- SemanticKeys 設定
- 生成

マウス操作だけでなく、**Alt 起点のキーボード操作**に対応します。

具体的なAltアクセラレーターはGUI設計Issueで決めます。

### CUI（MVP）

MVPの操作形態はCUIです。対応言語はGoとします。

責務コメントの抽出とMarkdown生成を行います。コマンド名・オプション名は§10で定めます。



## 3. 出力方式

次の方式を選択できるようにします。

### MVP: 1ファイルに集約

対象すべてを同じ Markdown ファイルへ出力します。

既定の出力先は入力ディレクトリ直下の responsibility-table.md とし、CUIの --output で変更できます。

### 後続機能: 単位ごとに分割

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

MVPで対応する言語はGoのみです。その他の言語は別Issueで追加します。

## 5. 宣言単位

MVPではソースコードの宣言を解析しません。責務表の1行は、コメント内のResponsibilityオブジェクトにある名前と責務の組です。

Goのstruct、interface、packageなどの宣言単位はMVPで使いません。Rustを含む他言語の宣言単位は、ソース構造を扱う後続設計で決めます。

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
    "Returns": ["戻り値", "出力"],
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

## 7. GUI の SemanticKeys 設定（MVP後）

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

MVPでは入力ディレクトリ直下の .json-like-comment-outs/semantic-keys.json を既定設定とします。CUIの --semantic-keys で別ファイルを指定できます。Global設定とGUI編集はMVP対象外です。

## 8. 責務の抽出

責務表生成時は、SemanticKeys により `Responsibility` に正規化された Key の値を使用します。

`Responsibility` の中では、名前をキー、責務を値とするオブジェクトを読み取ります。

```json
{
  "Responsibility": {
    "UserService": "ユーザー情報を管理する"
  }
}
```

この `UserService` を表の名前列、右側の説明を責務列へ出力します。

ツールは原則として:

- コメントにない責務をコードから推測しない
- AIで責務を自動生成しない
- 未知の Key を勝手に Responsibility とみなさない

MVPはコメントだけを調べるため、Responsibility項目がない宣言を列挙できません。そのため空行や警告は出さず、Responsibilityオブジェクトのエントリだけを出力します。

## 9. Markdown

MVPでは --locale ja|en を使い、既定は ja とします。見出しは「名前」「責務」または「Name」「Responsibility」とし、責務文は翻訳しません。

最小形式:

```markdown
| 名前 | 責務 |
| --- | --- |
| UserService | ユーザー情報を管理する |
```

言語によって「クラス名」の列名が不自然な場合は、`型名`、`宣言名` などへ変更できる設計を検討します。

## 10. MVPの決定事項

- 対応形態はCUI、対応言語はGoのみです。入力ディレクトリ以下のGoソースを再帰的に読みます。
- コマンド名は responsibility-table とし、必須の --source <dir>、任意の --output <file>、--semantic-keys <file>、--locale ja|en を受け付けます。
- ソース宣言やASTは解析しません。Responsibilityオブジェクトの各名前と責務の組を1行にします。
- Responsibility項目を含まない宣言は列挙できないため、行も警告も出しません。
- 出力は1ファイルに集約し、既定の出力先は入力ディレクトリ直下の responsibility-table.md です。パッケージなどによる分割は後続機能です。
- SemanticKeys設定の既定位置は入力ディレクトリ直下の .json-like-comment-outs/semantic-keys.json です。明示パスがあればそちらを使い、設定ファイルがなければ本仕様の既定別名を使います。
- 見出しは --locale ja|en で切り替え、既定は ja です。GUI、Altアクセラレーター、Go以外の言語、Rustなどの宣言単位は後続Issueに延期します。
