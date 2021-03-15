package ui

import (
	"gfast/erp/boot"

	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var fields = "id,no,name,py"
var optWhere = "state_id=4"
var optCols = map[string]string{
	"eba":     "id,no,name,py,tel,address,ebasq_id,emp_id,note",
	"sup":     fields,
	"res":     "id,no,name,py,model,pack_num,reskind_id",
	"emp":     fields,
	"account": fields,
	"subject": fields,
	"vtype":   fields,
	"brand":   fields,
	"reskind": fields,
	"state":   "*",
	"act":     "*",
	"unit":    fields,
	"ebasq":   "id,no,name,py,num,price",
	"invres":  "id,no,name,py,model",
	"inveba":  fields,
}

// Opt API管理对象
var Opt = &opt{}

type opt struct {
	name string
}

type OptParam struct {
}

//controller
func (ctrl *opt) Get(r *ghttp.Request) {
	optName := r.GetString("name")
	var data interface{}
	var err error
	if optName == "" {
		data, err = ctrl.getAllOption()
	} else {
		data, err = ctrl.getOneOption(optName)
	}
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/
func (ctrl *opt) getAllOption() (data map[string]interface{}, er error) {
	data = g.Map{}
	for k, v := range optCols {
		r, err := boot.ErpDB.Model(k).Fields(v).Where(optWhere).All()
		if err == nil {
			data[k] = r
		}
	}
	user, _ := g.DB().Model("user").Fields("id,user_name as no,user_nickname as name,'' as py").All()
	data["user"] = user
	return
}

func (ctrl *opt) getOneOption(optName string) (data gdb.Result, err error) {
	return boot.ErpDB.Model(optName).Fields(optCols[optName]).Where(optWhere).All()
}
