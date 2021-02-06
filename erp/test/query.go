package test

import (
	"gfast/erp/boot"
	"gfast/erp/e"

	"gfast/library/response"
	"gfast/library/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// Query API管理对象
var Query = &query{
	serv: map[string]queryServ{
		"sum":  getSum,
		"data": getData,
	},
}

type queryServ func(qName string, params RequestParams) (data interface{}, err error)

type query struct {
	serv map[string]queryServ
}

type RequestParams struct {
	PageNum   int
	PageSize  int
	WherePara map[string]string
	Groups    string
}

//controller
func (ctrl *query) Get(r *ghttp.Request) {
	var para RequestParams
	qName := r.GetString("name")
	qSub := r.GetString("sub")
	if err := r.Parse(&para); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}

	if ctrl.serv[qSub] == nil {
		err := gerror.New("资源子类不存在")
		g.Log().Error(err)
		response.FailJson(true, r, err.Error())
	}
	data, err := ctrl.serv[qSub](qName, para)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	response.SusJson(true, r, "成功", data)

}

//server
func makeMod(qName string, reqData RequestParams) *e.Mod {
	qm := boot.QueryMap[qName]
	mod := e.DB.From(qm.Model).Select(qm.Fields).Where(qm.WhereMap).WhereParas(reqData.WherePara).Group(reqData.Groups).Order(qm.Orders)
	return mod
}

//getData
/**
获取明细
*/

func getData(qName string, params RequestParams) (data interface{}, err error) {
	pageNum := params.PageNum
	pageSize := params.PageSize
	mod := makeMod(qName, params)

	//分页查询
	if pageNum == 0 {
		pageNum = 1
	}
	page := pageNum

	if pageSize == 0 {
		pageSize = service.AdminPageNum
	}
	r, err := mod.Page(pageNum, pageSize).All()
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

//getSum
/**
获取汇总
*/
func getSum(qName string, params RequestParams) (data interface{}, err error) {
	mod := makeMod(qName+"Sum", params)
	if data, err = mod.One(); err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	return
}
