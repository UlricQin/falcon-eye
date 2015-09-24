package global

import (
	"github.com/ulricqin/goutils/filetool"
	log "github.com/ulricqin/goutils/logtool"
	"os"
	"time"
)

const MaxCpustatHistory = 60

var CollBaseInfoInterval time.Duration
var HttpPort string
var Version string

// configuration
func initCfg() {
	initHttpConfig()
	initCollectBaseInfoInterval()
	initVersion()
}

func initVersion() {
	var err error
	Version, err = filetool.ReadFileToStringNoLn("VERSION")
	if err != nil {
		log.Fetal("read VERSION file fail")
		os.Exit(1)
	}
}

func initHttpConfig() {
	HttpPort = Config.String("port")
	if HttpPort == "" {
		log.Warn("http::port is blank. use default 1988")
		HttpPort = "1988"
	}
}

func initCollectBaseInfoInterval() {
	if UseSystemConfig {
		CollBaseInfoInterval = 1 * time.Second
		return
	}

	str := Config.String("base_interval_in_seconds")
	if str == "" {
		log.Warn("collector::base_interval_in_seconds is blank. use default 1s")
		CollBaseInfoInterval = 1 * time.Second
		return
	}

	v, err := Config.Int64("base_interval_in_seconds")
	if err != nil {
		log.Warn("collector::base_interval_in_seconds config error")
		os.Exit(1)
	}

	CollBaseInfoInterval = time.Duration(v) * time.Second
}
