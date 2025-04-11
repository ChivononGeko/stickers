package models

type Order struct {
	ClientName string         `json:"clientName"`
	CloseTime  string         `json:"closeTime"`
	Products   []OrderProduct `json:"products"`
}

type OrderProduct struct {
	Comment       string `json:"comment"`
	Count         int    `json:"count"`
	ID            string `json:"id"`
	Modifications string `json:"modifications"`
}

type OrderModification struct {
	M int `json:"m"`
	A int `json:"a"`
}
