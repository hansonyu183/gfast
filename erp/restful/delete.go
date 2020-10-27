package restful

import (
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type RfDelete struct {
	DeleteMod *gdb.Model
}

//controller
func (rf *RfDelete) Delete(r *ghttp.Request) {
	keys := r.Get("array")
	if list, err := rf.del(keys); err != nil {
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
func (po *RfDelete) del(keys interface{}) (rep interface{}, err error) {
	m := po.DeleteMod.Clone().WherePri(keys)
	if rep, err = m.Delete(); err != nil {
		return nil, err
	} else {
		return rep, nil
	}
}
