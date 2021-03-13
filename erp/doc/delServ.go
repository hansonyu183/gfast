package doc

import (
	"gfast/erp/boot"
)

func (ctrl *doc) del(docType string, id int) (err error) {
	_, err = boot.ErpDB.Model(docType).Where("id", id).Delete()
	return err
}
