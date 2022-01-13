package service

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const (
	dsnTemplate = "%s:%s@tcp(%s)/%s"
	infoSchema  = "information_schema"
)

var Cmd *CmdConfig

type CmdConfig struct {
	Dsn           string `json:"dsn"`
	TableName     string `json:"table_name"`
	Address       string `json:"address"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	DB            string `json:"db"`
	Args          string `json:"args"`
	InfoSchemaDsn string `json:"info_schema_dsn"`
}

func init() {
	ConfigInit()
}

func ConfigInit() {
	Cmd = new(CmdConfig)
	flag.StringVar(&Cmd.Address, "h", "", "")
	flag.StringVar(&Cmd.Username, "u", "", "")
	flag.StringVar(&Cmd.Password, "p", "", "")
	flag.StringVar(&Cmd.DB, "db", "", "")
	flag.StringVar(&Cmd.TableName, "t", "", "")
	flag.StringVar(&Cmd.Args, "a", "charset=utf8mb4&parseTime=True&loc=Local", "")
	flag.StringVar(&Cmd.Dsn, "dsn", "", "")
	flag.Parse()

	if !checkCmdConfig(Cmd) {
		b, err := json.Marshal(Cmd)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Error: args invalid, args: %s\n", b)
		os.Exit(400)
	}
	if Cmd.Dsn == "" {
		Cmd.Dsn = fmt.Sprintf(dsnTemplate, Cmd.Username, Cmd.Password, Cmd.Address, infoSchema)
		if len(Cmd.Args) > 0 {
			Cmd.Dsn += "?" + Cmd.Args
		}
	}
}

func checkCmdConfig(data *CmdConfig) (ok bool) {
	if data.Dsn != "" {
		return true
	}
	if len(data.TableName) <= 0 {
		return
	} else if len(data.DB) <= 0 {
		return
	} else if len(data.Username) <= 0 {
		return
	} else if len(data.Password) <= 0 {
		return
	} else if len(data.Address) <= 0 {
		return
	}
	ok = true
	return
}
