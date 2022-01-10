package main

import (
	"log"
	"log-survey/app/report"
	"log-survey/app/serverDetail"
	"log-survey/config"
	"os"
	"strconv"
)

func main() {
	// 引数の取得
	minTimeoutCount := getOneLine()
	if minTimeoutCount == 0 {
		return
	}
	// ログからサーバーの復旧時間を取り出す
	timeoutServerDetail := serverDetail.GetTimeoutServerList()

	// csv出力
	report.OutputCsv(timeoutServerDetail, minTimeoutCount)
}

// 引数の取得
func getOneLine() (minTimeoutCount int) {

	// 引数が指定された値と一致しない場合はエラー
	if len(os.Args)-1 != config.Config.ArgCount {
		log.Fatalf("%d argument must be specified.", config.Config.ArgCount)
	}

	// 引数の取得
	for i, v := range os.Args {
		switch i {
		case 0:
			//一番目はスキップする
			continue
		case 1:
			minTimeoutCount, err := strconv.Atoi(v)
			if err != nil {
				log.Println("Argument must be specified number.")
				log.Fatalln(err)
			}
			return minTimeoutCount
		}
	}

	// 返されない数値
	return 0
}
