package funcs

import (
	"github.com/ulricqin/falcon/collector"
	"github.com/ulricqin/goutils/slicetool"
)

func PortIsListen(port int64) bool {
	tcp := collector.ListenPorts("tcp")
	tcp6 := collector.ListenPorts("tcp6")
	udp := collector.ListenPorts("udp")
	udp6 := collector.ListenPorts("udp6")

	if slicetool.SliceContainsInt64(tcp, port) ||
		slicetool.SliceContainsInt64(tcp6, port) ||
		slicetool.SliceContainsInt64(udp, port) ||
		slicetool.SliceContainsInt64(udp6, port) {
		return true
	}
	return false
}
