package http

import (
	"fmt"
	"github.com/ulricqin/falcon-eye/funcs"
	"net/http"
)

/*
Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
sda               0.00     0.00    0.00    0.00     0.00     0.00     0.00     0.00    0.00    0.00    0.00   0.00   0.00
*/

func CfgIORouter() {

	m.Get("/proc/io", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		ret := [][]string{}
		for k, _ := range funcs.DiskStatsList {
			rio := funcs.IODelta(k, funcs.IOReadRequests)
			wio := funcs.IODelta(k, funcs.IOWriteRequests)
			delta_rsec := funcs.IODelta(k, funcs.IOReadSectors)
			delta_wsec := funcs.IODelta(k, funcs.IOWriteSectors)
			ruse := funcs.IODelta(k, funcs.IOMsecRead)
			wuse := funcs.IODelta(k, funcs.IOMsecWrite)
			use := funcs.IODelta(k, funcs.IOMsecTotal)
			n_io := rio + wio
			avgrq_sz := 0.0
			await := 0.0
			svctm := 0.0
			if n_io != 0 {
				avgrq_sz = float64(delta_rsec+delta_wsec) / float64(n_io)
				await = float64(ruse+wuse) / float64(n_io)
				svctm = float64(use) / float64(n_io)
			}

			item := []string{
				k,
				fmt.Sprintf("%d", funcs.IODelta(k, funcs.IOReadMerged)),
				fmt.Sprintf("%d", funcs.IODelta(k, funcs.IOWriteMerged)),
				fmt.Sprintf("%d", rio),
				fmt.Sprintf("%d", wio),
				fmt.Sprintf("%.2f", float64(delta_rsec)/2.0),
				fmt.Sprintf("%.2f", float64(delta_wsec)/2.0),
				fmt.Sprintf("%.2f", avgrq_sz),                                                    // avgrq-sz: delta(rsect+wsect)/delta(rio+wio)
				fmt.Sprintf("%.2f", float64(funcs.IODelta(k, funcs.IOMsecWeightedTotal))/1000.0), // avgqu-sz: delta(aveq)/s/1000
				fmt.Sprintf("%.2f", await),                                                       // await: delta(ruse+wuse)/delta(rio+wio)
				fmt.Sprintf("%.2f", svctm),                                                       // svctm: delta(use)/delta(rio+wio)
				fmt.Sprintf("%.2f%%", float64(use)/10.0),                                         // %util: delta(use)/s/1000 * 100%
			}
			ret = append(ret, item)
		}

		return RenderDataDto(ret)
	})

}
