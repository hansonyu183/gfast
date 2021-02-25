package desc

import (
	"gfast/erp/boot"

	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Desc API管理对象
var Desc = &desc{}

type desc struct {
	name string
}

type DescParam struct {
	Desc string
}

//controller
func (ctrl *desc) Get(r *ghttp.Request) {
	userName := r.GetString("user")
	data, err := ctrl.getData(userName)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//controller
func (ctrl *desc) Put(r *ghttp.Request) {
	var param DescParam
	r.Parse(&param)
	userName := r.GetString("user")
	err := ctrl.saveData(userName, param)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功")
}

//getData
/**
获取明细
*/

func (ctrl *desc) getData(userName string) (data interface{}, err error) {
	mod := boot.ErpDB.Table("desc").Fields("desc").Where("user", userName)
	r, err := mod.Value()

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}

	data = g.Map{
		"data": r,
	}
	return
}

func (ctrl *desc) saveData(userName string, desc DescParam) error {
	mod := boot.ErpDB.Table("desc").Data(g.Map{"user": userName, "desc": desc.Desc})
	_, err := mod.Replace()

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("保存数据失败")
	}

	return err
}
