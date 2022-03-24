package wx

import (
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/hhcool/gtls/log"
	"github.com/hhcool/gtls/structs"
)

func StructToMap(data interface{}) map[string]interface{} {
	return structs.Map(data)
}

// HTTP CLIENT
func debug() gout.DebugFunc {
	return func(o *gout.DebugOption) {
		o.Debug = true
		o.Write = log.SafeWriterLevel(log.Logger, 4)
	}
}
func (w *Wechat) Post(url string) *dataflow.DataFlow {
	if w.debug {
		return gout.POST(url).Debug(debug())
	}
	return gout.POST(url)
}
func (w *Wechat) Get(url string) *dataflow.DataFlow {
	if w.debug {
		return gout.GET(url).Debug(debug())
	}
	return gout.GET(url)
}
