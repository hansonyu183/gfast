package api

import (
	"database/sql"
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type EditParam interface {
	Create() (sql.Result, error)
	Update() (sql.Result, error)
}

type DelParam interface {
	Del() (sql.Result, error)
}

type GetVouParam interface {
	GetVou() (gdb.Record, error)
}

type Vou struct {
	GetVouParam GetVouParam
	EditParam   EditParam
	DelParam    DelParam
}

//controller
func (v *Vou) Get(r *ghttp.Request) {
	if err := r.Parse(v.GetVouParam); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	vou, err := v.GetVouParam.GetVou()
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	response.SusJson(true, r, "成功", vou)
}

func (v *Vou) Post(r *ghttp.Request) {
	if err := r.Parse(v.EditParam); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	if _, err := v.EditParam.Create(); err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "添加成功")
}

func (v *Vou) Put(r *ghttp.Request) {

	if err := r.Parse(v.EditParam); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	
	if _, err := v.EditParam.Update(); err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "修改成功")
}

func (v *Vou) Delete(r *ghttp.Request) {
	if err := r.Parse(v.DelParam); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	if _, err := v.DelParam.Del(); err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除信息成功")
}
