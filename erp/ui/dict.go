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
	ID       int        `orm:"id,primary"   json:"id"`
	DictName string     `orm:"dict_name"   json:"dictName"`
	DictType string     `orm:"dict_type"   json:"dictType"`
	Data     []DictData ` json:"data"`
}

type DictData struct {
	DictCode  int    `orm:"dict_code"   json:"dictCode"`
	DictLabel string `orm:"dict_label"   json:"dictLabel"`
	DictValue string `orm:"dict_value"   json:"dictValue"`
	IsDefault bool   `orm:"is_default"   json:"isDefault"`
	ReMark    string `orm:"remark"   json:"remark"`
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
		err = g.DB().Model("sys_dict_data").Where("dict_type", v.DictType).Structs(&dictData)
		if err != nil {
			continue
		}
		dicts[k].Data = dictData
	}

	return dicts, nil
}
