package restful

import (
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type RfPost struct {
	PostMod  *gdb.Model
	PostData interface{}
}

//controller
func (rf *RfPost) Post(r *ghttp.Request) {
	if err := r.Parse(rf.PostData); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}

	if rep, err := Create(rf.PostMod, rf.PostData); err != nil {
		response.FailJson(true, r, err.Error())
	} else {
		if rep != nil {
			response.SusJson(true, r, "成功", rep)
		} else {
			response.SusJson(true, r, "成功", g.Slice{})
		}
	}

}

//service
func Create(mod *gdb.Model, data interface{}) (rep interface{}, err error) {
	m := mod.Clone()
	if rep, err = m.Data(data).Insert(); err != nil {
		return nil, err
	} else {
		return rep, nil
	}
}
