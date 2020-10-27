package api

import (
	//"fmt"
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
)

type Dict struct {
	DictMods map[string]*gdb.Model
}

//controller
func (o *Dict) Get(r *ghttp.Request) {
	dict, err := GetDict(o.DictMods)

	if err != nil {
		response.FailJson(true, r, err.Error(), dict)
	}

	response.SusJson(true, r, "成功", dict)

}

func GetDict(DictMods map[string]*gdb.Model) (dict map[string]gdb.Result, err error) {
	dict = make(map[string]gdb.Result)
	var dbErr error
	for k, v := range DictMods {
		dict[k], dbErr = v.All()
		if dbErr != nil {
			gerror.Wrap(err, dbErr.Error())
		}
	}
	return
}
