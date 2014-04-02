package funcs

import (
	"github.com/ulricqin/falcon-eye/global"
	"github.com/ulricqin/falcon/collector"
)

var CpuNum int
var CpuSnapshootList [global.MaxCpustatHistory]*collector.CpuSnapshoot

func UpdateCpuStat() error {
	if CpuNum == 0 {
		CpuNum = collector.CpuNum()
	}

	snap, err := collector.CpuSnapShoot()
	if err != nil {
		return err
	}

	length := global.MaxCpustatHistory
	for i := length - 1; i > 0; i-- {
		CpuSnapshootList[i] = CpuSnapshootList[i-1]
	}

	CpuSnapshootList[0] = snap
	return nil
}

func CpuIdle() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].Idle-CpuSnapshootList[1].Idle) * invQuotient
}

func CpuBusy() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	dI := float64(CpuSnapshootList[0].Idle - CpuSnapshootList[1].Idle)
	return (dTot - dI) * 100 / dTot
}

func CpuUser() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].User-CpuSnapshootList[1].User) * invQuotient
}

func CpuNice() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].Nice-CpuSnapshootList[1].Nice) * invQuotient
}

func CpuSystem() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].System-CpuSnapshootList[1].System) * invQuotient
}

func CpuIowait() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].Iowait-CpuSnapshootList[1].Iowait) * invQuotient
}

func CpuIrq() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].Irq-CpuSnapshootList[1].Irq) * invQuotient
}

func CpuSoftIrq() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].SoftIrq-CpuSnapshootList[1].SoftIrq) * invQuotient
}

func CpuSteal() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].Steal-CpuSnapshootList[1].Steal) * invQuotient
}

func CpuGuest() float64 {
	dTot := float64(CpuSnapshootList[0].Total - CpuSnapshootList[1].Total)
	invQuotient := 100.00 / dTot
	return float64(CpuSnapshootList[0].Guest-CpuSnapshootList[1].Guest) * invQuotient
}

func CpuCnt() int {
	return CpuNum
}
