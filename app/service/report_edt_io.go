package service

import (
	"fmt"
	"gfast/app/dao"
	"gfast/app/model"
)

// ReportEdtIO 中间件管理服务
var ReportEdtIO = new(reportEdtIOService)

type reportEdtIOService struct{}

// 用户注册
func (s *reportEdtIOService) Get(r *model.ReportEdtIoApiReq) (data interface{}, err error) {
	offset := 0
	limit := 20
	if r.Page == 0 {
		return dao.ReportEdtIo.Columns, nil
	}
	offset = (r.Page - 1) * limit

	pa := fmt.Sprintf("%s,%d,%d", r.ToSQL(), offset, limit)
	if data, err = dao.ReportEdtIo.Get(pa); err != nil {
		//		if data, err = dao.ReportEdtIo.Get(r.ResID,r.BegDate,r.EndDate,r.SumKind,offset,limit); err != nil {
		return nil, err
	}
	return
}
