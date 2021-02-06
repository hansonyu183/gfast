package eba

import (
	"gfast/app/model"
	"gfast/erp/api"

	"github.com/gogf/gf/frame/g"
)

// List 客户资料管理
type List struct {
	*api.List
}

// ListParams 数据结构
type ListParams struct {
	EbaId   string `p:"id" orm:"eba.id"`
	EbaName string `p:"name" orm:"eba.name"`
	EmpId   string `p:"emp_id" `
}

type Eba struct {
	*model.Eba
	Emp Emp
}
type Emp struct {
	EmpId   string
	EmpName string
}

func (lp *ListParams) GetCount(likeStr string) (int, error) {
	m := g.DB("erp").Table("eba").LeftJoin("emp", "eba.emp_id=emp.emp_id").OmitEmpty().Where(lp)
	if likeStr != "" {
		likeStr = "%" + likeStr + "%"
		m = m.And("(eba.eba_name like ? or eba.easy_code like ?)", likeStr, likeStr)
	}
	return m.Count()
}

func (lp *ListParams) GetList(pageNum, pageSize int, likeStr string) (list interface{}, err error) {
	var ebas []*Eba
	m := g.DB("erp").Table("eba").LeftJoin("emp", "eba.emp_id=emp.emp_id").OmitEmpty().Where(lp)
	m = m.Fields("eba.*,emp.name emp_name")
	if likeStr != "" {
		likeStr = "%" + likeStr + "%"
		m = m.And("(eba.eba_name like ? or eba.easy_code like ?)", likeStr, likeStr)
	}
	err = m.Page(pageNum, pageSize).Fields("eba.*,emp.emp_id ,emp.name emp_name").Scan(&ebas)
	return ebas, err
}

func NewList() (c *List) {
	c = &List{
		&api.List{
			ListParam: &ListParams{},
		},
	}
	return
}
