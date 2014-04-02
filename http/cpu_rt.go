package http

import (
	"fmt"
	"github.com/ulricqin/falcon-eye/funcs"
	"net/http"
)

func CfgCpuRouter() {
	m.Get("/proc/cpu/num", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		return RenderDataDto(funcs.CpuCnt())
	})

	m.Get("/proc/cpu/usage", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if funcs.CpuSnapshootList[1] == nil {
			return RenderDataDto([][10]string{})
		}

		item := [10]string{
			fmt.Sprintf("%.1f%%", funcs.CpuIdle()),
			fmt.Sprintf("%.1f%%", funcs.CpuBusy()),
			fmt.Sprintf("%.1f%%", funcs.CpuUser()),
			fmt.Sprintf("%.1f%%", funcs.CpuNice()),
			fmt.Sprintf("%.1f%%", funcs.CpuSystem()),
			fmt.Sprintf("%.1f%%", funcs.CpuIowait()),
			fmt.Sprintf("%.1f%%", funcs.CpuIrq()),
			fmt.Sprintf("%.1f%%", funcs.CpuSoftIrq()),
			fmt.Sprintf("%.1f%%", funcs.CpuSteal()),
			fmt.Sprintf("%.1f%%", funcs.CpuGuest()),
		}

		return RenderDataDto([][10]string{item})
	})

}
