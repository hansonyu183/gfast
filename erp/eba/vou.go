package eba

import (
	"database/sql"
	"gfast/erp/api"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//客户资料管理
type Vou struct {
	*api.Vou
}

//model
//数据结构
type GetVouParam struct {
	Id string `p:"id" `
}

func (gp *GetVouParam) GetVou() (gdb.Record, error) {
	mod := g.DB("erp").Table("eba").OmitEmpty()
	return mod.FindOne(gp)
}

type EditParam struct {
	Id    string `p:"id" `
	Name  string `p:"name" `
	EmpId string `p:"emp_id" `
	Note  string `p:"note" `
}

func (ep *EditParam) Create() (sql.Result, error) {
	m := g.DB("erp").Table("eba").OmitEmpty()
	if rep, err := m.Data(ep).Insert(); err != nil {
		return nil, err
	} else {
		return rep, nil
	}
}

func (ep *EditParam) Update() (sql.Result, error) {
	m := g.DB("erp").Table("eba").OmitEmpty().Where("id", ep.Id)
	if rep, err := m.Data(ep).Update(); err != nil {
		return nil, err
	} else {
		return rep, nil
	}
}

type DelParam struct {
	Ids []string `p:"ids" `
}

func (dp *DelParam) Del() (sql.Result, error) {
	m := g.DB("erp").Table("eba").Where("id", dp.Ids)
	if rep, err := m.Delete(); err != nil {
		return nil, err
	} else {
		return rep, nil
	}
}

func NewVou() (c *Vou) {
	c = &Vou{
		&api.Vou{
			GetVouParam: &GetVouParam{},
			EditParam:   &EditParam{},
			DelParam:    &DelParam{},
		},
	}
	return
}
