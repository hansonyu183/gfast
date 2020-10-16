package restful

import (
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type Post struct {
	Params interface{}
	Mod    *gdb.Model
}

//controller
func (po *Post) Post(r *ghttp.Request) {
	if err := r.Parse(po.Params); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}

	if list, err := po.create(); err != nil {
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
func (po *Post) create() (data interface{}, err error) {
	m := po.Mod.Clone()
	if list, err := m.Data(po.Params).Insert(); err != nil {
		return nil, err
	} else {
		return list, nil
	}
}
