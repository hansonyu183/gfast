// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal



// Account is the golang structure for table account.
type Account struct {
    Id           string  `orm:"id,primary"     json:"id"`             //   
    Name         string  `orm:"name"           json:"name"`           //   
    Amo          float64 `orm:"amo"            json:"amo"`            //   
    InitAmo      float64 `orm:"init_amo"       json:"init_amo"`       //   
    Note         string  `orm:"note"           json:"note"`           //   
    CreateDate   string  `orm:"create_date"    json:"create_date"`    //   
    CreateUserId int     `orm:"create_user_id" json:"create_user_id"` //   
    State        string  `orm:"state"          json:"state"`          //   
    SubjectId    string  `orm:"subject_id"     json:"subject_id"`     //   
}