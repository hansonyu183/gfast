package vou

import (
	"gfast/app/restful"

	"github.com/gogf/gf/frame/g"
)

//客户资料管理
type EbaInfo struct {
	*restful.Get
	*restful.Post
}

func NewEbaInfo() (e *EbaInfo) {
	e = &EbaInfo{
		&restful.Get{
			Params: &GetParams{},
			Mod:    getMod,
		},
		&restful.Post{
			Params: &PostParams{},
			Mod:    postMod,
		},
	}
	return e
}

//model
//需求参数
type GetParams struct {
	EbaName string
	EmpId   string
}

type PostParams struct {
	EbaName string
	EmpId   string
}

var getMod = g.DB().Table("eba").Fields("eba_id,eba_name,emp_id")
var postMod = g.DB().Table("eba")
