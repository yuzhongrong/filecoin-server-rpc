package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"log"
)

type CmcTransationHistory struct {
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
	Address       string  `json:"address"`
	//Vin           string  `json:"vin"`
	//Vout          string `json:"vout"`

}



func (t *CmcTransationHistory) InsertTransation(db *gorm.DB) error {

	log.Println("-------开始插入一笔交易历史-----> ", t.Blockhash,t.Txid)
	if err := db.Where("blockhash = ? and txid = ?", t.Blockhash,t.Txid).FirstOrCreate(&t);err!=nil{
		return err.Error
	}
	return nil

}


func (t *CmcTransationHistory) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (t *CmcTransationHistory) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
