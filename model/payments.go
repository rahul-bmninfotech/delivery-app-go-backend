package model

import (
	"time"
)

// Payments stores sales which the buyer took on credit
type Payments struct {
	ID          int       `db:"id" json:"id"`
	SaleID      int       `db:"sale_id" json:"sale_id"`
	AddedDate   time.Time `db:"added_date" json:"added_date"`
	PaidDate    time.Time `db:"paid_date" json:"paid_date"`
	Balance     float32   `db:"balance" json:"balance"`
	IsPaid      int8      `db:"is_paid" json:"is_paid"`
	BuyerID     int       `db:"buyer_id" json:"buyer_id"`
	SaleGroupID int       `db:"salegroup_id" json:"salegroup_id"` // Sales added together will have the same group id
}
