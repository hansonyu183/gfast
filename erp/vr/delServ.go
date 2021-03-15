package vr

import (
	"gfast/erp/boot"
)

func (ctrl *vr) del(id int) (err error) {

	_, err = boot.ErpDB.Model("vr").Where("id", id).Delete()
	return err
}
