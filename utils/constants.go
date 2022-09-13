package utils

// Request Query Params
const (
	QueryUsername               = "username"
	QueryPassword               = "password"
	QueryDriverId               = "driver_id"
	QueryVehicleId              = "vehicle_id"
	QueryVehicleComment         = "comment"
	QueryRouteId                = "route_id"
	QueryBuyerId                = "buyer_id"
	QueryItemId                 = "item_id"
	QueryDemandNumber           = "demand_number"
	QueryItemQuantity           = "item_qty"
	QuerySaleItemId             = "sale_item_id"
	QuerySaleId                 = "sale_id"
	QueryInvoiceSaleIds         = "ids"
	QueryPaymentsBuyersID       = "buyers"
	QueryPaymentsBuyerID        = "buyer"
	QueryPaymentsSaleIDs        = "ids"
	QueryPaymentsBalance        = "bal"
	QueryPaymentGroupIDs        = "group_ids"
	QueryBuyerPriceIDs          = "buyer_id"
	QueryBuyerPriceItemID       = "i_id"
	QueryBuyerPriceBuyerID      = "b_id"
	QueryBuyerPriceValue        = "price"
	QueryBuyerPriceCategory     = "cate"
	QueryBuyerPriceID           = "id"
	QueryBuyerPriceNewValue     = "price"
	QueryInvoiceNumberParameter = "invoice_no"
)

// SQL Query Statements
const (
	DriverLoginQuery        = "SELECT id, username, password, address, contact_no, name, email, dob, doj, licenseno FROM driver"
	VehicleQuery            = "SELECT id, vehicle_no, comments, is_taken FROM vehicle"
	RouteQuery              = "SELECT id, routno, descr, is_taken FROM route"
	AllItemsQuery           = "SELECT id, name, item_unit, sale_price, purchased_price, descr, itemcategory, img FROM saleitem"
	BuyersQuery             = "SELECT id, name, address, latitude, longitude, route_id, contact_no, email, priority FROM buyer"
	DemandQuery             = "SELECT id, dnum, ddate, sitem, itemcategory, unit, rate, qty, amount, status, issueqty, vehicle FROM demand"
	AddSaleQuery            = "INSERT INTO sale(dnum, ddate, route, vehicle, driver, buyer, sitem, qty, qtyout, qtybal, idate, status, credit, sale_price) VALUES(?, DATE(?), ?, ?, ?, ?, ?, ?, ?, ?, DATE(?), ?, ?, ?)"
	DeleteSaleQuery         = "DELETE FROM sale WHERE id=?;"
	AllCategoriesQuery      = "SELECT id, cate FROM cbcategory;"
	PendingPaymentsQuery    = "SELECT id, sale_id, added_date, paid_date, balance, is_paid, buyer_id, salegroup_id FROM payments WHERE buyer_id=? AND is_paid=?;"
	AddPendingPaymentsQuery = "INSERT INTO payments(sale_id, added_date, paid_date, balance, is_paid, buyer_id, salegroup_id) VALUES(?, ?, ?, ?, ?, ?, ?);"
	SaleGroupsQuery         = "INSERT INTO salegroups(foo) VALUES(0);"
	FetchBuyerPricesQuery   = "SELECT id, item_id, buyer_id, price, item_category FROM buyer_item_prices WHERE buyer_id = ? ORDER BY id DESC;"
	AddBuyerPriceQuery      = "INSERT INTO buyer_item_prices(item_id, buyer_id, price, item_category) VALUES(?, ?, ?, ?);"
	UpdateBuyerPriceQuery   = "UPDATE buyer_item_prices SET price=? WHERE id=?;"
)

const (
	DemandLoaded          = "Loaded"
	LedgerStatusSold      = "Sold"
	LedgerStatusPurchased = "Purchase"
)
