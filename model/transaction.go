package model

import "time"

type TransactionStatus int

const (
	Waiting TransactionStatus = iota + 1
	Cancelled
	Paid
)

type Transaction struct {
	ID        string
	UserID    string
	Orders    []Order
	Status    TransactionStatus
	TotalFee  uint
	CreatedAt time.Time
}
