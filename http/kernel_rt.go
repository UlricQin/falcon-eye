package http

import (
	"github.com/ulricqin/falcon/collector"
	"github.com/ulricqin/goutils/systool"
	"net/http"
)

func CfgKernelRouter() {
	m.Get("/proc/kernel/hostname", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		hostname, err := collector.KernelHostname()
		if err != nil {
			return RenderErrDto(err.Error())
		}
		return RenderDataDto(hostname)
	})

	m.Get("/proc/kernel/maxproc", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		maxProc, err := collector.KernelMaxProc()
		if err != nil {
			return RenderErrDto(err.Error())
		}
		return RenderDataDto(maxProc)
	})

	m.Get("/proc/kernel/maxfiles", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		maxFiles, err := collector.KernelMaxFiles()
		if err != nil {
			return RenderErrDto(err.Error())
		}
		return RenderDataDto(maxFiles)
	})

	m.Get("/proc/kernel/version", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		ver, err := systool.CmdOutNoLn("uname", "-r")
		if err != nil {
			return RenderErrDto(err.Error())
		}
		return RenderDataDto(ver)
	})
}
