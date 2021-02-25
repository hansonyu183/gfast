package doc

import (
	"fmt"
	"gfast/erp/boot"
	"strings"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

//getData
/**
获取明细
*/
func (ctrl *doc) getData(docType string, docID int) (data map[string]interface{}, err error) {
	tb, fks, subs, err := boot.ErpDB.GeTbFkSub(docType, docID)
	if err != nil {
		return nil, err
	}
	data = g.Map{}
	if tb != nil {
		if docID == 0 {
			no, _ := MakeDocNewNo(docType)
			tb["no"] = gvar.New(no, true)
		}
		data["form"] = tb
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
		data["tables"] = tbs
	}
	var nextID, preID *g.Var
	if docID != 0 {
		nextID, err = boot.ErpDB.Table(docType).Where("id>?", docID).Value("min(id)")
		if err != nil {
			return nil, err
		}
		preID, err = boot.ErpDB.Table(docType).Where("id<?", docID).Value("max(id)")

		if err != nil {
			return nil, err
		}
	} else {
		preID, err = boot.ErpDB.Table(docType).Value("max(id)")
		if err != nil {
			return nil, err
		}
	}
	data["nextId"] = nextID
	data["preId"] = preID
	return
}

func MakeDocNewNo(docType string) (newNo string, err error) {
	maxID, err := boot.ErpDB.Model(docType).Value("max(id)")
	if err != nil {
		return "", err
	}
	newNo = fmt.Sprintf("%s%d%s", "K", maxID.Int(), strings.ToUpper(docType[0:1]))
	return
}
