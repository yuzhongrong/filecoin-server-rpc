package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type CmcRecord struct {
	ID            int     `gorm:"primary_key" json:"id"`
	Amount        float64 `json:"amount"`
	NowAllamount  float64 `json:"now_allamount"`
	LastAllamount float64 `json:"last_allamount"`
	Blockhash     string  `json:"blockhash"`
	Timestamp     int64   `json:"timestamp"`
	From          string  `json:"from"` //from
	To            string  `json:"to"`   //to
	Txid          string  `json:"txid"`

}

func (t *CmcRecord) RelateTransation(db *gorm.DB) {
	log.Println("------RelateTransation-----> ")
}
