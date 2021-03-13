package vr

import (
	"gfast/erp/boot"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

func (ctrl *vr) insertData(vrType string, data *PostData) (dID uint, err error) {
	id, err := boot.ErpDB.Model("vr").Value("max(id)")
	if err != nil {
		return 0, err
	}
	var vtypeID *gvar.Var
	if vtypeID, err = boot.ErpDB.Model("vtype").Fields("id").Where("no", vrType).Value(); err != nil {
		return 0, err
	}

	vid := id.Uint() + 1
	err = boot.ErpDB.Transaction(func(tx *gdb.TX) error {
		if data.Vr == nil {
			return gerror.New("表单无数据")
		}
		data.Vr.ID = vid
		data.Vr.StateID = 2
		data.Vr.VtypeID = vtypeID.Uint()
		if _, err = tx.Model("vr").Insert(data.Vr); err != nil {
			return err
		}
		if data.VrMain == nil {
			data.VrMain.VID = vid
			if _, err = tx.Model("vr_main").Insert(data.VrMain); err != nil {
				return err
			}
		}
		if len(data.VrAmo) > 0 {
			for _, v := range data.VrAmo {
				v.VID = vid
			}
			if _, err = tx.Model("vr_amo").Insert(data.VrAmo); err != nil {
				return err
			}
		}
		if len(data.VrNum) > 0 {
			for _, v := range data.VrNum {
				v.VID = vid
			}
			if _, err = tx.Model("vr_num").Insert(data.VrNum); err != nil {
				return err
			}
		}
		if len(data.VrPf) > 0 {
			for _, v := range data.VrPf {
				v.VID = vid
			}
			if _, err = tx.Model("vr_pf").Insert(data.VrPf); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		g.Log().Error(err)
		//	err = gerror.New("新增失败")
		return 0, err
	}
	return vid, nil
}

func (ctrl *vr) updateData(vrType string, vrID uint, data *PostData) (dID uint, err error) {
	err = boot.ErpDB.Transaction(func(tx *gdb.TX) error {
		if data.Vr == nil {
			return gerror.New("表单无数据")
		}
		if _, err = tx.Model("vr").Where("id", vrID).Update(data.Vr); err != nil {
			return err
		}
		if data.VrMain == nil {
			if _, err = tx.Model("vr_main").Where("vid", vrID).Update(data.VrMain); err != nil {
				return err
			}
		}
		if len(data.VrAmo) > 0 {
			if _, err = tx.Model("vr_amo").Where("vid", vrID).Save(data.VrAmo); err != nil {
				return err
			}
		}
		if len(data.VrNum) > 0 {
			if _, err = tx.Model("vr_num").Where("vid", vrID).Save(data.VrNum); err != nil {
				return err
			}
		}
		if len(data.VrPf) > 0 {
			if _, err = tx.Model("vr_pf").Where("vid", vrID).Delete(); err != nil {
				return err
			}
			if _, err = tx.Model("vr_pf").Insert(data.VrPf); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		g.Log().Error(err)
		//	err = gerror.New("新增失败")
		return 0, err
	}
	return vrID, nil
}
