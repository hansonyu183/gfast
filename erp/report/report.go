package report

import (
	"fmt"

	"gfast/erp/boot"

	"gfast/library/response"
	"gfast/library/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

// Report API管理对象
var Report = new(report)

type report struct {
}

type ReportParams struct {
	PageNum  int
	PageSize int
	FormData string
}

//controller
func (ctrl *report) Get(r *ghttp.Request) {
	var para ReportParams
	repName := r.GetString("name")
	if err := r.Parse(&para); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	page, report, err := ctrl.GetServ(repName, para)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	result := g.Map{
		"report": report,
		"page":   page,
	}

	response.SusJson(true, r, "成功", result)

}

//service
/**
获取列表数据
*/
func (ctrl *report) GetServ(repName string, params ReportParams) (page int, report interface{}, err error) {
	pageNum := params.PageNum
	pageSize := params.PageSize
	sql := makeSql(repName, params.FormData)

	//分页查询
	if pageNum == 0 {
		pageNum = 1
	}
	page = pageNum

	if pageSize == 0 {
		pageSize = service.AdminPageNum
	}

	offset := (pageNum - 1) * pageSize
	if report, err = boot.ErpDB.GetAll(sql, offset, pageSize, "item"); err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	return
}

func makeSql(repName, queryData string) (sql string) {
	formData := gconv.Map(queryData)
	pa := "("
	for _, v := range boot.ReportMap[repName].Params {
		if formData[v] == nil {
			pa = fmt.Sprintf("%s'%s',", pa, "")
		} else {
			pa = fmt.Sprintf("%s'%s',", pa, formData[v])
		}
	}
	pa = pa + "?,?,?)"
	sql = boot.ReportMap[repName].SQL + pa
	return
}
