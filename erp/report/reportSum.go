package report

import (
	"gfast/erp/boot"
	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// ReportSum API管理对象
var ReportSum = new(reportSum)

type reportSum struct {
}

//controller
func (ctrl *reportSum) Get(r *ghttp.Request) {
	var para ReportParams
	repName := r.GetString("name")
	if err := r.Parse(&para); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	sum, err := ctrl.GetServ(repName, para)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	result := g.Map{
		"sum": sum,
	}

	response.SusJson(true, r, "成功", result)

}

//service
/**
获取列表数据
*/
func (ctrl *reportSum) GetServ(repName string, params ReportParams) (sum interface{}, err error) {
	sql := makeSql(repName, params.FormData)

	if sum, err = boot.ErpDB.GetOne(sql, 0, 0, "sum"); err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
	}
	return
}
