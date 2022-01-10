package main

import (
	"log-survey/app/report"
	"log-survey/app/serverDetail"
)

func main() {
	// ログからサーバーの復旧時間を取り出す
	timeoutServerDetail := serverDetail.GetTimeoutServerList()

	// csv出力
	report.OutputCsv(timeoutServerDetail)
}
