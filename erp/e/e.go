package e

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// DB erp数据库
var DB = &EDB{gdb: g.DB("erp")}

type EDB struct {
	gdb gdb.DB
}
