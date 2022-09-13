package model

type Buyer struct {
	Id            int    `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	Address       string `db:"address" json:"address"`
	Latitude      string `db:"latitude" json:"latitude"`
	Longitude     string `db:"longitude" json:"longitude"`
	RouteId       int    `db:"route_id" json:"route_id"`
	ContactNumber string `db:"contact_no" json:"contact_no"`
	Email         string `db:"email" json:"email"`
	Priority      int    `db:"priority" json:"priority"`
}

type BuyerItemPrice struct {
	ID           int    `db:"id" json:"id"`
	ItemID       int    `db:"item_id" json:"item_id"`
	BuyerID      int    `db:"buyer_id" json:"buyer_id"`
	Price        string `db:"price" json:"price"`
	ItemCategory string `db:"item_category" json:"item_category"`
}

type BuyerPriority struct {
	Id       int `db:"id" json:"id"`
	DriverID int `db:"driver_id" json:"driver_id"`
	BuyerID  int `db:"buyer_id" json:"buyer_id"`
	Priority int `db:"priority" json:"priority"`
}

type BuyerPriorities struct {
	Id         int    `db:"id" json:"id"`
	DriverID   int    `db:"driver_id" json:"driver_id"`
	RouteID    int    `db:"route_id" json:"route_id"`
	Priorities string `db:"priorities" json:"priorities"`
}
