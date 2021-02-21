package codegen

import (
	"flag"
)

var Cmd *CmdConfig

type CmdConfig struct {
	Dns       string `json:"dns"`
	TableName string `json:"table_name"`
}

func ConfigInit() {
	Cmd = new(CmdConfig)
	flag.StringVar(&Cmd.Dns, "dsn", "", "")
	flag.StringVar(&Cmd.TableName, "table_name", "", "")
	flag.Parse()
	if !checkCmdConfig(Cmd){
		panic("args invalid")
	}
}

func checkCmdConfig(data *CmdConfig) (ok bool) {
	if len(data.TableName) <=0 {
		return
	}else if len(data.Dns) <=0 {
		return
	}
	return
}