package funcs

import (
	"github.com/ulricqin/falcon/collector"
)

var DiskStatsList = make(map[string][2]*collector.DiskStats)

func UpdateDiskStats() error {
	dsList, err := collector.ListDiskStats()
	if err != nil {
		return err
	}

	for i := 0; i < len(dsList); i++ {
		device := dsList[i].Device
		DiskStatsList[device] = [2]*collector.DiskStats{dsList[i], DiskStatsList[device][0]}
	}
	return nil
}

func IOReadRequests(arr [2]*collector.DiskStats) uint64 {
	return arr[0].ReadRequests - arr[1].ReadRequests
}

func IOReadMerged(arr [2]*collector.DiskStats) uint64 {
	return arr[0].ReadMerged - arr[1].ReadMerged
}

func IOReadSectors(arr [2]*collector.DiskStats) uint64 {
	return arr[0].ReadSectors - arr[1].ReadSectors
}

func IOMsecRead(arr [2]*collector.DiskStats) uint64 {
	return arr[0].MsecRead - arr[1].MsecRead
}

func IOWriteRequests(arr [2]*collector.DiskStats) uint64 {
	return arr[0].WriteRequests - arr[1].WriteRequests
}

func IOWriteMerged(arr [2]*collector.DiskStats) uint64 {
	return arr[0].WriteMerged - arr[1].WriteMerged
}

func IOWriteSectors(arr [2]*collector.DiskStats) uint64 {
	return arr[0].WriteSectors - arr[1].WriteSectors
}

func IOMsecWrite(arr [2]*collector.DiskStats) uint64 {
	return arr[0].MsecWrite - arr[1].MsecWrite
}

func IOMsecTotal(arr [2]*collector.DiskStats) uint64 {
	return arr[0].MsecTotal - arr[1].MsecTotal
}

func IOMsecWeightedTotal(arr [2]*collector.DiskStats) uint64 {
	return arr[0].MsecWeightedTotal - arr[1].MsecWeightedTotal
}

func IODelta(device string, f func([2]*collector.DiskStats) uint64) uint64 {
	val, ok := DiskStatsList[device]
	if !ok {
		return 0
	}

	if val[1] == nil {
		return 0
	}
	return f(val)
}
