package doc

import (
	"gfast/erp/boot"

	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// DocList API管理对象
var DocList = &docList{}

type docList struct {
	name string
}

type DocListParam struct {
}

//controller
func (ctrl *docList) Get(r *ghttp.Request) {
	docType := r.GetString("type")
	data, err := ctrl.getData(docType)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/

func (ctrl *docList) getData(docType string) (data interface{}, err error) {
	mod := boot.ErpDB.Table(docType)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	r, err := mod.All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}

	data = g.Map{
		"tables": g.Map{
			docType: r,
		},
	}
	return
}
