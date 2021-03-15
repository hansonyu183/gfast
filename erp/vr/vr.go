package vr

import (
	"gfast/erp/vr/model"
	"gfast/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// Vr API管理对象
var Vr = &vr{}

type vr struct {
}

//controller
func (ctrl *vr) Get(r *ghttp.Request) {
	vrType := r.GetString("type")
	vID := r.GetInt("id")
	data, err := ctrl.getData(vrType, vID)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

type PostData struct {
	Vr     *model.Vr
	VrMain *model.VrMain
	VrAmo  []*model.VrAmo
	VrNum  []*model.VrNum
	VrPf   []*model.VrPf
}

//controller
func (ctrl *vr) Post(r *ghttp.Request) {
	var vrData *PostData
	var err error
	if err = r.Parse(&vrData); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	vrType := r.GetString("type")
	vID := r.GetUint("id")
	var effectID uint
	if vID == 0 {
		effectID, err = ctrl.insertData(vrType, vrData)
	} else {
		effectID, err = ctrl.updateData(vrType, vID, vrData)
	}

	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", effectID)
}

//controller
func (ctrl *vr) Patch(r *ghttp.Request) {
	var err error

	vrType := r.GetString("type")
	vID := r.GetInt("id")
	actID := r.GetInt("actID")
	if err = ctrl.handelAct(vrType, vID, actID); err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功")
}

//controller
func (ctrl *vr) Delete(r *ghttp.Request) {
	vID := r.GetInt("id")
	err := ctrl.del(vID)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功")
}
