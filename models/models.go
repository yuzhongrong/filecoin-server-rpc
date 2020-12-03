package models

import (
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}
