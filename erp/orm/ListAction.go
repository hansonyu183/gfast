package orm

import (
	"gfast/library/response"
	"gfast/library/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type ListAction struct {
	ListMod    interface{}
	QueryParam interface{}
}

// @Summary 读取多条
// @Description 读取多条
// @Tags 读取多条
// @Param is page parames and SearchParams date like "{"pageNum": 1,"pageSize": 15, id"：1，“name": john}"
// @Success 0 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /getList
// @Security
func (act *ListAction) Get(r *ghttp.Request) {

	if err := r.Parse(act.QueryParam); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	pn := r.GetInt("pageNum")
	ps := r.GetInt("pageSize")
	total, page, err := Get(pn, ps, act.ListMod, act.QueryParam)

	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	result := g.Map{
		"total": total,
		"list":  act.ListMod,
		"page":  page,
	}

	response.SusJson(true, r, "成功", result)

}

//service
func Get(pageNum, pageSize int, mod, params interface{}) (total int64, page int, err error) {

	if dbErr := DB.Model(mod).Where(params).Count(&total).Error; dbErr != nil {
		g.Log().Error(dbErr)
		err = gerror.New("获取总行数失败")
	}

	if pageNum == 0 {
		pageNum = 1
	}
	page = pageNum

	if pageSize == 0 {
		pageSize = service.AdminPageNum
	}
	if dbErr := DB.Preload("AppEmp").Where(params).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(mod).Error; dbErr != nil {
		g.Log().Error(dbErr)
		err = gerror.New("获取数据失败")
		return
	}

	return
}

//service
func GetItem(mod, params interface{}) (total int64, page int, err error) {

	if dbErr := DB.Where(params).Find(mod).Error; dbErr != nil {
		g.Log().Error(dbErr)
		err = gerror.New("获取数据失败")
		return
	}

	return
}
