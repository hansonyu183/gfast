package vr

import (
	"gfast/erp/boot"
)

var actMap = map[string]func(vrType string, id int) (err error){
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

func (ctrl *vr) handelAct(vrType string, vrID, actID int) (err error) {
	var act *Act
	if err = boot.ErpDB.Model("act").Where("id", actID).Struct(&act); err != nil {
		return err
	}
	if err = actMap[act.No](vrType, vrID); err != nil {
		return err
	}

	if _, err = boot.ErpDB.Model("vr").Where("id", vrID).Data("state_id", act.SetStateID).Update(); err != nil {
		return err
	}

	return nil
}

func doCheck(vrType string, id int) (err error) {

	return nil
}

func autoApprove(vrType string, id int) (msg string,err error) {
	
	return "",nil
}

func checkPrice(){

}

func doUnCheck(vrType string, id int) (err error) {
	return nil
}

func doApprove(vrType string, id int) (err error) {
	return nil
}

func doUnApprove(vrType string, id int) (err error) {
	return nil
}
func doFinish(vrType string, id int) (err error) {
	return nil
}
func doUnFinish(vrType string, id int) (err error) {
	return nil
}
func doAbort(vrType string, id int) (err error) {
	return nil
}
func doUnAbort(vrType string, id int) (err error) {
	return nil
}
