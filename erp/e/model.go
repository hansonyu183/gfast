package e

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/database/gdb"
)

type Mod struct {
	fields     string
	from       string
	where      map[string]string
	group      string
	order      string
	whereParas map[string]string
	Params     []string
	Proce      string
	offset     int
	limit      int
}

func (db *EDB) From(from string) *Mod {
	mo := new(Mod)
	mo.from = from
	return mo
}

func (mo *Mod) Where(where map[string]string) *Mod {
	mo.where = where
	return mo
}

func (mo *Mod) WhereParas(whereParas map[string]string) *Mod {
	mo.whereParas = whereParas
	return mo
}

func (mo *Mod) Select(fields string) *Mod {
	if mo.fields != "" {
		mo.fields = mo.fields + "," + fields
	} else {
		mo.fields = fields
	}
	return mo
}

func (mo *Mod) Group(group string) *Mod {
	mo.group = group
	return mo
}

func (mo *Mod) Order(order string) *Mod {
	mo.order = order
	return mo
}

func (mo *Mod) Page(page, limit int) *Mod {
	if page <= 0 {
		page = 1
	}
	mo.limit = limit
	mo.offset = limit * page
	return mo
}

func (mo *Mod) sql() (sql string) {
	sql = " from " + mo.from + " where 1 " + mo.whereSQL()
	if mo.group != "" && mo.group != "{}" {
		sql = sql + " group by " + mo.group
	}
	if mo.order != "" && mo.order != "{}" {
		sql = sql + " order by " + mo.order
	}
	if mo.limit > 0 {
		sql = sql + " limit " + strconv.Itoa(mo.offset) + "," + strconv.Itoa(mo.limit)
	}
	return
}

func (mo *Mod) whereSQL() (sql string) {
	for key, val := range mo.where {
		if v, ok := mo.whereParas[key]; ok {
			if v != "" {
				s := strings.Replace(val, "?", v, -1)
				sql = sql + " and " + s
			} //存在
		}
	}
	//fmt.Printf("whereMap%+v  wherePars:%+v  sql:%s", mo.where, mo.whereParas, sql)
	return
}

func (mo *Mod) Count() (c int, e error) {
	sql := "select count(1) " + mo.sql()
	return DB.gdb.GetCount(sql)
}

func (mo *Mod) One() (r gdb.Record, e error) {
	sql := ""
	if mo.fields == "" {
		sql = "select * "
	} else {
		sql = "select " + mo.fields
	}
	sql = sql + mo.sql()
	fmt.Println(sql)
	return DB.gdb.GetOne(sql)
}

func (mo *Mod) All() (r gdb.Result, e error) {
	sql := ""
	if mo.fields == "" {
		sql = "select * "
	} else {
		sql = "select " + mo.fields
	}
	sql = sql + mo.sql()
	return DB.gdb.GetAll(sql)
}
