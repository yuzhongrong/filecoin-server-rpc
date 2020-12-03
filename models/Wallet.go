package models

type Wallet struct {

	ID            int     `gorm:"primary_key" json:"id"`
	Address     string  `json:"address"`
	Outaddress     string  `json:"outaddress"`
	Starthight  int64   `json:"starthight"`
	Current int64 `json:"current"`//当前查询到哪个块


}

