package vr

import (
	"gfast/erp/boot"

	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// VrList API管理对象
var VrList = &vrList{}

type vrList struct {
	name string
}

type VrListParam struct {
}

//controller
func (ctrl *vrList) Get(r *ghttp.Request) {
	vrType := r.GetString("type")
	data, err := ctrl.getData(vrType)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/

func (ctrl *vrList) getData(vrType string) (data interface{}, err error) {

	mod := boot.ErpDB.Model("vr")

	r, err := mod.All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}

	data = g.Map{
		"tables": g.Map{
			vrType: r,
		},
	}
	return
}
