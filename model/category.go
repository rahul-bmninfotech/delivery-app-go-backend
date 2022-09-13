package model

// CBCategory represents a SaleItem category
type CBCategory struct {
	ID       int    `db:"id" json:"id"`
	Category string `db:"cate" json:"cate"`
}
