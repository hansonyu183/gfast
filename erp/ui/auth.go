package ui

import (
	"gfast/erp/boot"
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Auth API管理对象
var Auth = &auth{}

type auth struct {
	name string
}

//controller
func (ctrl *auth) Get(r *ghttp.Request) {
	userName := r.GetString("user")
	data, err := ctrl.getData(userName)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/
type ResponseGet struct {
	Vr  gdb.Result `json:"vr"`
	Doc gdb.Result `json:"doc"`
}

func (ctrl *auth) getData(userName string) (data *ResponseGet, err error) {
	sql := "SELECT substring(v1, 3) FROM gfast.casbin_rule where v0=(select concat('u_',id) from gfast.user where user_name=?)"
	roleID, err := g.DB().GetValue(sql, userName)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return nil, err
	}

	vr, err := boot.ErpDB.Model("auth_vr").Fields("id,type,act_ids").Where("role_id", roleID).All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return nil, err
	}

	doc, err := boot.ErpDB.Model("auth_doc").Fields("id,type,act_ids").Where("role_id", roleID).All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return nil, err
	}

	data = &ResponseGet{
		Vr:  vr,
		Doc: doc,
	}
	return data, nil
}
