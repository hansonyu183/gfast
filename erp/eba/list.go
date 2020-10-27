package eba

import (
	"gfast/erp/api"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//客户资料管理
type List struct {
	*api.List
}

//model
//数据结构
type ListParams struct {
	Id    string `p:"id" orm:"eba.id"`
	Name  string `p:"name" orm:"eba.name"`
	EmpId string `p:"emp_id" `
}

func (lp *ListParams) GetCount(likeStr string) (int, error) {
	m := g.DB("erp").Table("eba").LeftJoin("emp", "eba.emp_id=emp.id").OmitEmpty().Where(lp)
	if likeStr != "" {
		likeStr = "%" + likeStr + "%"
		m = m.And("(eba.name like ? or eba.py like ?)", likeStr, likeStr)
	}
	return m.Count()
}

func (lp *ListParams) GetList(pageNum, pageSize int, likeStr string) (list gdb.Result, err error) {
	m := g.DB("erp").Table("eba").LeftJoin("emp", "eba.emp_id=emp.id").OmitEmpty().Where(lp)
	m=m.Fields("eba.*,emp.name emp_name")
	if likeStr != "" {
		likeStr = "%" + likeStr + "%"
		m = m.And("(eba.name like ? or eba.py like ?)", likeStr, likeStr)
	}
	return m.Page(pageNum, pageSize).All()
}

func NewList() (c *List) {
	c = &List{
		&api.List{
			ListParam: &ListParams{},
		},
	}
	return
}

