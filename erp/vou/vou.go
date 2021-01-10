package vou

import (
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type VouCtrl struct {
}

type VouData struct {
	Main map[string]interface{}
	Item map[string]interface{}
	Sub  map[string]interface{}
}

//controller
func (ctrl *VouCtrl) Get(r *ghttp.Request) {
	vouData, err := GetVou(r.GetString("name"), r.GetInt("id"))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	result := g.Map{
		"data": vouData,
	}
	response.SusJson(true, r, "成功", result)

}

func (ctrl *VouCtrl) Delete(r *ghttp.Request) {
	if err := DelVou(r.GetString("name"), r.GetInt("id")); err != nil {
		response.FailJson(true, r, "删除失败")
	}
	response.SusJson(true, r, "删除成功")
}

func (ctrl *VouCtrl) Post(r *ghttp.Request) {
	var vouData VouData
	if err := r.Parse(&vouData); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	if err := CreateVou(r.GetString("name"), &vouData); err != nil {
		response.FailJson(true, r, "新增失败")
	}
	response.SusJson(true, r, "新增成功")
}

func (ctrl *VouCtrl) Put(r *ghttp.Request) {
	var vouData VouData
	if err := r.Parse(&vouData); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	if err := CreateVou(r.GetString("name"), &vouData); err != nil {
		response.FailJson(true, r, "新增失败")
	}
	response.SusJson(true, r, "新增成功")
}

func CreateVou(name string, vouData *VouData) error {
	return g.DB("erp").Transaction(func(tx *gdb.TX) error {
		db := g.DB("erp")
		if _, err := db.Table(name).Data(vouData.Main).Insert(); err != nil {
			return err
		}
		if _, err := db.Table(name + "_item").Data(vouData.Item).Insert(); err != nil {
			return err
		}
		if _, err := db.Table(name + "_sub").Data(vouData.Sub).Insert(); err != nil {
			return err
		}
		return nil
	})
}

func UpdateVou(name string, vouData *VouData) error {
	return g.DB("erp").Transaction(func(tx *gdb.TX) error {
		db := g.DB("erp")
		if _, err := db.Table(name).Data(vouData.Main).Save(); err != nil {
			return err
		}
		if _, err := db.Table(name + "_item").Data(vouData.Item).Save(); err != nil {
			return err
		}
		if _, err := db.Table(name + "_sub").Data(vouData.Sub).Save(); err != nil {
			return err
		}
		return nil
	})
}

func DelVou(name string, id int) error {
	return g.DB("erp").Transaction(func(tx *gdb.TX) error {
		db := g.DB("erp")
		if _, err := db.Table(name).Delete("voucher_id", id); err != nil {
			return err
		}
		if _, err := db.Table(name+"_item").Delete("voucher_id", id); err != nil {
			return err
		}
		if _, err := db.Table(name+"_sub").Delete("voucher_id", id); err != nil {
			return err
		}
		return nil
	})
}

//service
/**
获取列表数据
*/
func GetVou(name string, id int) (vouData map[string]interface{}, err error) {
	var main gdb.Record
	var item, sub gdb.Result
	err = g.DB("erp").Transaction(func(tx *gdb.TX) (err error) {
		db := g.DB("erp")
		if main, err = db.Table(name).One("voucher_id", id); err != nil {
			return err
		}
		if item, err = db.Table(name+"_item").All("voucher_id", id); err != nil {
			return err
		}
		if sub, err = db.Table(name+"_sub").All("voucher_id", id); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	vouData = g.Map{
		"main": main,
		"item": item,
		"sub":  sub,
	}
	return
}

func New() *VouCtrl {
	return &VouCtrl{}
}
