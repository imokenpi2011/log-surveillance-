package serverDetail

import (
	"bufio"
	"fmt"
	"log"
	"log-survey/app/model"
	"log-survey/config"
	"os"
	"strings"
)

// タイムアウトしたサーバーの一覧を復旧時間も合わせて出す
func GetTimeoutServerList(minTimeoutCount int) []*model.TimeoutServer {

	// タイムアウトしたサーバーの詳細一覧
	var timeoutServerDetail = []*model.TimeoutServer{}

	// タイムアウト中のIP一覧
	var timeoutServerIpList = []string{}

	// ファイルの読み込み
	file, err := os.Open(config.Config.InputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// 一行ずつ読み込む
	serverLog := bufio.NewScanner(file)
	for serverLog.Scan() {
		// テキストを読み込む[日時, IP/サブネットマスク, 応答時間]
		logRow := serverLog.Text()
		rowSlice := strings.Split(logRow, ",")

		// 指定フォーマットになっていない場合はスキップ
		if len(rowSlice) != 3 {
			continue
		}

		// ping応答がないログかどうか検証(応答時間が「-」)
		if rowSlice[2] == "-" {
			// ping応答がない場合
			// 応答がないサーバー一覧に追加されていない場合のみ処理を行う
			if !isIpContains(rowSlice[1], timeoutServerIpList) {
				// 応答がないサーバー一覧に追加
				timeoutServerIpList = append(timeoutServerIpList, rowSlice[1])
				// 応答がないサーバー詳細にタイムアウト開始時刻、IPを登録
				timeoutServerDetail = registTimeoutServerDetail(rowSlice, timeoutServerDetail)
				fmt.Println(timeoutServerDetail)
			}
		} else {
			// ping応答がある場合
			// 応答がないサーバー一覧にある場合、復帰した事にする
			if isIpContains(rowSlice[1], timeoutServerIpList) {
				// 応答がないサーバー一覧から削除
				timeoutServerIpList = deleteRecoveredIp(rowSlice[1], timeoutServerIpList)
				// 応答がないサーバー詳細に復帰時刻を登録
				timeoutServerDetail = registRecoverServerDetail(rowSlice, timeoutServerDetail)
				fmt.Println(timeoutServerDetail)
			}
		}
	}

	return timeoutServerDetail
}
