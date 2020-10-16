package restful

import (
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type Get struct {
	Params interface{}
	Mod    *gdb.Model
}

//controller
func (rg *Get) Get(r *ghttp.Request) {
	if err := r.Parse(&rg.Params); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}

	if list, err := rg.getList(rg.Params); err != nil {
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
/**
获取列表数据
*/
func (rg *Get) getList() (data interface{}, err error) {
	m := rg.Mod.Clone()
	if list, err := m.Where(rg.Params).All(); err != nil {
		return nil, err
	} else {
		return list, nil
	}
}
