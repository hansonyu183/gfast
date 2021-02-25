package desc

import (
	"gfast/erp/boot"

	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// DBDesc API管理对象
var DBDesc = &dbDesc{}

type dbDesc struct {
	name string
}

type DBDescParam struct {
	DBDesc string
}

//controller
func (ctrl *dbDesc) Get(r *ghttp.Request) {
	tbName := r.GetString("table")
	data, err := ctrl.getData(tbName)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/

func (ctrl *dbDesc) getData(tbName string) (data interface{}, err error) {
	mod := boot.ErpDB
	r, err := mod.GetAll("desc " + tbName)

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}

	data = g.Map{
		"data": r,
	}
	return
}
