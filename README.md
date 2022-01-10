# log-survey_Q1

## Q1実装方針
[実装方針](https://github.com/imokenpi2011/log-survey/wiki/10_Q1-%E5%AE%9F%E8%A3%85%E6%96%B9%E9%87%9D)を参照

## 概要
- 特定のサーバーに向けてpingしたログを読み込む想定のプログラム。
- ログは以下の形式で記載されている。
`[yyyymmddHHMM形式の日時, IPアドレス, レスポンスタイム]`

## 内容
### パッケージ構造(共通)
```
../log-survey
├── README.md
├── app
│   ├── model
│   │   └── timeoutServer.go    // CSV出力に必要な情報を記載
│   ├── report
│   │   └── report.go   // CSV出力の部分を担当
│   └── serverDetail
│       ├── logic.go    // メインのロジックを担当
│       └── serverDetail.go // メインロジックで使用するサブメソッドの定義
├── config
│   └── config.go   // 設定ファイルの読み込み
├── config.ini  // 設定ファイル
├── go.mod
├── go.sum
├── main.go // 実行ファイル。メインロジックとCSV出力を呼び出す
└── server.log  // 読み込みファイル
```
### 実行方法
1. `$git clone https://github.com/imokenpi2011/log-survey`を実行し、`cd log-survey`を実行する。
2. Q1の場合は`$git checkout Q1`の様にしてチャックアウトする。
3. 設定ファイルに読み込みファイルと出力形式を定義
```
[file]
inputfile = ./server.log    // 読み込むファイルの場所(main.goの相対パス又は絶対パス)
outputfile = report_[date].csv     // 出力ファイル名の定義。[Qx]ブランチでは[date]を含めるとファイル名にyyyymmddHHMM形式の日付を含ませることができる。必ずcsv形式で指定すること。
```
4. `go run main.go`を実行。
5. その他期待値はmainブランチのREADME.mdを参考にしてください。