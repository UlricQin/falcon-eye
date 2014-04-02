package http

import (
	"fmt"
	"github.com/ulricqin/falcon-eye/funcs"
	"github.com/ulricqin/goutils/formatter"
	"net/http"
)

func CfgNetRouter() {

	m.Get("/proc/net/rate", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		ret := [][]string{}
		for k, _ := range funcs.NetIfList {
			item := []string{
				k,
				fmt.Sprintf("%s/s", formatter.DisplaySize(float64(funcs.NetReceiveBytesRate(k)))),
				fmt.Sprintf("%s/s", formatter.DisplaySize(float64(funcs.NetTransmitBytesRate(k)))),
				fmt.Sprintf("%s/s", formatter.DisplaySize(float64(funcs.NetTotalBytesRate(k)))),
				fmt.Sprintf("%d/s", funcs.NetDroppedRate(k)),
			}
			ret = append(ret, item)
		}

		return RenderDataDto(ret)
	})

}
