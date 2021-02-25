package doc

import (
	"gfast/erp/boot"
	"gfast/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// Doc API管理对象
var DocAct = &docAct{}

type docAct struct {
}

//controller
func (ctrl *docAct) Post(r *ghttp.Request) {
	var docData *DocData
	var err error
	if err = r.Parse(&docData); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	docType := r.GetString("type")
	docID := r.GetInt("id")
	act := r.GetString("act")
	var dID int
	dID, err = ctrl.handelAct(docType, act, docID, docData)

	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", dID)
}

var actMap = map[string]func(docType string, id int, data *DocData) (dID int, err error){
	"save":      doSave,
	"check":     doCheck,
	"approve":   doApprove,
	"finish":    doFinish,
	"abort":     doAbort,
	"del":       doDel,
	"unCheck":   doUnCheck,
	"unApprove": doUnApprove,
	"unFinish":  doUnFinish,
	"unAbort":   doUnAbort,
}

func (ctrl *docAct) handelAct(docType, act string, id int, data *DocData) (dID int, err error) {
	return actMap[act](docType, id, data)
}

func updateState(docType, act string, id int) (dID int, err error) {
	setStateID, err := boot.ErpDB.Model("act").Where("no", act).Value("set_state_id")
	if err != nil {
		return 0, err
	}
	_, err = boot.ErpDB.Model(docType).Where("id", id).Data("state_id", setStateID).Update()
	return id, err
}

func doSave(docType string, id int, data *DocData) (dID int, err error) {
	if id == 0 {
		dID, err = insertData(docType, data)
	} else {
		dID, err = saveData(docType, id, data)
	}
	return dID, err
}

func doDel(docType string, id int, data *DocData) (dID int, err error) {
	_, err = boot.ErpDB.Model(docType).Where("id", id).Delete()
	return 0, err
}

func doCheck(docType string, id int, data *DocData) (dID int, err error) {
	return updateState(docType, "check", id)
}

func doUnCheck(docType string, id int, data *DocData) (dID int, err error) {
	return updateState(docType, "unCheck", id)
}

func doApprove(docType string, id int, data *DocData) (dID int, err error) {
	return updateState(docType, "approve", id)
}

func doUnApprove(docType string, id int, data *DocData) (dID int, err error) {
	return updateState(docType, "unApprove", id)
}
func doFinish(docType string, id int, data *DocData) (dID int, err error) {
	return updateState(docType, "finish", id)
}
func doUnFinish(docType string, id int, data *DocData) (dID int, err error) {
	return updateState(docType, "unFinish", id)
}
func doAbort(docType string, id int, data *DocData) (dID int, err error) {
	return updateState(docType, "abort", id)
}
func doUnAbort(docType string, id int, data *DocData) (dID int, err error) {
	return updateState(docType, "unAbort", id)
}
