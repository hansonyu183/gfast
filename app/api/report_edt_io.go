package report

import (
	"gfast/app/model"
	"gfast/app/service"
	"gfast/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// ReportEdtIO API管理对象
var ReportEdtIO = new(reportEdtIOApi)

type reportEdtIOApi struct{}

// @summary 库存IO报表查询接口
// @tags    库存服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "库存IO报表请求"
// @router  /report/edtIO [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *reportEdtIOApi) Get(r *ghttp.Request) {
	var (
		apiReq *model.ReportEdtIoApiReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data, err := service.ReportEdtIO.Get(apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", data)
	}
}
