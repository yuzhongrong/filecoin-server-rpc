package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type FilTransation struct {
	Model
	Coinbase      bool    `json:"coinbase"`
	BlockHight    uint64  `json:"blockhight"`
	Blockhash     string  `json:"blockhash"`
	Blocktime     int64   `json:"blocktime"`
	Confirmations int64   `json:"confirmations"`
	Hex           string  `json:"hex"`
	Amount        float64 `json:"amount"`
	Fee           float64 `json:"fee"`
	Category      string  `json:"category"`
	Locktime      int64   `json:"locktime"`
	Size          int64   `json:"size"`
	Time          int64   `json:"time"`
	Txid          string  `json:"txid"`
	Version       int64   `json:"version"`
	Generated     bool    `json:"generated"`
	Address       string  `json:"address"` //from
	To            string  `json:"to"`      //to
	Direct         string  `json:"direct"`

}



func (t *FilTransation) AddTransation(db *gorm.DB) error {
	err := db.Where("txid = ?", t.Txid).FirstOrCreate(&t)
	if err != nil {
		return err.Error
	}
	return nil
}


func (t *FilTransation)RelatedFilIntegral(db *gorm.DB) error {
	db.Where("txid = ?", t.Txid).FirstOrCreate(&FilIntegral{
		Txid:          t.Txid,
		Amount:        t.Amount,
		From:          t.Address,
		To:            t.To,
		State:         0,
		Direct:        t.Direct,

	})
	return nil
}

func (t *FilTransation) FindOneTransation(db *gorm.DB) (string,error) {
	err := db.Where("txid = ?", t.Txid).Find(&t)
	if err != nil {
		return "",err.Error
	}
	return t.Txid,nil
}


func (t *FilTransation) GetLastTransation(db *gorm.DB) (string, error) {

	//db.Model(&CmcTransation{}).Where("id","id?=")
	//return err
	return "", nil

}


func (t *FilTransation) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (t *FilTransation) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
