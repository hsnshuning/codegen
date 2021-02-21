package service

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const (
	dsnTemplate = "root:qwer1234@tcp(127.0.0.1:33061)/demo?charset=utf8mb4&parseTime=True&loc=Local"
)

var Cmd *CmdConfig

type CmdConfig struct {
	Dns       string `json:"dns"`
	TableName string `json:"table_name"`
}

func init() {
	ConfigInit()
}

func ConfigInit() {
	Cmd = new(CmdConfig)
	flag.StringVar(&Cmd.Dns, "dsn", "", "")
	flag.StringVar(&Cmd.TableName, "table_name", "", "")
	flag.Parse()
	if !checkCmdConfig(Cmd) {
		b, err := json.Marshal(Cmd)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Error: args invalid, args: %s\n", b)
		os.Exit(400)
	}
}

func checkCmdConfig(data *CmdConfig) (ok bool) {
	if len(data.TableName) <= 0 {
		return
	} else if len(data.Dns) <= 0 {
		return
	}
	ok = true
	return
}
