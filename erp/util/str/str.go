package str

import (
	"fmt"
	"reflect"
	"strings"
)

// CovMapKeyToArray 转换map描述字符串，输出Key的数组 {"k1":"v","k2":"v"}=>[k1,k2]
func CovMapKeyToArray(mapStr string) (keyStrArray []string) {
	mapStr = strings.TrimPrefix(mapStr, "{")
	strArr := strings.Split(mapStr, ",")
	for _, v := range strArr {
		i := strings.Index(v, ":")
		if i >= 1 {
			keyStrArray = append(keyStrArray, v[1:i-1])
		}
	}
	return
}

// ParamsSQL 转换struct 到 sql存储过程的参数数
func ParamsSQL(params interface{}) (sql string) {
	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)
	for k := 0; k < t.NumField(); k++ {
		sql = sql + fmt.Sprintf("%v", v.Field(k).Interface()) + ","
	}
	sql = strings.TrimSuffix(sql, ",")
	return
}
