package http

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/ulricqin/falcon-eye/funcs"
	"github.com/ulricqin/falcon-eye/global"
	"github.com/ulricqin/goutils/logtool"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Dto struct {
	Succ bool
	Msg  string
	Data interface{}
}

func ErrDto(message string) Dto {
	return Dto{Succ: false, Msg: message}
}

func DataDto(d interface{}) Dto {
	return Dto{Succ: true, Msg: "", Data: d}
}

func RenderErrDto(message string) string {
	dto := ErrDto(message)
	bs, err := json.Marshal(dto)
	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}

func RenderDataDto(d interface{}) string {
	dto := DataDto(d)
	bs, err := json.Marshal(dto)

	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}

var m *martini.ClassicMartini

func StartHttp() {

	p, err := strconv.Atoi(global.HttpPort)
	if err != nil {
		logtool.Fetal("port[%s] format error", global.HttpPort)
		os.Exit(1)
	}

	if funcs.PortIsListen(int64(p)) {
		logtool.Fetal("port[%d] is in listen", p)
		os.Exit(1)
	}

	m = martini.Classic()

	m.Use(render.Renderer(render.Options{
		Funcs: []template.FuncMap{{
			"nl2br":      nl2br,
			"htmlquote":  htmlQuote,
			"str2html":   str2html,
			"dateformat": dateFormat,
		}},
	}))

	m.Get("/healthz", func() string {
		return "ok"
	})

	m.Get("/", func(re render.Render) {
		m := make(map[string]string)
		m["version"] = global.Version
		re.HTML(200, "index", m)
	})

	CfgKernelRouter()
	CfgSystemRouter()
	CfgCpuRouter()
	CfgMemRouter()
	CfgDfRouter()
	CfgNetRouter()
	CfgIORouter()

	logtool.Info("use http port: %s", global.HttpPort)
	http.ListenAndServe(":"+global.HttpPort, m)
}
