package vr

import (
	"fmt"
	"gfast/erp/boot"
	"strings"
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gdb"
)

type ResponseGet struct {
	NextID *gvar.Var  `json:"nextID"`
	PreID  *gvar.Var  `json:"preID"`
	Vr     gdb.Record `json:"vr"`
	VrMain gdb.Record `json:"vrMain"`
	VrAmo  gdb.Result `json:"vrAmo"`
	VrNum  gdb.Result `json:"vrNum"`
	VrPf   gdb.Result `json:"vrPf"`
}

//getData
/**
获取明细
*/
func (ctrl *vr) getData(vtype string, vrID int) (data *ResponseGet, err error) {
	if vrID == 0 {
		return getNewVr(vtype)
	}
	var vtypeID *gvar.Var
	if vtypeID, err = boot.ErpDB.Model("vtype").Fields("id").Where("no", vtype).Value(); err != nil {
		return nil, err
	}
	data = &ResponseGet{}
	data.Vr, data.VrMain, data.VrAmo, data.VrNum, data.VrPf, err = getVrByID(vrID)
	if err != nil {
		return nil, err
	}

	if data.NextID, err = boot.ErpDB.Model("vr").Where("vtype_id=? and id>?", vtypeID, vrID).Value("min(id)"); err != nil {
		return nil, err
	}
	if data.PreID, err = boot.ErpDB.Model("vr").Where("vtype_id=? and id<?", vtypeID, vrID).Value("max(id)"); err != nil {
		return nil, err
	}
	return data, nil
}

func getVrByID(id int) (vr, vrMain gdb.Record, vrAmo, vrNum, vrPf gdb.Result, err error) {
	db := boot.ErpDB
	if vr, err = db.Model("vr").Where("id", id).One(); err != nil {
		return
	}
	if vrMain, err = db.Model("vr_main").Where("vid", id).One(); err != nil {
		return
	}
	if vrAmo, err = db.Model("vr_amo").Where("vid", id).All(); err != nil {
		return
	}
	if vrNum, err = db.Model("vr_num").Where("vid", id).All(); err != nil {
		return
	}
	if vrPf, err = db.Model("vr_pf").Where("vid", id).All(); err != nil {
		return
	}
	return
}

func getNewVr(vtype string) (data *ResponseGet, err error) {
	var vtypeID *gvar.Var
	if vtypeID, err = boot.ErpDB.Model("vtype").Fields("id").Where("no", vtype).Value(); err != nil {
		return nil, err
	}
	var vrMain gdb.Record
	var vrNum gdb.Record
	var vrAmo gdb.Record
	var vrPf gdb.Record

	if vrMain, err = boot.ErpDB.GetTbDefaultData("vr_main"); err != nil {
		return nil, err
	}

	if vrNum, err = boot.ErpDB.GetTbDefaultData("vr_num"); err != nil {
		return nil, err
	}
	vrNum["id"] = gvar.New(1, true)

	if vrAmo, err = boot.ErpDB.GetTbDefaultData("vr_amo"); err != nil {
		return nil, err
	}
	vrAmo["id"] = gvar.New(1, true)

	if vrPf, err = boot.ErpDB.GetTbDefaultData("vr_pf"); err != nil {
		return nil, err
	}
	vrPf["id"] = gvar.New(1, true)
	vrPf["iid"] = gvar.New(1, true)
	//vrPf["id"] = gvar.New(1, true)
	//vrPf["iid"] = gvar.New(1, true)*/

	vr := make(gdb.Record)
	data = &ResponseGet{
		Vr:     vr,
		VrMain: vrMain,
		VrNum:  []gdb.Record{vrNum},
		VrAmo:  []gdb.Record{vrAmo},
		VrPf:   []gdb.Record{vrPf},
	}
	var no, toDay string
	if no, toDay, err = MakeVrNewNo(vtype); err != nil {
		return nil, err
	}

	if data.PreID, err = boot.ErpDB.Model("vr").Where("vtype_id", vtypeID).Value("max(id)"); err != nil {
		return nil, err
	}
	data.Vr["id"] = gvar.New(0, true)
	data.Vr["no"] = gvar.New(no, true)
	data.Vr["date"] = gvar.New(toDay, true)

	return data, nil
}

func MakeVrNewNo(vtype string) (newNo, toDay string, err error) {
	toDay = time.Now().Format("20060102")
	var toDayVrNum int
	if toDayVrNum, err = boot.ErpDB.Model("vr").Where("date", toDay).Count(); err != nil {
		return
	}
	newNo = fmt.Sprintf("%s%s-%d", strings.ToUpper(vtype), toDay, toDayVrNum+1)
	return
}
