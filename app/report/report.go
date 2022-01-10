package report

import (
	"encoding/csv"
	"log"
	"log-survey/app/model"
	"log-survey/config"
	"os"
	"strings"
	"time"

	"github.com/leekchan/timeutil"
)

// csv形式でファイルを出力する
func OutputCsv(timeoutServerDetail []*model.TimeoutServer) {
	// タイムアウトしたサーバーがあるか検証
	if len(timeoutServerDetail) > 0 {
		log.Printf("Timeout server found. Start generating report.")
		log.Printf("%d records found.\n", len(timeoutServerDetail))

		// []ファイル名を生成する
		fileName := strings.Replace(config.Config.OutputFile, "[date]", getDateTimeString(), 1)
		// 出力対象のファイルを読み込む
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		// ある場合はCSV形式にしてフォーマット
		outputServerDetail := formatServerDetail(timeoutServerDetail)
		// CSV書き込み処理
		cw := csv.NewWriter(file)
		defer cw.Flush()
		cw.WriteAll(outputServerDetail)
		log.Println("Success generating report.")

		return
	}
	log.Printf("Timeout server not found. All green.")
}

// csv出力形式にフォーマットした配列を返す
func formatServerDetail(timeoutServerDetail []*model.TimeoutServer) (formatedServerList [][]string) {
	// 最初の要素にCSVのヘッダを記載
	csvLabel := []string{"Server IP", "Timeout time", "Recovered time"}
	formatedServerList = append(formatedServerList, csvLabel)

	// 出力要素の作成
	for _, serverDetail := range timeoutServerDetail {
		// サーバーが停止中の場合は現在も停止中であるメッセージを記載
		if serverDetail.TimeoutEnd == "" {
			serverDetail.TimeoutEnd = "Currently not responding."
		}

		// 文字列の配列に変換
		strDetail := []string{serverDetail.Ip, serverDetail.TimeoutStart, serverDetail.TimeoutEnd}
		// 要素の追加
		formatedServerList = append(formatedServerList, strDetail)
	}

	return formatedServerList
}

// ファイル用にyyyymmddHHMMSS形式の文字列を取得する
func getDateTimeString() string {
	n := time.Now()
	curTimeStr := timeutil.Strftime(&n, "%Y%m%d%H%M%S")
	return string(curTimeStr)
}
