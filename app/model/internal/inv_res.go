// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal



// InvRes is the golang structure for table inv_res.
type InvRes struct {
    Id      int     `orm:"id,primary" json:"id"`       //   
    Name    string  `orm:"name"       json:"name"`     //   
    Model   string  `orm:"model"      json:"model"`    //   
    TaxNo   string  `orm:"tax_no"     json:"tax_no"`   //   
    Py      string  `orm:"py"         json:"py"`       //   
    State   string  `orm:"state"      json:"state"`    //   
    InitNum float64 `orm:"init_num"   json:"init_num"` //   
    EndNum  float64 `orm:"end_num"    json:"end_num"`  //   
}