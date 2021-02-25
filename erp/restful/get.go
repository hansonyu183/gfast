package restful

import (
	"gfast/library/response"

	"gfast/library/service"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type RfGet struct {
	GetMod    *gdb.Model
	GetParams interface{}
}

type ListPage struct {
	PageNum  int
	PageSize int
}

//controller
func (rf *RfGet) Get(r *ghttp.Request) {

	if err := r.Parse(rf.GetParams); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	pg := new(ListPage)
	pg.PageNum = r.GetInt("pageNum")
	pg.PageSize = r.GetInt("pageSize")
	total, page, list, err := GetList(rf.GetMod, pg, rf.GetParams)

	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	result := g.Map{
		"total": total,
		"list":  list,
		"page":  page,
	}

	response.SusJson(true, r, "成功", result)

}

//service
/**
获取列表数据
*/
func GetList(mod *gdb.Model, pg *ListPage, params interface{}) (total, page int, list gdb.Result, err error) {
	m := mod.Clone().Where(params)

	total, err = m.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
	}

	if pg.PageNum == 0 {
		pg.PageNum = 1
	}

	page = pg.PageNum

	if pg.PageSize == 0 {
		pg.PageSize = service.AdminPageNum
	}
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
	}
	list, err = m.Page(page, pg.PageSize).All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}

	return
}
