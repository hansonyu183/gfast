package dict

import (
	"gfast/erp/api"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//客户资料管理
type Dict struct {
	*api.Dict
}

func NewDict() (o *Dict) {
	ds := make(map[string]*gdb.Model)
	ds["emp"] = g.DB("erp").Table("emp").Fields("id,name,py")
	ds["eba"] = g.DB("erp").Table("eba").Fields("id,name,py")
	o = &Dict{
		&api.Dict{
			DictMods: ds,
		},
	}
	return
}
