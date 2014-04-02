package http

import (
	"fmt"
	"github.com/ulricqin/falcon/collector"
	"net/http"
)

func CfgSystemRouter() {
	m.Get("/proc/system/date", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		out, err := collector.SystemDate()
		if err != nil {
			return RenderErrDto(err.Error())
		}
		return RenderDataDto(out)
	})

	m.Get("/proc/system/uptime", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		arr, err := collector.SystemUptime()
		if err != nil {
			return RenderErrDto(err.Error())
		}

		return RenderDataDto(fmt.Sprintf("%d days %d hours %d minutes", arr[0], arr[1], arr[2]))
	})

	m.Get("/proc/system/loadavg", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		cpuNum := collector.CpuNum()
		load, err := collector.LoadAvg()
		if err != nil {
			return RenderErrDto(err.Error())
		}

		ret := [3][2]interface{}{
			[2]interface{}{load.Avg1min, int64(load.Avg1min * 100.0 / float64(cpuNum))},
			[2]interface{}{load.Avg5min, int64(load.Avg5min * 100.0 / float64(cpuNum))},
			[2]interface{}{load.Avg15min, int64(load.Avg15min * 100.0 / float64(cpuNum))},
		}

		return RenderDataDto(ret)
	})

}
