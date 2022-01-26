/*
Copyright © 2022 qing.wang@ucloud.cn

*/
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"usmgr/cmd"
	"usmgr/config"
)

func main() {
	mod := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(config.USMgrLogs, mod, config.DefaultFileMod)
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		panic(err)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: file})

	cmd.Execute()
}
