package boot

import (
	"strings"

	"github.com/gogf/gf/frame/g"
)

type Report struct {
	Name   string
	SQL    string
	Params []string
	Select string
	From string
	Where []string
	Group string
	Having string
	Order string
}

type report struct {
	Name   string
	SQL    string
	Params string
	Select string
	From string
	Where string
	Group string
	Having string
	Order string
}

var ReportMap map[string]Report

func ReLoadReportMap() {
	ReportMap = make(map[string]Report)
	r, err := g.DB("erp").Table("app_def_report").All()
	if err != nil {
		return
	}
	var t report
	for _, v := range r {
		if err1 := v.Struct(&t); err1 != nil {
			return
		}
		//ld.ParamArray = strings.Split(ld.Params, ",")
		ReportMap[t.Name] = Report{
			t.Name,
			t.SQL,
			strings.Split(t.Params, ","),
			t.Select,
			t.From,
			strings.Split(t.Where, ","),
			t.Group,
			t.Having,
			t.Order,
		}
	}
}

func init() {
	ReLoadReportMap()
}
