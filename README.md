# aozorandom

青空文庫の作家名または著作名を入力すると、その作家のランダムな著作を返すツールです。

## 機能

- 作家名で著作を、または著作名で作家名を検索（部分一致）
- 該当する著作の中からランダムに1作品を選んで表示

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
./aozora <作家名・著作名>

# 例
$ ./aozora 夏目漱石
作家: 夏目 漱石
作品: こころ
出版社: 岩波書店

$ ./aozora 羅生
作家: 芥川 竜之介
作品: 羅生門
出版社: 岩波書店
```

## 依存ライブラリ

| ライブラリ                            | 用途                                 |
| ------------------------------------- | ------------------------------------ |
| `golang.org/x/text/encoding/japanese` | Shift_JIS → UTF-8 変換               |
| `golang.org/x/text/transform`         | エンコーディング変換のストリーム処理 |

## データソース

[青空文庫](https://www.aozora.gr.jp/) が公開している作品リストCSVを使用しています。

- 取得URL: `https://www.aozora.gr.jp/index_pages/list_person_all.zip`

初回実行は最新のリストを取得し、最新の作品情報を反映しつつ7日間キャッシュします。
