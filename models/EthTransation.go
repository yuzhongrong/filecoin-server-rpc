package models

type EthTransation struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Result []Result `json:"result"`
}
type Result struct {
	BlockNumber string `json:"blockNumber"`
	TimeStamp string `json:"timeStamp"`
	Hash string `json:"hash"`
	Nonce string `json:"nonce"`
	BlockHash string `json:"blockHash"`
	TransactionIndex string `json:"transactionIndex"`
	From string `json:"from"`
	To string `json:"to"`
	Value string `json:"value"`
	Gas string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	IsError string `json:"isError"`
	TxreceiptStatus string `json:"txreceipt_status"`
	Input string `json:"input"`
	ContractAddress string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed string `json:"gasUsed"`
	Confirmations string `json:"confirmations"`
}
