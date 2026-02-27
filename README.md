# aozorandom

青空文庫の作家名を入力すると、その作家のランダムな著作を返すツールです。

## 機能

- 作家名（部分一致）で著作を検索
- 該当する著作の中からランダムに1作品を選んで表示
- 作品名と出版社を出力

## 必要環境

- Go 1.26以上

## インストール

```bash
git clone https://github.com/hrn-tmy/aozorandom.git
cd aozorandom
go mod tidy
go build -o aozora .
```

## 使い方

```bash
./aozora <作家名>
```

### 実行例

```bash
$ ./aozora 夏目漱石
作品: こころ
出版社: 岩波書店

$ ./aozora 芥川
作品: 羅生門
出版社: 岩波書店
```

作家名は部分一致で検索されます。例えば `芥川` と入力すると `芥川龍之介` の作品が対象になります。

## 依存ライブラリ

```bash
go get golang.org/x/text
```

| パッケージ                            | 用途                                 |
| ------------------------------------- | ------------------------------------ |
| `golang.org/x/text/encoding/japanese` | Shift_JIS → UTF-8 変換               |
| `golang.org/x/text/transform`         | エンコーディング変換のストリーム処理 |

## データソース

[青空文庫](https://www.aozora.gr.jp/) が公開している作品リストCSV（Shift_JIS形式）を使用しています。

- 取得URL: `https://www.aozora.gr.jp/index_pages/list_person_all.zip`

ツール実行のたびに最新のリストを取得するため、常に最新の作品情報が反映されます。
