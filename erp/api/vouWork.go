package api

import (
	"database/sql"
	"gfast/library/response"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

var workTbMod *gdb.Model

type VouWork struct {
}

type WorkParam struct {
	VoucherIds []int
}

//controller
func (w *VouWork) Check(r *ghttp.Request) {
	wp := &WorkParam{}
	if err := r.Parse(wp); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	sr, err := Check(wp.VoucherIds)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	ef, _ := sr.RowsAffected()
	if ef == 0 {
		response.FailJson(true, r, "处理失败")
	}

	response.SusJson(true, r, "处理成功")
}

func (w *VouWork) UnCheck(r *ghttp.Request) {
	wp := &WorkParam{}
	if err := r.Parse(wp); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	sr, err := UnCheck(wp.VoucherIds)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	ef, _ := sr.RowsAffected()
	if ef == 0 {
		response.FailJson(true, r, "处理失败")
	}

	response.SusJson(true, r, "处理成功")
}

func (w *VouWork) Finish(r *ghttp.Request) {
	wp := &WorkParam{}
	if err := r.Parse(wp); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	sr, err := Finish(wp.VoucherIds)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	ef, _ := sr.RowsAffected()
	if ef == 0 {
		response.FailJson(true, r, "处理失败")
	}

	response.SusJson(true, r, "处理成功")
}

func (w *VouWork) UnFinish(r *ghttp.Request) {
	wp := &WorkParam{}
	if err := r.Parse(wp); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	sr, err := UnFinish(wp.VoucherIds)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	ef, _ := sr.RowsAffected()
	if ef == 0 {
		response.FailJson(true, r, "处理失败")
	}

	response.SusJson(true, r, "处理成功")
}

func (w *VouWork) Abandon(r *ghttp.Request) {
	wp := &WorkParam{}
	if err := r.Parse(wp); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	sr, err := Abandon(wp.VoucherIds)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	ef, _ := sr.RowsAffected()
	if ef == 0 {
		response.FailJson(true, r, "处理失败")
	}

	response.SusJson(true, r, "处理成功")
}

func (w *VouWork) UnAbandon(r *ghttp.Request) {
	wp := &WorkParam{}
	if err := r.Parse(wp); err != nil {
		response.FailJson(true, r, err.(*gvalid.Error).FirstString())
	}
	sr, err := UnAbandon(wp.VoucherIds)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	ef, _ := sr.RowsAffected()
	if ef == 0 {
		response.FailJson(true, r, "处理失败")
	}

	response.SusJson(true, r, "处理成功")
}

func Check(ids []int) (sql.Result, error) {
	return workTbMod.Where("vouhcer_id", ids).Where("state", "").Update("state", "c")
}

func UnCheck(ids []int) (sql.Result, error) {
	return workTbMod.Where("vouhcer_id", ids).Where("state", "c").Update("state", "")
}

func Finish(ids []int) (sql.Result, error) {
	return workTbMod.Where("vouhcer_id", ids).Where("state", "c").Update("state", "f")
}

func UnFinish(ids []int) (sql.Result, error) {
	return workTbMod.Where("vouhcer_id", ids).Where("state", "f").Update("state", "c")
}

func Abandon(ids []int) (sql.Result, error) {
	return workTbMod.Where("vouhcer_id", ids).Where("state", "").Update("state", "a")
}

func UnAbandon(ids []int) (sql.Result, error) {
	return workTbMod.Where("vouhcer_id", ids).Where("state", "a").Update("state", "")
}

func init() {
	workTbMod = g.DB("erp").Table("vouhcer")
}

func NewWork() *VouWork {
	return &VouWork{}
}
