package api

import (
	"gfast/library/response"

	"gfast/library/service"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type List struct {
	ListParam ListParam
}

type ListParam interface {
	GetList(pageNum, pageSize int, likeStr string) (list gdb.Result, err error)
	GetCount(likeStr string) (int, error)
}

//controller
func (l *List) Get(r *ghttp.Request) {

	if err := r.Parse(l.ListParam); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	pageNum := r.GetInt("pageNum")
	pageSize := r.GetInt("pageSize")
	likeStr := r.GetString("likeStr")
	total, page, list, err := GetList(pageNum, pageSize, l.ListParam, likeStr)

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
func GetList(pageNum, pageSize int, lp ListParam, likeStr string) (total, page int, list gdb.Result, err error) {
	total, err = lp.GetCount(likeStr)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
	}

	if pageNum == 0 {
		pageNum = 1
	}

	page = pageNum

	if pageSize == 0 {
		pageSize = service.AdminPageNum
	}
	list, err = lp.GetList(pageNum, pageSize, likeStr)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}

	return
}
