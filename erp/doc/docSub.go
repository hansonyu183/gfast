package doc

import (
	"fmt"
	"gfast/erp/boot"
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

// DocSub API管理对象
var DocSub = &docSub{}

type docSub struct {
}

type PostSubData struct {
	Tables map[string]string
	KeyIDs map[string]string
}

//controller
func (ctrl *docSub) Post(r *ghttp.Request) {
	var docSubData *PostSubData
	var err error
	if err = r.Parse(&docSubData); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	docID := r.GetInt("docID")
	if docID == 0 {
		response.FailJson(true, r, "无主档案")
	}
	docType := r.GetString("docType")

	err = ctrl.saveData(docType, docID, docSubData)

	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功")
}

func (ctrl *docSub) saveData(docType string, docID int, data *PostSubData) (err error) {
	err = boot.ErpDB.Transaction(func(tx *gdb.TX) error {
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

			if _, err = tx.Model(tb).Insert(tbData); err != nil {
				return err
			}

		}
		return nil
	})
	if err != nil {
		g.Log().Error(err)
		//	err = gerror.New("保存失败")
		return err
	}
	return nil
}
