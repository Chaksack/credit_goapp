package models

type Bank struct {
	Id            uint   `json:"id"`
	BankName      string `json:"bank_name"`
	AccountHolder string `json:"account_holder"`
	AccountNumber int    `json:"account_number"`
}

type Momo struct {
	Id           uint   `json:"id"`
	Network      string `json:"network"`
	MobileNumber int    `json:"mobile_number"`
}
