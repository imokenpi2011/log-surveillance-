# log-survey(ログ監視アプリケーション)

## はじめに
各回答はブランチに対応しています。
例 ) Q1の回答→Q1ブランチ

実装の解説に関しては各ブランチの`README.md`を参照してください。

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
1. `git clone https://github.com/imokenpi2011/log-survey`を実行し、`cd log-survey`を実行する。
2. 設定ファイルに読み込みファイルと出力形式を定義
```
[file]
inputfile = ./server.log    // 読み込むファイルの場所(main.goの相対パス又は絶対パス)
outputfile = report_[date].csv     // 出力ファイル名の定義。[Qx]ブランチでは[date]を含めるとファイル名にyyyymmddHHMM形式の日付を含ませることができる。必ずcsv形式で指定すること。
```
3. `go run main.go`を実行。 ※Q2以降は引数が無いとエラーになります。
4. ログ内に出力対象の行がある場合、下記のログが出ます。
- 2行対象の場合
```
2022/01/10 19:31:47 Timeout server found. Start generating report.
2022/01/10 19:31:47 2 records found.
2022/01/10 19:31:47 Success generating report.
```
5. 指定したディレクトリ内にcsv形式のレポートが生成されます。
- 2の設定通りだと下記の中身の`report_20220110193147.csv`が生成される
```
Server IP,Timeout time,Recovered time
10.20.30.2/16,20201019133325,20201019133525
192.168.1.2/24,20201019133435,20201019133835
```