# Comment Generator / コメント自動挿入ツール

**まだできてないよ...**
**Status: Planned — まだ実装されていません。**

クラス・関数・メソッドの宣言を読み取り、JSON-like Comment Out の雛形を自動挿入するツールです。

## やりたいこと

例えば次の関数がある場合、

```csharp
User GetUser(string userId, bool includeDeleted)
```

次のような雛形を生成します。

```text
{
  責務:
  処理: [
  ]
  引数: [
    userId:
    includeDeleted:
  ]
  戻り値: [
  ]
}
```

クラスなら、実際のフィールド名を読み取って Fields の雛形を生成することも想定します。

## 方針

このツールは**説明文を勝手に考えることを目的にしません。**

まずは、

- 宣言を読む
- 引数名やフィールド名を取得する
- 必要な Semantic Key の枠を作る
- 人間が説明を書く

という補助ツールにします。

将来的に IDE 連携やエディタ拡張を作る可能性もあります。

## 今後決めること

- 挿入対象の判定
- 既存コメントがある場合の動作
- クラス Fields の自動取得範囲
- 戻り値名の扱い
- SemanticKeys に応じたキー名の選択
- GUI / CUI / Editor Extension の優先順位
