package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

// 設定一覧を指定
type ConfigList struct {
	InputFile  string // 読み込みファイル
	OutputFile string // 出力ファイル
	ArgCount   int    // 許可する引数の数
}

//Configインスタンスを定義
var Config ConfigList

//初期処理
func init() {
	//設定の読み込み処理
	LoadConfig()
}

//設定を読み込む
func LoadConfig() {
	// 設定を記載したiniファイルを読み込む
	configDir := "config.ini"
	//root直下の設定ファイルを読み込む
	cfg, err := ini.Load(configDir)
	if err != nil {
		log.Fatalln(err)
	}

	//読み込んだ値を設定する
	argCount, _ := cfg.Section("args").Key("argcount").Int()
	Config = ConfigList{
		InputFile:  cfg.Section("file").Key("inputfile").String(),
		OutputFile: cfg.Section("file").Key("outputfile").String(),
		ArgCount:   argCount,
	}
}
