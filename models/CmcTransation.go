package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CmcTransation struct {
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

	//Vin           []Vin  `json:"vin"`
	//Vout          []Vout `json:"vout"`
}

type ScriptP struct {
	Addresses []string `json:"addresses"`
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int64    `json:"reqSigs"`
	Type      string   `json:"type"`
}

type ScriptS struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type Vout struct {
	N            int64   `json:"n"`
	ScriptPubKey ScriptP `json:"scriptPubKey"`
	Value        float64 `json:"value"`
}

type Vin struct {
	ScriptSig ScriptS `json:"scriptSig"`
	Sequence  int64   `json:"sequence"`
	Txid      string  `json:"txid"`
	Vout      int64   `json:"vout"`
}

func (t *CmcTransation) AddTransation(db *gorm.DB) error {
	err := db.Where("txid = ?", t.Txid).FirstOrCreate(&t)
	if err != nil {
		return err.Error
	}
	return nil
}

func ExistTransationByID(db *gorm.DB, txId string, blockhash string) bool {
	var transation CmcTransation
	db.Where("txid = ? AND blockhash = ?", txId, blockhash).Find(&transation)
	return transation.Txid != "" && transation.Blockhash != ""

}

func (t *CmcTransation) GetLastTransation(db *gorm.DB) (string, error) {

	//db.Model(&CmcTransation{}).Where("id","id?=")
	//return err
	return "", nil

}

func (t *CmcTransation) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (t *CmcTransation) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
