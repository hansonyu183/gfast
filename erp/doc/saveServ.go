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

type DocData struct {
	Form   map[string]string
	Tables map[string]string
}

func insertData(docType string, data *DocData) (dID int, err error) {
	id, err := boot.ErpDB.Table(docType).Value("max(id)")
	if err != nil {
		return 0, err
	}
	docID := id.Int() + 1
	idKey := fmt.Sprintf("%s_id", docType)
	err = boot.ErpDB.Transaction(func(tx *gdb.TX) error {
		// user
		fmData := data.Form
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
				tbData[i] = TrimMapNull(v)
				if len(tbData[i]) > 0 {
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

func saveData(docType string, docID int, data *DocData) (dID int, err error) {
	idKey := fmt.Sprintf("%s_id", docType)
	err = boot.ErpDB.Transaction(func(tx *gdb.TX) error {
		fmData := data.Form
		if len(fmData) > 0 {
			if _, err = tx.Table(docType).Where("id", docID).Data(fmData).Update(); err != nil {
				return err
			}
		} else {
			err = gerror.New("表单无有效数据")
			return err
		}
		fmt.Printf("数据：%+v\n", data.Tables)

		for tb, v := range data.Tables {
			if _, err = tx.Table(tb).Delete(idKey, docID); err != nil {
				return err
			}
			tbData := gconv.SliceMap(v)
			for i, v := range tbData {
				tbData[i] = TrimMapNull(v)
				if len(tbData[i]) > 0 {
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
		//	err = gerror.New("保存失败")
		return 0, err
	}
	return docID, nil
}

func TrimMapNull(mapData map[string]interface{}) map[string]interface{} {
	for k, v := range mapData {
		if v == nil {
			delete(mapData, k)
		}
	}
	return mapData
}
