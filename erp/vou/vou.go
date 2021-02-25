package vou

import (
	"gfast/erp/boot"

	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Vou API管理对象
var Vou = &vouData{}

type vouData struct {
}

type VouParam struct {
}

//controller
func (ctrl *vouData) Get(r *ghttp.Request) {
	vouDataID := r.GetInt("id")
	data, err := ctrl.getData(vouDataID)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/

func (ctrl *vouData) getData(vouDataID int) (data interface{}, err error) {
	vr := boot.ErpDB.Table("vr")
	vMain := boot.ErpDB.Table("vr_main")
	vItemAmo := boot.ErpDB.Table("vr_item_amo")
	vItemNum := boot.ErpDB.Table("vr_item_num")

	vrData, err := vr.Where("vid", vouDataID).One()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	vMainData, err := vMain.Where("vid", vouDataID).One()
	vItemAmoData, err := vItemAmo.Where("vid", vouDataID).All()
	vItemNumData, err := vItemNum.Where("vid", vouDataID).All()
	fmData := g.Map{}
	if vrData != nil {
		fmData["vr"] = vrData
	}
	if vMainData != nil {
		fmData["main"] = vMainData
	}
	tbData := g.Map{}
	if vItemAmoData != nil {
		tbData["amoItem"] = vItemAmoData
	}
	if vItemNumData != nil {
		tbData["numItem"] = vItemNumData
	}
	data = g.Map{
		"forms":  fmData,
		"tables": tbData,
	}
	return
}
