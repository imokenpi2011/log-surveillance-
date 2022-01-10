package serverDetail

import (
	"log-survey/app/model"
)

// ping応答がないサーバー一覧に含まれているか検証
func isIpContains(serverIp string, timeoutServerIpList []string) bool {
	if serverIp == "" {
		return false
	}

	// ping応答がないサーバーが含まれるか精査
	for _, timeoutServerIp := range timeoutServerIpList {
		// 含まれている場合はtrueを返す
		if serverIp == timeoutServerIp {
			return true
		}
	}

	// 含まれていない場合はfalse
	return false
}

// 復帰したサーバーIPを一覧から削除
func deleteRecoveredIp(serverIp string, timeoutServerIpList []string) (serverList []string) {
	if serverIp == "" {
		return
	}

	// 削除対象のIPのみを抜いた新しい配列を作成
	for _, timeoutServerIp := range timeoutServerIpList {
		if serverIp != timeoutServerIp {
			serverList = append(serverList, timeoutServerIp)
		}
	}

	return serverList
}

// タイムアウトしたログの時間、IP情報を書き込む
func registTimeoutServerDetail(timeoutServer []string, timeoutServerDetail []*model.TimeoutServer) []*model.TimeoutServer {
	// フォーマットg違う場合は何もしない
	if len(timeoutServer) != 3 {
		return timeoutServerDetail
	}

	// タイムアウトしたサーバーの詳細にIPとタイムアウト時刻,連続タイムアウト回数を書き込む
	timeoutServerDetail = append(timeoutServerDetail, &model.TimeoutServer{
		Ip:           timeoutServer[1],
		TimeoutStart: timeoutServer[0],
		TimeoutCount: 1,
	})

	return timeoutServerDetail
}

// 復帰したサーバーのログの時間を書き込む
func registRecoverServerDetail(recoveredServer []string, timeoutServerDetail []*model.TimeoutServer) []*model.TimeoutServer {
	// フォーマットが違う場合は何もしない
	if len(recoveredServer) != 3 {
		return timeoutServerDetail
	}

	// タイムアウトしたサーバーの詳細一覧を検索し、該当のIPのデータに復帰時間を書き込む
	for _, serverDetail := range timeoutServerDetail {
		if recoveredServer[1] == serverDetail.Ip {
			// 一度復帰したログには書き込まない
			if serverDetail.TimeoutEnd == "" {
				serverDetail.TimeoutEnd = recoveredServer[0]
			}
		}
	}

	return timeoutServerDetail
}

// 連続タイムアウト回数のカウントアップを実行
func countupTimeoutCount(targetServerIp string, timeoutServerDetail []*model.TimeoutServer) []*model.TimeoutServer {
	// 文字列がない場合は何もしない
	if targetServerIp == "" {
		return timeoutServerDetail
	}

	// タイムアウトしたサーバーの詳細一覧を検索し、該当のIPのデータのタイムアウト回数を増やす
	for _, serverDetail := range timeoutServerDetail {
		if targetServerIp == serverDetail.Ip {
			// 一度復帰したログには書き込まない
			if serverDetail.TimeoutEnd == "" {
				// カウントアップ
				serverDetail.TimeoutCount++
			}
		}
	}

	return timeoutServerDetail
}
