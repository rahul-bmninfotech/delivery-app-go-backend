package model

import "time"

type Demand struct {
	Id            int       `db:"id" json:"id"`
	DemandNumber  string    `db:"dnum" json:"dnum"`
	DemandDate    time.Time `db:"ddate" json:"ddate"`
	SaleItem      int       `db:"sitem" json:"sitem"`
	ItemCategory  string    `db:"itemcategory" json:"itemcategory"`
	SaleUnit      string    `db:"unit" json:"unit"`
	Rate          string    `db:"rate" json:"rate"`
	Quantity      int       `db:"qty" json:"qty"`
	Amount        string    `db:"amount" json:"amount"`
	Status        string    `db:"status" json:"status"`
	IssueQuantity int       `db:"issueqty" json:"issueqty"`
	VehicleId     int       `db:"vehicle" json:"vehicle"`
}

type LoadedDemand struct {
	Id            int       `db:"id" json:"id"`
	DemandNumber  string    `db:"dnum" json:"dnum"`
	DemandDate    time.Time `db:"ddate" json:"ddate"`
	SaleUnit      string    `db:"unit" json:"unit"`
	Rate          string    `db:"rate" json:"rate"`
	Quantity      int       `db:"qty" json:"qty"`
	Amount        string    `db:"amount" json:"amount"`
	Status        string    `db:"status" json:"status"`
	IssueQuantity int       `db:"issueqty" json:"issueqty"`
	Vehicle       `json:"vehicle"`
	CBCategory    `json:"itemcategory"`
	SaleItem      `json:"saleitem"`
}
