package ui

import (
	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Dict API管理对象
var Dict = &dict{}

type dict struct {
}

//controller
func (ctrl *dict) Get(r *ghttp.Request) {
	data, err := ctrl.getData()
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/
type SysDict struct {
	ID       int
	DictName string
	DictType string
	Data     []DictData
}

type DictData struct {
	DictCode  int
	DictLabel string
	DictValue string
	IsDefault bool
	ReMark    string
}

func (ctrl *dict) getData() (data []SysDict, err error) {
	dicts := ([]SysDict)(nil)
	err = g.DB().Model("sys_dict_type").Structs(&dicts)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return nil, err
	}
	for k, v := range dicts {
		dictData := ([]DictData)(nil)
		err = g.DB().Model("sys_dict_data").Where("dict_type", v.DictType).Structs(&dictDatas)
		if err != nil {
			continue
		}
		dicts[k].Data = dictData
	}

	return dicts, nil
}
