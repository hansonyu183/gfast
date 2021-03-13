package doc

import (
	"fmt"
	"gfast/erp/boot"
	"strings"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gdb"
)

type ResponseGet struct {
	NextID *gvar.Var             `json:"nextID"`
	PreID  *gvar.Var             `json:"preID"`
	Forms  map[string]gdb.Record `json:"forms"`
	Tables map[string]gdb.Result `json:"tables"`
}

//getData
/**
获取明细
*/
func (ctrl *doc) getData(docType string, docID int) (data *ResponseGet, err error) {
	tb, fks, subs, err := boot.ErpDB.GeTbFkSub(docType, docID)
	if err != nil {
		return nil, err
	}
	data = &ResponseGet{}
	if tb != nil {
		data.Forms = make(map[string]gdb.Record)
		if docID == 0 {
			no, _ := MakeDocNewNo(docType)
			tb["no"] = gvar.New(no, true)
		}
		data.Forms[docType] = tb
	}
	tbs := make(map[string]gdb.Result)
	if fks != nil {
		tbs = fks
	}
	if subs != nil {
		for k, v := range subs {
			tbs[k] = v
		}
	}
	if len(tbs) > 0 {
		data.Tables = tbs
	}
	if docID != 0 {
		data.NextID, err = boot.ErpDB.Table(docType).Where("id>?", docID).Value("min(id)")
		if err != nil {
			return nil, err
		}
		data.PreID, err = boot.ErpDB.Table(docType).Where("id<?", docID).Value("max(id)")

		if err != nil {
			return nil, err
		}
	} else {
		data.PreID, err = boot.ErpDB.Table(docType).Value("max(id)")
		if err != nil {
			return nil, err
		}
	}
	return data, err
}

func MakeDocNewNo(docType string) (newNo string, err error) {
	maxID, err := boot.ErpDB.Model(docType).Value("max(id)")
	if err != nil {
		return "", err
	}
	newNo = fmt.Sprintf("%s%d%s", "K", maxID.Int(), strings.ToUpper(docType[0:1]))
	return
}
