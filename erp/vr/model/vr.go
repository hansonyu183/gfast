// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package model

// Vr is the golang structure for table vr.
type Vr struct {
	ID           uint   `orm:"id,primary"    json:"id"`            //
	No           string `orm:"no,unique"     json:"no"`            //
	Date         string `orm:"date"        json:"date"`            //
	CreateDate   string `orm:"create_date"    json:"createDate"`   //
	CheckDate    string `orm:"check_date"     json:"checkDate"`    //
	CreateUserID uint   `orm:"create_user_id" json:"createUserId"` //
	CheckUserID  uint   `orm:"check_user_id"  json:"checkUserId"`  //
	VtypeID      uint   `orm:"vtype_id"       json:"vtypeId"`      //
	StateID      uint   `orm:"state_id"       json:"stateId"`      //
}
