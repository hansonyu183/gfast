package doc

import (
	"gfast/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// Doc API管理对象
var Doc = &doc{}

type doc struct {
}

type DocParam struct {
}

//controller
func (ctrl *doc) Get(r *ghttp.Request) {
	docType := r.GetString("type")
	docID := r.GetInt("id")
	data, err := ctrl.getData(docType, docID)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}
