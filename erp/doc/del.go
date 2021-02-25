package doc

import (
	"gfast/erp/boot"
)

/*
//controller
func (ctrl *doc) Delete(r *ghttp.Request) {
	docType := r.GetString("type")
	docID := r.GetInt("id")
	err := ctrl.delData(docType, docID)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功")
}*/

func (ctrl *doc) delData(docType string, docID int) (err error) {
	_, err = boot.ErpDB.DelTbFkSub(docType, docID)
	return
}
