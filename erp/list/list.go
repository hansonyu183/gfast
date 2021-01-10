package list

import (
	"gfast/erp/def"

	"gfast/library/response"
	"gfast/library/service"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"

	"strings"
)

type QueryParams struct {
	PageNum  int
	PageSize int
	FormData string
}


type ListCtrl struct {
}

//controller
func (ctrl *ListCtrl) Get(r *ghttp.Request) {
	var para QueryParams
	if err := r.Parse(&para); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	total, page, list, err := GetList(r.GetString("name"), para)
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

func (ctrl *ListCtrl) Delete(r *ghttp.Request) {
	para := r.GetMap()
	if err := Del(r.GetString("name"), para); err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除信息成功")
}

func Del(listName string, para map[string]interface{}) error {
	li := def.ListMap()[listName]
	delTables := strings.Split(li.Del, ",")
	delete(para, "name")
	return g.DB("erp").Transaction(func(tx *gdb.TX) error {
		for _, v := range delTables {
			_, err := g.DB("erp").Table(v).Delete(para)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

//service
/**
获取列表数据
*/
func GetList(listName string, params QueryParams) (total, page int, list interface{}, err error) {
	pageNum := params.PageNum
	pageSize := params.PageSize
	li := def.ListMap()[listName]
	formData := gconv.Map(params.FormData)
	likeStr := gconv.String(formData["likeStr"])
	delete(formData, "likeStr")

	mod := g.DB("erp").Model(li.From).Where(formData).OmitEmpty()
	//处理like查询
	if likeStr != "" && li.Like != "" {
		//替换？号
		likeSql := strings.ReplaceAll(li.Like, "?", "%"+likeStr+"%")
		mod = mod.And(likeSql)
	}
	//获取总计数
	total, err = mod.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
	}
	//分页查询
	if pageNum == 0 {
		pageNum = 1
	}
	page = pageNum

	if pageSize == 0 {
		pageSize = service.AdminPageNum
	}
	if list, err = mod.Fields(li.Select).Page(pageNum, pageSize).All(); err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	return
}

func New() *ListCtrl {
	return &ListCtrl{}
}
