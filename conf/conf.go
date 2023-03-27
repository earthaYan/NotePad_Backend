package conf

import (
	"NotePad/model"
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	SERVER_PORT string
	DB_ADDR     string
	DB_TYPE     string
	DB_USER     string
	DB_PASS     string
	DB_NAME     string
)

func Init() {
	cfg, err := ini.Load("./conf/conf.ini")
	if err != nil {
		fmt.Printf("加载配置文件错误：%f", err)
		os.Exit(1)
	}
	loadServer(cfg)
	loadMysql(cfg)
	path := strings.Join([]string{DB_USER, ":", DB_PASS, "@tcp(", DB_ADDR, ")/", DB_NAME, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	model.InitDataBase(path)
}
func loadServer(cfg *ini.File) {
	SERVER_PORT = cfg.Section("server").Key("port").String()
}

func loadMysql(cfg *ini.File) {
	mysql := cfg.Section("mysql")
	DB_ADDR = mysql.Key("db_addr").String()
	DB_TYPE = mysql.Key("db_type").String()
	DB_USER = mysql.Key("db_user").String()
	DB_PASS = mysql.Key("db_pass").String()
	DB_NAME = mysql.Key("db_name").String()
}
