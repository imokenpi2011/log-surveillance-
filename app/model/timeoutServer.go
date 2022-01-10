package model

type TimeoutServer struct {
	Ip           string // サーバーIP
	TimeoutStart string // タイムアウト開始時刻
	TimeoutEnd   string // 復帰時刻
	TimeoutCount int    // 連続タイムアウト回数
}
