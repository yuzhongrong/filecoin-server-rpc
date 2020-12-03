package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type FilIntegral struct {
	Model
	Amount        float64 `json:"amount"`
	Txid          string  `json:"txid"`
	Version       int64   `json:"version"`
	From       string  `json:"address"` //from
	To            string  `json:"to"`      //to
	Direct         string  `json:"direct"`
	State         int64  `json:"state"`//状态0:新交易,1代表已经处理

}


func (t *FilIntegral) FilIntegral(db *gorm.DB) error {
	err := db.Where("txid = ?", t.Txid).FirstOrCreate(&t)
	if err != nil {
		return err.Error
	}
	return nil
}






func (t *FilIntegral) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (t *FilIntegral) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
