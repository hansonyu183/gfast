package boot

import (
	"database/sql"
	"strings"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// ErpDB erp数据库
var ErpDB = &EDB{g.DB("erp")}

type EDB struct {
	gdb.DB
}

type TbNames struct {
	TableName string
}

type TableDesc struct {
	Field   string
	Type    string
	Default string
}

type AssTableName struct {
	MainTableKey    string
	MiddelTableName string
	TargetTableName string
}

func (eb *EDB) GetFKTbsName(table string) (names []AssTableName, err error) {
	like := "fk\\_%" + table + "%"
	sql := "SELECT TABLE_NAME FROM	information_schema.TABLES WHERE	table_schema = 'kyerp'	AND TABLE_NAME LIKE '" + like + "'"
	ns, err := eb.GetArray(sql)
	if err != nil {
		return
	}
	for _, v := range ns {
		midTb := v.String()
		tagTb := strings.TrimPrefix(midTb, "fk_")
		tagTb = strings.TrimPrefix(tagTb, table+"_")
		tagTb = strings.TrimSuffix(tagTb, "_"+table)
		names = append(names, AssTableName{
			table + "_id",
			midTb,
			tagTb,
		})
	}
	return
}

func (eb *EDB) GetSubTbsName(table string) (names []AssTableName, err error) {
	like := table + "\\_%"
	sql := "SELECT TABLE_NAME FROM	information_schema.TABLES WHERE	table_schema = 'kyerp'	AND TABLE_NAME LIKE '" + like + "'"
	ns, err := eb.GetArray(sql)
	if err != nil {
		return
	}
	for _, v := range ns {
		names = append(names, AssTableName{
			table + "_id",
			"",
			v.String(),
		})
	}
	return
}

func (eb *EDB) GetTbDefaultData(tbName string) (data gdb.Record, err error) {
	var desc []TableDesc
	err = eb.GetStructs(&desc, "desc "+tbName)
	//desc,err:=tx.GetAll("desc "+tbName)
	if err != nil {
		return nil, err
	}
	//data = make(gdb.Record)
	sql := "select "
	for _, v := range desc {
		//data[v.Field] = gvar.New(v.Default, true)
		sql += "default(" + v.Field + ") as " + v.Field + ","
	}
	sql = strings.TrimSuffix(sql, ",")
	sql += " from " + tbName + " limit 1"
	data, err = eb.GetOne(sql)
	return
}

func (eb *EDB) GeTbFkSub(tbName string, id int) (tbData gdb.Record, fksData map[string]gdb.Result, subsData map[string]gdb.Result, err error) {
	err = eb.Transaction(func(tx *gdb.TX) error {
		if id == 0 { //新建记录
			tbData, err = eb.GetTbDefaultData(tbName)
		} else {
			tbData, err = tx.Table(tbName).Where("id", id).One()
		}
		if err != nil {
			return err
		}
		if fksData, err = eb.GetFKData(tx, tbName, id); err != nil {
			return err
		}
		if subsData, err = eb.GetSubData(tx, tbName, id); err != nil {
			return err
		}
		return err
	})
	return
}

func (eb *EDB) GetFKData(tx *gdb.TX, table string, id int) (fksData map[string]gdb.Result, err error) {
	assTbs, err := eb.GetFKTbsName(table)
	if err != nil {
		return
	}
	fksData = make(map[string]gdb.Result)
	for _, v := range assTbs {
		var tagIds gdb.Result
		if id == 0 {
			tagIds = nil
		} else {
			tagIds, err = tx.Table(v.MiddelTableName).Fields(v.TargetTableName+"_id").Where(v.MainTableKey, id).All()
		}
		if err == nil {
			fksData[v.MiddelTableName] = tagIds
		}
		//}
	}
	return
}

func (eb *EDB) GetSubData(tx *gdb.TX, table string, id int) (subsData map[string]gdb.Result, err error) {
	assTbs, err := eb.GetSubTbsName(table)
	if err != nil {
		return
	}

	subsData = make(map[string]gdb.Result)
	for _, v := range assTbs {
		var r gdb.Result
		if id == 0 {
			r = nil
		} else {
			r, err = tx.Table(v.TargetTableName).Where(v.MainTableKey, id).All()
		}
		if err == nil {
			subsData[v.TargetTableName] = r
		}
	}
	return
}

func (eb *EDB) DelTbFkSub(table string, id int) (result sql.Result, err error) {
	err = eb.Transaction(func(tx *gdb.TX) error {
		result, err = tx.Table(table).Delete("id", id)
		if err != nil {
			return err
		}
		if err = eb.DelFKData(tx, table, id); err != nil {
			return nil
		}
		if err = eb.DelSubData(tx, table, id); err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return
	}
	return
}

func (eb *EDB) DelFKData(tx *gdb.TX, table string, id int) (err error) {
	preFix := "fk_" + table + "_"
	sql := "SELECT TABLE_NAME FROM	information_schema.TABLES WHERE	table_schema = 'kyerp'	AND TABLE_NAME LIKE '" + preFix + "%'"
	tbs, err := eb.GetArray(sql)
	if err != nil {
		return
	}
	for _, val := range tbs {
		fkTb := val.String()
		_, err := tx.Table(fkTb).Delete(table+"_id", id)
		if err != nil {
			continue
		}

	}
	return err
}

func (eb *EDB) DelSubData(tx *gdb.TX, table string, id int) (err error) {
	preFix := table + "_"
	sql := "SELECT TABLE_NAME FROM	information_schema.TABLES WHERE	table_schema = 'kyerp'	AND substr(TABLE_NAME,1,4) ='" + preFix + "'"
	tbs, err := eb.GetArray(sql)
	if err != nil {
		return
	}
	for _, val := range tbs {
		tb := val.String()
		_, err := tx.Table(tb).Delete(table+"_id", id)
		if err == nil {
			continue
		}
	}
	return err
}
