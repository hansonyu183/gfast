package doc

import (
	"gfast/erp/boot"

	"gfast/library/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// DocOption API管理对象
var DocOption = &docOption{}

type docOption struct {
	name string
}

type DocOptionParam struct {
}

//controller
func (ctrl *docOption) Get(r *ghttp.Request) {
	docType := r.GetString("type")

	data, err := ctrl.getData(docType)

	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", data)
}

//getData
/**
获取明细
*/

func (ctrl *docOption) getData(docType string) (data interface{}, err error) {
	if docType == "ALL" {
		return ctrl.getAllOption()
	}

	mod := boot.ErpDB.Table(docType).Fields("id,no,name,py")
	r, err := mod.All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}

	data = g.Map{
		docType: r,
	}
	return
}

func (ctrl *docOption) getAllOption() (data map[string]interface{}, er error) {
	db := boot.ErpDB
	fdb := g.DB()
	cols := "id,no,name,py"
	where := "state_id=4"
	docs := map[string]string{
		"eba":     cols,
		"sup":     cols,
		"res":     cols,
		"emp":     cols,
		"account": cols,
		"subject": cols,
		"vtype":   cols,
		"brand":   cols,
		"reskind": cols,
		"state":   "*",
		"act":     "*",
		"unit":    cols,
		"ebasq":   cols,
		"invres":  cols,
		"inveba":  cols,
	}

	data = g.Map{}

	for k, v := range docs {
		r, err := db.Table(k).Fields(v).Where(where).All()
		if err == nil {
			data[k] = r
		}

	}
	user, _ := fdb.Table("user").Fields("id,user_name as no,user_nickname as name,'' as py").All()
	data["user"] = user
	return
}
