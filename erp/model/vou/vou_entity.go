// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package vou

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table vou.
type Entity struct {
    Id            int    `orm:"id,primary"      json:"id"`              //   
    No            string `orm:"no,unique"       json:"no"`              //   
    Date          string `orm:"date"            json:"date"`            //   
    Type          string `orm:"type"            json:"type"`            //   
    Note          string `orm:"note"            json:"note"`            //   
    State         string `orm:"state"           json:"state"`           //   
    CreateDate    string `orm:"create_date"     json:"create_date"`     //   
    CreateUserId  int    `orm:"create_user_id"  json:"create_user_id"`  //   
    UpdateDate    string `orm:"update_date"     json:"update_date"`     //   
    UpdateDateId  int    `orm:"update_date_id"  json:"update_date_id"`  //   
    CheckDate     string `orm:"check_date"      json:"check_date"`      //   
    CheckUserId   int    `orm:"check_user_id"   json:"check_user_id"`   //   
    FinishDate    string `orm:"finish_date"     json:"finish_date"`     //   
    FinishUserId  int    `orm:"finish_user_id"  json:"finish_user_id"`  //   
    AbandonDate   string `orm:"abandon_date"    json:"abandon_date"`    //   
    AbandonUserId int    `orm:"abandon_user_id" json:"abandon_user_id"` //   
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}