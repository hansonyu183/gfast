package doc

import (
	"gfast/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// Doc API管理对象
var Doc = &doc{}

type doc struct {
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

type PostData struct {
	Forms  map[string]string
	Tables map[string]string
}

//controller
func (ctrl *doc) Post(r *ghttp.Request) {
	var docData *PostData
	var err error
	if err = r.Parse(&docData); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	docType := r.GetString("type")
	docID := r.GetInt("id")
	var dID int
	if docID == 0 {
		dID, err = ctrl.insertData(docType, docData)
	} else {
		dID, err = ctrl.updateData(docType, docID, docData)
	}

	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", dID)
}

//controller
func (ctrl *doc) Patch(r *ghttp.Request) {
	var err error

	docType := r.GetString("type")
	docID := r.GetInt("id")
	actID := r.GetInt("actID")
	if err = ctrl.handelAct(docType, docID, actID); err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功")
}

//controller
func (ctrl *doc) Delete(r *ghttp.Request) {
	docType := r.GetString("type")
	docID := r.GetInt("id")
	err := ctrl.del(docType, docID)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功")
}
