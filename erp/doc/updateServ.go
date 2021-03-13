package doc

import (
	"fmt"
	"gfast/erp/boot"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func (ctrl *doc) updateData(docType string, docID int, data *PostData) (dID int, err error) {
	err = boot.ErpDB.Transaction(func(tx *gdb.TX) error {
		fmData := gconv.Map(data.Forms[docType])
		if len(fmData) > 0 {
			if _, err = tx.Table(docType).Where("id", docID).Data(fmData).Update(); err != nil {
				return err
			}
		} else {
			err = gerror.New("表单无有效数据")
			return err
		}

		idKey := fmt.Sprintf("%s_id", docType)
		for tb, v := range data.Tables {
			if _, err = tx.Table(tb).Delete(idKey, docID); err != nil {
				return err
			}
			tbData := gconv.SliceMap(v)

			for i, v := range tbData {
				if len(v) > 0 {
					tbData[i][idKey] = docID
				} else {
					tbData = append(tbData[:i], tbData[i+1:]...)
				}
			}

			if len(tbData) == 0 {
				continue
			}

			if _, err = tx.Table(tb).Insert(tbData); err != nil {
				return err
			}

		}
		return nil
	})
	if err != nil {
		g.Log().Error(err)
		//	err = gerror.New("保存失败")
		return 0, err
	}
	return docID, nil
}

/*
func TrimMapNull(mapData map[string]interface{}) map[string]interface{} {
	for k, v := range mapData {
		if v == nil {
			delete(mapData, k)
		}
	}
	return mapData
}
*/
