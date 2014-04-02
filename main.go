package main

import (
	"flag"
	"github.com/ulricqin/falcon-eye/funcs"
	"github.com/ulricqin/falcon-eye/global"
	"github.com/ulricqin/falcon-eye/http"
	"github.com/ulricqin/goutils/filetool"
	"github.com/ulricqin/goutils/logtool"
	"time"
)

func main() {
	var cfgFile string
	flag.StringVar(&cfgFile, "c", "", "configuration file")
	flag.Parse()

	if cfgFile == "" {
		cfgFile, _ = filetool.RealPath("cfg.ini")
		logtool.Warn("no configuration file specified. use default: %s", cfgFile)
	}

	global.InitEnv(cfgFile)

	go InitDataHistory()

	http.StartHttp()

	select {}
}

func InitDataHistory() {
	for {
		funcs.UpdateCpuStat()
		funcs.UpdateIfStat()
		funcs.UpdateDiskStats()
		time.Sleep(global.CollBaseInfoInterval)
	}
}
