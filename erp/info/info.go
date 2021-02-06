package info

import (
	"gfast/erp/boot"
	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Report API管理对象
var Info = new(info)

type info struct {
}

//controller
func (ctrl *info) Get(r *ghttp.Request) {
	infoName := r.GetString("name")
	info, err := ctrl.GetServ(infoName)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	result := g.Map{
		"info": info,
	}

	response.SusJson(true, r, "成功", result)

}

//service
/**
获取列表数据
*/
func (ctrl *info) GetServ(infoName string) (info interface{}, er error) {
	db := boot.ErpDB
	switch infoName { //finger is declared in switch
	case "ERP":
		appDict, err := db.Table("app_dict").Fields("id,name").All()
		eba, err := db.Table("eba").Fields("id,name").All()
		sup, err := db.Table("sup").Fields("id,name").All()
		res, err := db.Table("res").Fields("id,name").All()
		emp, err := db.Table("emp").Fields("id,name").All()
		account, err := db.Table("account").Fields("id,name").All()
		subject, err := db.Table("subject").Fields("id,name").All()
		vrType, err := db.Table("vr_type").Fields("id,name").All()
		info = g.Map{
			"app_dict": appDict,
			"eba":      eba,
			"sup":      sup,
			"res":      res,
			"emp":      emp,
			"account":  account,
			"subject":  subject,
			"vr_type":  vrType,
		}
		er = err
	default: //default case
		info, er = db.Table(infoName).Fields("id,name").All()
	}
	if er != nil {
		g.Log().Error(er)
		return nil, gerror.New("获取数据失败")
	}
	return
}
