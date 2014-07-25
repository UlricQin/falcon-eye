package funcs

import (
	"github.com/ulricqin/falcon/collector"
	"github.com/ulricqin/goutils/slicetool"
)

func PortIsListen(port int64) bool {
	return slicetool.SliceContainsInt64(collector.ListenPorts(), port)
}
