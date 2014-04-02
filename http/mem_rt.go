package http

import (
	"github.com/ulricqin/falcon/collector"
	"net/http"
)

func CfgMemRouter() {
	m.Get("/proc/mem", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		mem, err := collector.MemInfo()
		if err != nil {
			return RenderErrDto(err.Error())
		}
		memFree := mem.MemFree + mem.Buffers + mem.Cached
		memUsed := mem.MemTotal - memFree
		return RenderDataDto([]interface{}{mem.MemTotal / 1024, memUsed / 1024, memFree / 1024})
	})

}
