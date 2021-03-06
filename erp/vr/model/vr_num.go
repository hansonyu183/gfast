// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package model

// VrNum is the golang structure for table vr_num.
type VrNum struct {
	ID         uint    `orm:"id,primary"   json:"id"`         //
	VID        uint    `orm:"vid,primary"  json:"vid"`        //
	PackNum    float64 `orm:"pack_num"     json:"packNum"`    //
	PerPackNum float64 `orm:"per_pack_num" json:"perPackNum"` //
	Num        float64 `orm:"num"          json:"num"`        //
	Price      float64 `orm:"price"        json:"price"`      //
	Amo        float64 `orm:"amo"          json:"amo"`        //
	TaxRate    float64 `orm:"tax_rate"     json:"taxRate"`    //
	TaxAmo     float64 `orm:"tax_amo"      json:"taxAmo"`     //
	CostPrice  float64 `orm:"cost_price"   json:"costPrice"`  //
	CostAmo    float64 `orm:"cost_amo"     json:"costAmo"`    //
	ItNote     string  `orm:"it_note"      json:"itNote"`     //
	RefVID    uint    `orm:"ref_vid"    json:"refVId"`    //
	RefIID    uint    `orm:"ref_iid"    json:"refIId"`    //
	RefVrNo    string  `orm:"ref_vr_no"    json:"refVrNo"`    //
	EdtIo      string  `orm:"edt_io"       json:"edtIo"`      //
	InvIo      string  `orm:"inv_io"       json:"invIo"`      //
	QtIo       string  `orm:"qt_io"        json:"qtIo"`       //
	ResID      uint    `orm:"res_id"       json:"resId"`      //
	InvresID   uint    `orm:"invres_id"    json:"invresId"`   //
	CompanyID  uint    `orm:"company_id"   json:"companyId"`  //
	PrintRes   string  `orm:"print_res"    json:"printRes"`   //
	PrintModel string  `orm:"print_model"  json:"printModel"` //
	KpiNum     float64 `orm:"kpi_num"      json:"kpiNum"`     //
	KpiAmo     float64 `orm:"kpi_amo"      json:"kpiAmo"`     //
}
