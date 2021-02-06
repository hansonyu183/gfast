package report

import (
	"gfast/erp/boot"
	"gfast/library/response"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// ReportCount API管理对象
var ReportCount = new(reportCount)

type reportCount struct {
}

//controller
func (ctrl *reportCount) Get(r *ghttp.Request) {
	var para ReportParams
	repName := r.GetString("name")
	if err := r.Parse(&para); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	count, err := ctrl.GetServ(repName, para)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	result := g.Map{
		"count": count,
	}

	response.SusJson(true, r, "成功", result)

}

//service
/**
获取列表数据
*/
func (ctrl *reportCount) GetServ(repName string, params ReportParams) (count int, err error) {
	sql := makeSql(repName, params.FormData)
	//获取总计数
	var t *gvar.Var
	t, err = boot.ErpDB.GetValue(sql, 0, 0, "count")
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
	}
	count = t.Int()
	return
}
