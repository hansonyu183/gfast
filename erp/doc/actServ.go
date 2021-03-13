package doc

import (
	"fmt"
	"gfast/erp/boot"
)

var actMap = map[string]func(docType string, id int) (err error){
	"check":     doCheck,
	"approve":   doApprove,
	"finish":    doFinish,
	"abort":     doAbort,
	"unCheck":   doUnCheck,
	"unApprove": doUnApprove,
	"unFinish":  doUnFinish,
	"unAbort":   doUnAbort,
}

type Act struct {
	No         string
	SetStateID int
}

func (ctrl *doc) handelAct(docType string, docID, actID int) (err error) {
	var act *Act
	if err = boot.ErpDB.Model("act").Where("id", actID).Struct(&act); err != nil {
		return err
	}
	fmt.Print("%+v\n", act)
	if err = actMap[act.No](docType, docID); err != nil {
		return err
	}

	if _, err = boot.ErpDB.Model(docType).Where("id", docID).Data("state_id", act.SetStateID).Update(); err != nil {
		return err
	}

	return nil
}

func doCheck(docType string, id int) (err error) {
	return nil
}

func doUnCheck(docType string, id int) (err error) {
	return nil
}

func doApprove(docType string, id int) (err error) {
	return nil
}

func doUnApprove(docType string, id int) (err error) {
	return nil
}
func doFinish(docType string, id int) (err error) {
	return nil
}
func doUnFinish(docType string, id int) (err error) {
	return nil
}
func doAbort(docType string, id int) (err error) {
	return nil
}
func doUnAbort(docType string, id int) (err error) {
	return nil
}
