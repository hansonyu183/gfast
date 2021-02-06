package query

import (
	"fmt"
	"gfast/erp/boot"
	"strings"

	"gfast/library/response"
	"gfast/library/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// EdtIoGroup API管理对象
var EdtIoGroup = &edtIoGroup{
	"que_edt_io_group",
}

type edtIoGroup struct {
	name string
}

type edtIoGroupParam struct {
	BegDate  string
	EndDate  string
	ResId    string
	EmpId    string
	UnitId   string
	VrType   string
	GroupBy  []string
	PageNum  int
	PageSize int
}

//controller
func (ctrl *edtIoGroup) Get(r *ghttp.Request) {
	var para edtIoGroupParam
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

func (ctrl *edtIoGroup) getData(params edtIoGroupParam) (data interface{}, err error) {
	pageNum := params.PageNum
	pageSize := params.PageSize

	//分页查询
	if pageNum == 0 {
		pageNum = 1
	}
	page := pageNum

	if pageSize == 0 {
		pageSize = service.AdminPageNum
	}
	params.PageNum = pageSize * (page - 1)
	params.PageSize = pageSize
	sql := "call " + ctrl.name + ctrl.paramsSQL(params)
	fmt.Println(sql)
	r, err := boot.ErpDB.GetAll(sql)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	data = g.Map{
		"data": r,
		"page": page,
	}
	return
}

func (ctrl *edtIoGroup) paramsSQL(params edtIoGroupParam) (sql string) {
	grp := strings.Join(params.GroupBy, ",")
	sql = fmt.Sprintf(
		"('%s','%s','%s','%s','%s','%s','%s',%d,%d)",
		params.BegDate, params.EndDate, params.ResId, params.EmpId, params.UnitId, params.VrType,
		grp, params.PageNum, params.PageSize)
	return
}
