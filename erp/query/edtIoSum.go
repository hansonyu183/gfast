package query

import (
	"fmt"
	"gfast/erp/boot"

	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// EdtIoSum API管理对象
var EdtIoSum = &edtIoSum{
	"que_edt_io_sum",
}

type edtIoSum struct {
	name string
}

type edtIoSumParam struct {
	BegDate string
	EndDate string
	ResId   string
	EmpId   string
	UnitId  string
	VrType  string
}

//controller
func (ctrl *edtIoSum) Get(r *ghttp.Request) {
	var para edtIoSumParam
	if err := r.Parse(&para); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}

	data, err := ctrl.getData(para)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/

func (ctrl *edtIoSum) getData(params edtIoSumParam) (data interface{}, err error) {
	sql := "call " + ctrl.name + ctrl.paramsSQL(params)

	data, err = boot.ErpDB.GetOne(sql)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	return
}

func (ctrl *edtIoSum) paramsSQL(params edtIoSumParam) (sql string) {
	sql = fmt.Sprintf(
		"('%s','%s','%s','%s','%s','%s')",
		params.BegDate, params.EndDate, params.ResId, params.EmpId, params.UnitId, params.VrType)
	return
}
