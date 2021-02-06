package e

import (
	"strings"

	"github.com/gogf/gf/database/gdb"
)

type Proce struct {
	proce     string
	params    []string
	paramsVal map[string]string
	offset    int
	limit     int
}

func (db *EDB) Proce(procName string, params []string) *Proce {
	pr := new(Proce)
	pr.proce = procName
	pr.params = params
	return pr
}

func (pr *Proce) ParamsVal(paramsVal map[string]string) *Proce {
	pr.paramsVal = paramsVal
	return pr
}

func (pr *Proce) paramsSQL() (sql string) {
	for _, val := range pr.params {
		sql = sql + "'" + pr.paramsVal[val] + "',"
	}
	sql = strings.TrimSuffix(sql, ",")
	//fmt.Printf("whereMap%+v  wherePars:%+v  sql:%s", mo.where, mo.whereParas, sql)
	return
}

func (pr *Proce) Call(where map[string]string) (r gdb.Result, e error) {
	sql := "call " + pr.proce + "(" + pr.paramsSQL() + ")"
	return DB.gdb.GetAll(sql)
}
