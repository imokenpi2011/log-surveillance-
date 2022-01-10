package report

import (
	"encoding/csv"
	"log-survey/app/model"
	"log-survey/config"
	"os"
)

func OutputCsv(timeoutServerDetail []*model.TimeoutServer) {
	outputServerDetail := formatServerDetail(timeoutServerDetail)

	if len(outputServerDetail) > 0 {
		file, err := os.OpenFile(config.Config.OutputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		cw := csv.NewWriter(file)
		cw.WriteAll(outputServerDetail)
		defer cw.Flush()
	}

}

// csv出力形式にフォーマットした配列を返す
func formatServerDetail(timeoutServerDetail []*model.TimeoutServer) (formatedServerList [][]string) {
	// 故障したサーバーが存在しない場合は空の配列を返す
	if len(timeoutServerDetail) == 0 {
		return
	}

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
