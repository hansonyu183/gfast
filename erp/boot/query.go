package boot

import (
	"gfast/erp/util/str"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type query struct {
	Fields   string
	Model    string
	Groups   []string
	Orders   string
	Cols     []string
	WhereMap map[string]string
	Params []string
	Proce	string
}

type queryForm struct {
	Groups    []string
	WherePara []string
	Params []string
}

type queryView struct {
	Form *queryForm
	Cols []string
}

type queryTable struct {
	Name     string
	Fields   string
	Model    string
	Groups   string
	Orders   string
	Cols     string
	WhereMap string
	Params string
	Proce	string
}

var QueryMap map[string]query
var QueryViewMap map[string]queryView

func ReLoadQueryMap() {
	QueryMap = make(map[string]query)
	QueryViewMap = make(map[string]queryView)
	r, err := g.DB("erp").Table("app_query").All()
	if err != nil {
		return
	}
	var t queryTable
	for _, v := range r {
		if err1 := v.Struct(&t); err1 != nil {
			return
		}
		groups := strings.Split(t.Groups, ",")
		cols := strings.Split(t.Cols, ",")
		pars:=strings.Split(t.Params, ",")
		QueryMap[t.Name] = query{
			t.Fields,
			t.Model,
			groups,
			t.Orders,
			cols,
			gconv.MapStrStr(t.WhereMap),
			pars,
			t.Proce,
		}
		QueryViewMap[t.Name] = queryView{
			&queryForm{
				groups,
				str.CovMapKeyToArray(t.WhereMap),
				pars,
			},
			cols,
		}
	}
}

func init() {
	ReLoadQueryMap()
}
