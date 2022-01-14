package model

type Order struct {
	ID            string
	UserID        string
	TransactionID string
	ProductName   string
	Price         uint
	Quantity      uint
}
