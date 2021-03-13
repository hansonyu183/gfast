package vr

import (
	"gfast/erp/boot"

	"github.com/gogf/gf/container/gvar"
)

func (ctrl *vr) del(vrType string, id int) (err error) {
	var vtypeID *gvar.Var
	if vtypeID, err = boot.ErpDB.Model("vtype").Fields("id").Where("no", vrType).Value(); err != nil {
		return err
	}

	_, err = boot.ErpDB.Model("vr").Where("id,vtype_id", id, vtypeID).Delete()
	return err
}
