package restful

import (
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type RfPut struct {
	PutMod  *gdb.Model
	PutData interface{}
}

//controller
func (rf *RfPut) Put(r *ghttp.Request) {
	if err := r.Parse(rf.PutData); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	if list, err := Update(rf.PutMod,rf.PutData); err != nil {
		response.FailJson(true, r, err.Error())
	} else {
		if list != nil {
			response.SusJson(true, r, "成功", list)
		} else {
			response.SusJson(true, r, "成功", g.Slice{})
		}
	}
}

//service
func Update(mod *gdb.Model, data interface{}) (rep interface{}, err error) {
	m := mod.Clone().WherePri(data)
	if rep, err = m.Data(data).Update(); err != nil {
		return nil, err
	} else {
		return rep, nil
	}
}
