package funcs

import (
	"github.com/ulricqin/falcon/collector"
)

var NetIfList = make(map[string][2]*collector.NetIf)

func UpdateIfStat() error {
	netIfs, err := collector.NetIfs()
	if err != nil {
		return err
	}

	for i := 0; i < len(netIfs); i++ {
		iface := netIfs[i].Iface
		NetIfList[iface] = [2]*collector.NetIf{netIfs[i], NetIfList[iface][0]}
	}
	return nil
}

func NetReceiveBytesRate(iface string) int64 {
	val, ok := NetIfList[iface]
	if !ok {
		return 0
	}

	if val[1] == nil {
		return 0
	}

	return val[0].InBytes - val[1].InBytes
}

func NetTransmitBytesRate(iface string) int64 {
	val, ok := NetIfList[iface]
	if !ok {
		return 0
	}

	if val[1] == nil {
		return 0
	}

	return val[0].OutBytes - val[1].OutBytes
}

func NetTotalBytesRate(iface string) int64 {
	val, ok := NetIfList[iface]
	if !ok {
		return 0
	}

	if val[1] == nil {
		return 0
	}

	return val[0].TotalBytes - val[1].TotalBytes
}

func NetDroppedRate(iface string) int64 {
	val, ok := NetIfList[iface]
	if !ok {
		return 0
	}

	if val[1] == nil {
		return 0
	}

	return val[0].TotalDropped - val[1].TotalDropped
}
