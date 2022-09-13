package model

import "time"

// Sale represents a sale made to a customer
type Sale struct {
	Id              int       `db:"id" json:"id"`
	DemandNumber    string    `db:"dnum" json:"dnum"`
	DemandDate      time.Time `db:"ddate" json:"ddate"`
	Route           string    `db:"route" json:"route"`
	Vehicle         string    `db:"vehicle" json:"vehicle"`
	Driver          string    `db:"driver" json:"driver"`
	Buyer           string    `db:"buyer" json:"buyer"`
	SaleItem        int       `db:"sitem" json:"sitem"`
	Quantity        string    `db:"qty" json:"qty"`
	QuantityOut     int       `db:"qtyout" json:"qtyout"`
	QuantityBalance int       `db:"qtybal" json:"qtybal"`
	IssuedDate      time.Time `db:"idate" json:"idate"`
	Status          string    `db:"status" json:"status"`
	Invoice         string    `db:"invoice" json:"invoice"`
	Credit          string    `db:"credit" json:"credit"`
	SalePrice       string    `db:"sale_price" json:"sale_price"`
	InvoiceNumber   string    `db:"invoice_no" json:"invoice_no"`
	HasVat          bool      `db:"has_vat" json:"has_vat"`
}

// Returned represents returned sales
type Returned struct {
	Id              int       `db:"id" json:"id"`
	DemandNumber    string    `db:"dnum" json:"dnum"`
	DemandDate      time.Time `db:"ddate" json:"ddate"`
	Route           string    `db:"route" json:"route"`
	Vehicle         string    `db:"vehicle" json:"vehicle"`
	Driver          string    `db:"driver" json:"driver"`
	Buyer           string    `db:"buyer" json:"buyer"`
	SaleItem        int       `db:"sitem" json:"sitem"`
	Quantity        int       `db:"qty" json:"qty"`
	QuantityOut     int       `db:"qtyout" json:"qtyout"`
	QuantityBalance int       `db:"qtybal" json:"qtybal"`
	IssuedDate      time.Time `db:"idate" json:"idate"`
	Status          string    `db:"status" json:"status"`
}

type SaleItemPriority struct {
	Id       int `db:"id" json:"id"`
	ItemId   int `db:"item_id" json:"item_id"`
	Priority int `db:"priority" json:"priority"`
}

// SaleItem represents an item for sale
type SaleItem struct {
	Id             int    `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	ItemUnit       string `db:"item_unit" json:"item_unit"`
	SalePrice      string `db:"sale_price" json:"sale_price"`
	PurchasedPrice string `db:"purchased_price" json:"purchased_price"`
	Description    string `db:"descr" json:"descr"`
	ItemCategory   string `db:"itemcategory" json:"itemcategory"`
	ImagePath      string `db:"img" json:"img"`
}

type SaleGroupInvoice struct {
	ID          int    `db:"salegroup_id" json:"salegroup_id"`
	InvoiceName string `db:"invoice" json:"invoice"`
}
