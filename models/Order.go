package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Order struct {
	Id            int     `gorm:"primary_key" json:"id"`
	Orderid        string `json:"orderid"`
	Txid          string  `json:"txid"`
	Bindaddress string `json:"bindaddress"`
	Amount     float64  `json:"amount"`
	State  int `json:"state"`


}

func (o *Order) Order(db *gorm.DB) {
	log.Println("------RelateTransation-----> ")
}

//func (o *Order) GetLastOrder(db *gorm.DB){
//	db.Last(&o,"pay_status=?",1)
//}
