# XML ↔ JSON-like Converter

**Status: Planned — まだ実装されていません。**

XML Documentation コメントと JSON-like Comment Outs を相互変換するツールです。

## やりたいこと

主に C# の XML Documentation を対象として、

```xml
/// <summary>ユーザーを取得する</summary>
/// <param name="userId">対象ユーザーID</param>
/// <returns>ユーザー情報</returns>
```

から、

```text
{
  責務: ユーザーを取得する
  引数: [
    userId: 対象ユーザーID
  ]
  戻り値: [
    user: ユーザー情報
  ]
}
```

のような変換を行います。

逆方向にも変換できるようにし、

- 普段は JSON-like Comment Outs で書く
- 必要な場合だけ XML Documentation を生成する

という使い方も想定します。

## 想定する対応

- summary ↔ Responsibility
- param ↔ Param
- returns ↔ Return
- exception ↔ Error
- remarks / note 系 ↔ Note

完全に1対1対応できない要素をどう扱うかは今後決めます。

## 今後決めること

- C# 以外の XML Documentation 形式を扱うか
- Action を XML 側でどう表現するか
- see / seealso / cref の扱い
- inheritdoc の扱い
- 変換不能な情報の保持方法
- ファイル単位 / プロジェクト単位の変換方法
