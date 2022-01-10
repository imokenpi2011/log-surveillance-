package report

import (
	"encoding/csv"
	"log"
	"log-survey/app/model"
	"log-survey/config"
	"os"
)

// csv形式でファイルを出力する
func OutputCsv(timeoutServerDetail []*model.TimeoutServer) {
	// 出力対象のファイルを読み込む
	file, err := os.OpenFile(config.Config.OutputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// CSV書き込み処理
	cw := csv.NewWriter(file)
	defer cw.Flush()

	// タイムアウトしたサーバーがあるか検証
	if len(timeoutServerDetail) > 0 {
		// ある場合はCSV形式にしてフォーマット
		outputServerDetail := formatServerDetail(timeoutServerDetail)
		cw.WriteAll(outputServerDetail)
	} else {
		// 無い場合はメッセージのみ書き込む
		cw.Write([]string{"All server green."})
	}

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
