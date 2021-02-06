package query

import (
	"fmt"
	"gfast/erp/boot"
	"strings"

	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// EdtIoGroupSum API管理对象
var EdtIoGroupSum = &edtIoGroupSum{
	"que_edt_io_group_sum",
}

type edtIoGroupSum struct {
	name string
}

type edtIoGroupSumParam struct {
	BegDate string
	EndDate string
	ResId   string
	EmpId   string
	UnitId  string
	VrType  string
	GroupBy []string
}

//controller
func (ctrl *edtIoGroupSum) Get(r *ghttp.Request) {
	var para edtIoGroupSumParam
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

func (ctrl *edtIoGroupSum) getData(params edtIoGroupSumParam) (data interface{}, err error) {
	sql := "call " + ctrl.name + ctrl.paramsSQL(params)

	data, err = boot.ErpDB.GetOne(sql)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	return
}

func (ctrl *edtIoGroupSum) paramsSQL(params edtIoGroupSumParam) (sql string) {
	grp := strings.Join(params.GroupBy, ",")
	sql = fmt.Sprintf(
		"('%s','%s','%s','%s','%s','%s','%s')",
		params.BegDate, params.EndDate, params.ResId, params.EmpId, params.UnitId, params.VrType,
		grp)
	return
}
