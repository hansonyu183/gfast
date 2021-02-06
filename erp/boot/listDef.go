package boot

import (
	"sync"

	"github.com/gogf/gf/frame/g"
)

type List struct {
	Name   string
	Select string
	From   string
	Like   string
	Del    string
}

var listMap map[string]List

var once sync.Once

func ListMap() map[string]List {
	once.Do(func() {
		ReLoadListMap()
	})
	return listMap
}

func ReLoadListMap() {
	listMap = make(map[string]List)
	r, err := g.DB("erp").Table("app_def_list").All()
	if err != nil {
		return
	}
	for _, v := range r {
		var ld List
		if err1 := v.Struct(&ld); err1 != nil {
			return
		}

		listMap[ld.Name] = ld
	}
}

func init() {
	ReLoadListMap()
}
