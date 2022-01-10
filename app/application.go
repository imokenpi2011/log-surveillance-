package app

import (
	"bufio"
	"fmt"
	"log-survey/app/model"
	"log-survey/app/report"
	"log-survey/config"
	"os"
	"strings"
)

// タイムアウトしたサーバーの詳細一覧
var TimeoutServerDetail = []*model.TimeoutServer{}

// タイムアウト中のIP一覧
var TimeoutServerIpList = []string{}

func ReadFile() {

	// ファイルの読み込み
	file, err := os.Open(config.Config.InputFile)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(rowSlice)

		// ping応答がないログかどうか検証(応答時間が「-」)
		if rowSlice[2] == "-" {
			fmt.Println("ping timeout.")
			// ping応答がない場合
			// 応答がないサーバー一覧に追加されていない場合のみ処理を行う
			if !IsIpContains(rowSlice[1]) {
				// 応答がないサーバー一覧に追加
				TimeoutServerIpList = append(TimeoutServerIpList, rowSlice[1])
				// 応答がないサーバー詳細にタイムアウト開始時刻、IPを登録
				registTimeoutServerDetail(rowSlice)
				fmt.Println(TimeoutServerIpList)
				fmt.Println(TimeoutServerDetail)
			}
		} else {
			// ping応答がある場合
			// 応答がないサーバー一覧にある場合、復帰した事にする
			if IsIpContains(rowSlice[1]) {
				fmt.Println("recovered.")
				// 応答がないサーバー一覧から削除
				TimeoutServerIpList = DeleteRecoveredIp(rowSlice[1])
				// 応答がないサーバー詳細に復帰時刻を登録
				registRecoverServerDetail(rowSlice)
				fmt.Println(TimeoutServerIpList)
				fmt.Println(TimeoutServerDetail)
			}
		}
	}

	// csv出力
	report.OutputCsv(TimeoutServerDetail)
}

// ping応答がないサーバー一覧に含まれているか検証
func IsIpContains(serverIp string) bool {
	if serverIp == "" {
		return false
	}

	// ping応答がないサーバーが含まれるか精査
	for _, timeoutServerIp := range TimeoutServerIpList {
		// 含まれている場合はtrueを返す
		if serverIp == timeoutServerIp {
			return true
		}
	}

	// 含まれていない場合はfalse
	return false
}

// 復帰したサーバーIPを一覧から削除
func DeleteRecoveredIp(serverIp string) (serverList []string) {
	if serverIp == "" {
		return
	}

	// 削除対象のIPのみを抜いた新しい配列を作成
	for _, timeoutServerIp := range TimeoutServerIpList {
		if serverIp != timeoutServerIp {
			serverList = append(serverList, timeoutServerIp)
		}
	}

	return serverList
}

// タイムアウトしたログの時間、IP情報を書き込む
func registTimeoutServerDetail(timeoutServer []string) {
	if len(timeoutServer) != 3 {
		return
	}

	// タイムアウトしたサーバーの詳細にIPとタイムアウト時刻を書き込む
	TimeoutServerDetail = append(TimeoutServerDetail, &model.TimeoutServer{
		Ip:           timeoutServer[1],
		TimeoutStart: timeoutServer[0],
	})
}

// 復帰したサーバーのログの時間を書き込む
func registRecoverServerDetail(recoveredServer []string) {
	if len(recoveredServer) != 3 {
		return
	}

	// タイムアウトしたサーバーの詳細一覧を検索し、該当のIPのデータに復帰時間を書き込む
	for _, serverDetail := range TimeoutServerDetail {
		if recoveredServer[1] == serverDetail.Ip {
			fmt.Println("matched.")
			// 一度復帰したログには書き込まない
			if serverDetail.TimeoutEnd == "" {
				serverDetail.TimeoutEnd = recoveredServer[0]
			}
		}
	}
}
