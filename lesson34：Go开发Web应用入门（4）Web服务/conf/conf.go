package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

type tomlConf struct {
	IsDebug bool   `toml:"isDebug"`
	SeHost  string `toml:"SeHost"`
	SePort  string `toml:"SePort"`

	PostgreSql DBConf `toml:"PostgreSqlConf"`
	MySql      DBConf `toml:"MySqlConf"`
}

type DBConf struct {
	DBType   string
	UserName string
	Password string
	DBHost   string
	DBPort   string
	DBName   string
}

//App App的相关配置项
var App *tomlConf

func init() {
	App = new(tomlConf)
}

//Init 初始化配置文件
func Init(filePath string) {
	_, err := toml.DecodeFile(RealFilePath(filePath), App)
	if err != nil {
		fmt.Println("Fail to load Configuration file, try again in 1 minute，error:", err)
		time.Sleep(time.Minute)
		Init(filePath)
	}

}
