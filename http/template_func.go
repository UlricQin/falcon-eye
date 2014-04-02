package http

import (
	"html/template"
	"strings"
	"time"
)

func nl2br(in string) (out string) {
	out = strings.Replace(in, "\n", "<br>", -1)
	return
}

func htmlQuote(src string) string {
	text := string(src)
	text = strings.Replace(text, "&", "&amp;", -1)
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	text = strings.Replace(text, "'", "&#39;", -1)
	text = strings.Replace(text, "\"", "&quot;", -1)
	text = strings.Replace(text, "“", "&ldquo;", -1)
	text = strings.Replace(text, "”", "&rdquo;", -1)
	text = strings.Replace(text, " ", "&nbsp;", -1)
	return strings.TrimSpace(text)
}

func str2html(raw string) template.HTML {
	return template.HTML(raw)
}

func dateFormat(unixnano int64) string {
	return time.Unix(0, unixnano).Format("2006-01-02 15:04:05")
}
