package global

import (
	"github.com/astaxie/beego/config"
	"github.com/ulricqin/goutils/filetool"
	log "github.com/ulricqin/goutils/logtool"
	"github.com/ulricqin/goutils/systool"
	"os"
	"runtime"
)

var Config config.Configer
var UseSystemConfig bool = false

func InitEnv(configPath string) {

	if !filetool.IsExist(configPath) {
		log.Warn("configuration file[%s] is nonexistent", configPath)
		UseSystemConfig = true
	}

	if !UseSystemConfig {
		var err error
		Config, err = config.NewConfig("ini", configPath)
		if err != nil {
			log.Fetal("configuration file[%s] cannot parse. ", configPath)
			os.Exit(1)
		}
	}

	if !UseSystemConfig {
		log.SetLevelWithDefault(Config.String("log::level"), "info")
	} else {
		log.SetLevel("info")
	}

	initCfg()

	defaultPidPath := "/var/run/falcon_eye/falcon_eye.pid"
	if !UseSystemConfig {
		if iniPidPath := Config.String("common::pid"); iniPidPath != "" {
			defaultPidPath = iniPidPath
		}
	}
	systool.WritePidFile(defaultPidPath)

	runtime.GOMAXPROCS(runtime.NumCPU())
}
