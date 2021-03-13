package doc

import (
	"fmt"
	"gfast/erp/boot"
	"strconv"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func (ctrl *doc) insertData(docType string, data *PostData) (dID int, err error) {
	id, err := boot.ErpDB.Table(docType).Value("max(id)")
	if err != nil {
		return 0, err
	}
	docID := id.Int() + 1
	idKey := fmt.Sprintf("%s_id", docType)
	err = boot.ErpDB.Transaction(func(tx *gdb.TX) error {
		// user
		fmData := gconv.Map(data.Forms[docType])
		if len(fmData) > 0 {
			fmData["id"] = strconv.Itoa(docID)
			fmData["state_id"] = "2"
			if _, err = tx.Table(docType).Insert(fmData); err != nil {
				return err
			}
		} else {
			err = gerror.New("表单无数据")
			return err
		}
		for tb, v := range data.Tables {
			tbData := gconv.SliceMap(v)
			for i, v := range tbData {
				if len(v) > 0 {
					tbData[i][idKey] = docID
				} else {
					tbData = append(tbData[:i], tbData[i+1:]...)
				}
			}
			if len(tbData) > 0 {
				if _, err = tx.Table(tb).Insert(tbData); err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		g.Log().Error(err)
		//	err = gerror.New("新增失败")
		return 0, err
	}
	return docID, nil
}
