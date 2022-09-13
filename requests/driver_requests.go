package requests

import (
	db "DeliveryApp/database"
	"DeliveryApp/model"
	"DeliveryApp/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// DriverLoginVerify verifies login details for a Driver
func DriverLoginVerify(username, password string) (driver model.Driver, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}
	fmt.Println(username, password)
	queryCondition := "WHERE username=? AND password=?;"
	query := fmt.Sprintf("%s %s", utils.DriverLoginQuery, queryCondition)
	err = dtb.QueryRowx(query, username, password).StructScan(&driver)
	return
}

// DriverPostComment posts a Driver's comments about a vehicle
func DriverPostComment(vehicleID string, comment string) (ok bool, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	commentQuery := "UPDATE vehicle SET comments=CONCAT(?, comments) WHERE id=?"
	_, err = dtb.Exec(commentQuery, comment, vehicleID)
	if err != nil {
		return
	}
	ok = true
	return
}

// FetchDriverBuyers returns slice of Buyers assigned to that Route
func FetchDriverBuyers(routeID string) (buyers []model.Buyer, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}
	queryCondition := "WHERE route_id=?;"
	buyersQuery := fmt.Sprintf("%s %s", utils.BuyersQuery, queryCondition)
	rows, err := dtb.Queryx(buyersQuery, routeID)
	if err != nil {
		return
	}
	defer rows.Close()

	var buyer model.Buyer
	for rows.Next() {
		err = rows.StructScan(&buyer)
		if err != nil {
			continue
		}
		buyers = append(buyers, buyer)
	}
	return
}

// FetchAllItems returns slice of all SaleItems
func FetchAllItems() (items []model.SaleItem, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	query := utils.AllItemsQuery
	var item model.SaleItem
	rows, err := dtb.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&item)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, item)
	}
	return
}

// DriverFetchDemands fetches all Demands for the provided Vehicle
func DriverFetchDemands(vehicleID string) (demands []model.LoadedDemand, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	query := "SELECT d.id, d.dnum, d.ddate, d.unit, d.rate, d.qty, d.amount, d.status, d.issueqty, s.id, s.name, s.item_unit, s.sale_price, s.purchased_price, s.descr, s.itemcategory, s.img, c.id, c.cate, v.id, v.vehicle_no, v.comments, v.is_taken FROM demand d, saleitem s, cbcategory c, vehicle v WHERE d.sitem = s.id AND d.itemcategory = c.id AND d.vehicle = v.id AND d.vehicle=? AND d.status=?;"
	rows, err := dtb.Query(query, vehicleID, utils.DemandLoaded)
	if err != nil {
		return
	}

	defer rows.Close()
	var ld model.LoadedDemand
	for rows.Next() {
		err = rows.Scan(
			&ld.Id,
			&ld.DemandNumber,
			&ld.DemandDate,
			&ld.ItemUnit,
			&ld.Rate,
			&ld.Quantity,
			&ld.Amount,
			&ld.Status,
			&ld.IssueQuantity,
			&ld.SaleItem.Id,
			&ld.SaleItem.Name,
			&ld.SaleItem.ItemUnit,
			&ld.SaleItem.SalePrice,
			&ld.SaleItem.PurchasedPrice,
			&ld.SaleItem.Description,
			&ld.SaleItem.ItemCategory,
			&ld.SaleItem.ImagePath,
			&ld.CBCategory.ID,
			&ld.CBCategory.Category,
			&ld.Vehicle.Id,
			&ld.Vehicle.VehicleNumber,
			&ld.Vehicle.Comments,
			&ld.Vehicle.IsTaken,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		demands = append(demands, ld)
	}
	return
}

/************************************************************************
	NEW METHODS
 ************************************************************************/

// DriverFetchVehicles returns slice of all Vehicles that aren't taken
func DriverFetchVehicles() (vehicles []model.Vehicle, err error) {

	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	// queryCondition := "WHERE is_taken=0;"
	// query := fmt.Sprintf("%s %s", utils.VehicleQuery, queryCondition)
	query := utils.VehicleQuery

	rows, err := dtb.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var vehicle model.Vehicle
	for rows.Next() {
		err = rows.StructScan(&vehicle)
		if err != nil {
			fmt.Println(err)
			continue
		}
		vehicles = append(vehicles, vehicle)
	}
	return
}

// MarkVehicleTaken marks a Vehicle with the provided id taken
func MarkVehicleTaken(r *http.Request) (ok bool, err error) {
	// values := r.URL.Query()
	// vehicleID := values.Get(utils.QueryVehicleId)
	return ok, nil

	// dtb, err := db.MakeDB()
	// if err != nil {
	// 	return
	// }
	// query := "UPDATE vehicle SET is_taken=1 WHERE id=?;"
	// res, err := dtb.Exec(query, vehicleID)
	// if err != nil {
	// 	return
	// }
	// rowsAffected, err := res.RowsAffected()
	// if err != nil {
	// 	return
	// }
	// ok = rowsAffected == 1
	// return
}

// DriverFetchRoutes fetches all routes that are not taken
func DriverFetchRoutes() (routes []model.Route, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	// queryCondition := "WHERE is_taken=0;"
	// query := fmt.Sprintf("%s %s", utils.RouteQuery, queryCondition)
	query := utils.RouteQuery

	rows, err := dtb.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var route model.Route
	for rows.Next() {
		err = rows.StructScan(&route)
		if err != nil {
			fmt.Println(err)
			continue
		}
		routes = append(routes, route)
	}
	return

}

// MarkRouteTaken marks a route as taken
func MarkRouteTaken(r *http.Request) (ok bool, err error) {
	values := r.URL.Query()
	routeID := values.Get(utils.QueryRouteId)

	dtb, err := db.MakeDB()
	if err != nil {
		return
	}
	query := "UPDATE route SET is_taken=1 WHERE id=?;"
	res, err := dtb.Exec(query, routeID)
	if err != nil {
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	ok = rowsAffected == 1
	return
}

// AddSale adds the provided Sales
// Data is sent in the body
func AddSale(r *http.Request) (sale model.Sale, err error) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &sale)
	if err != nil {
		return
	}
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	query := utils.AddSaleQuery
	fmt.Println(query)
	res, err := dtb.Exec(query,
		sale.DemandNumber,
		sale.DemandDate,
		sale.Route,
		sale.Vehicle,
		sale.Driver,
		sale.Buyer,
		sale.SaleItem,
		sale.Quantity,
		sale.QuantityOut,
		sale.QuantityBalance,
		sale.IssuedDate,
		sale.Status,
	)

	if err != nil {
		return
	}
	rowID, err := res.LastInsertId()
	if err != nil {
		return
	}
	sale.Id = int(rowID)
	return
}

//func AddReturned(){
//
//}

// DeleteSale delets a sale with the given ID
func DeleteSale(r *http.Request) (ok bool, err error) {
	values := r.URL.Query()
	saleID := values.Get(utils.QuerySaleId)

	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	query := utils.DeleteSaleQuery
	_, err = dtb.Exec(query, saleID)
	if err != nil {
		return
	}
	ok = true
	return
}

// FetchCategories fetches all CBCategories
func FetchCategories() (cats []model.CBCategory, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}
	query := utils.AllCategoriesQuery
	rows, err := dtb.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var cat model.CBCategory

	for rows.Next() {
		err = rows.StructScan(&cat)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cats = append(cats, cat)
	}
	return
}

// MarkRouteFree marks a Route free
func MarkRouteFree(r *http.Request) (ok bool, err error) {
	return true, nil
	// 	values := r.URL.Query()
	// 	routeID := values.Get(utils.QueryRouteId)
	// 	dtb, err := db.MakeDB()
	// 	if err != nil {
	// 		return
	// 	}

	// 	query := "UPDATE route SET is_taken=0 WHERE id=?;"
	// 	_, err = dtb.Exec(query, routeID)
	// 	if err != nil {
	// 		return
	// 	}

	// 	ok = err == nil
	// 	return
}

// MarkVehicleFree marks a Vehicle free
func MarkVehicleFree(r *http.Request) (ok bool, err error) {
	return true, nil
	// values := r.URL.Query()
	// vehicleID := values.Get(utils.QueryVehicleId)
	// dtb, err := db.MakeDB()
	// if err != nil {
	// 	return
	// }
	// query := "UPDATE vehicle SET is_taken=0 WHERE id=?;"
	// _, err = dtb.Exec(query, vehicleID)
	// if err != nil {
	// 	return
	// }
	// ok = err == nil
	// return
}

// MarkRouteAndVehicleFree marks Route and Vehicle free.
// If either Route or Vehicle id is -1, it is not marked free.
func MarkRouteAndVehicleFree(r *http.Request) (ok bool, err error) {
	return true, nil
	// values := r.URL.Query()
	// routeID := values.Get(utils.QueryRouteId)
	// vehicleID := values.Get(utils.QueryVehicleId)

	// if routeID == "-1" {
	// 	return MarkVehicleFree(r)
	// } else if vehicleID == "-1" {
	// 	return MarkRouteFree(r)
	// } else {
	// 	routeOk, err := MarkRouteFree(r)
	// 	vehicleOk, err := MarkVehicleFree(r)
	// 	return routeOk && vehicleOk, err
	// }
}

// ReqAddSales adds the Sales
// Sales are sent in the body
func ReqAddSales(r *http.Request) (sales []model.Sale, err error) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &sales)
	if err != nil {
		return
	}

	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	tx, err := dtb.Begin()
	if err != nil {
		return
	}

	stmt, err := tx.Prepare(utils.AddSaleQuery)
	for idx, sale := range sales {
		res, err := stmt.Exec(
			sale.DemandNumber,
			sale.DemandDate,
			sale.Route,
			sale.Vehicle,
			sale.Driver,
			sale.Buyer,
			sale.SaleItem,
			sale.Quantity,
			sale.QuantityOut,
			sale.QuantityBalance,
			sale.IssuedDate,
			sale.Status,
			sale.Credit,
			sale.SalePrice,
		)

		if err != nil {
			continue
		}
		rowID, err := res.LastInsertId()
		if err != nil {
			continue
		}
		sales[idx].Id = int(rowID)
	}
	tx.Commit()

	if sales[0].Credit == "NO" {
		return
	}
	err = ReqAddPendingPayments(sales)
	return
}

// ReqSaveInvoice saves the invoice to disk and adds the invoice title to the relevant Sales
func ReqSaveInvoice(r *http.Request) (ok bool, err error) {
	v := r.URL.Query()
	tempIDs := v.Get(utils.QueryInvoiceSaleIds)
	invoiceNumber := v.Get(utils.QueryInvoiceNumberParameter)
	var invoiceName string

	r.ParseMultipartForm(32 << 20)

	file, fileHeader, err := r.FormFile("invoice")
	if err != nil {
		return
	}
	defer file.Close()

	filePath, err := utils.MediaDir()
	if err != nil {
		fmt.Println("No Media Directory Found")
		return
	}

	invoiceFile, err := os.OpenFile(filepath.Join(filePath, fileHeader.Filename), os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer invoiceFile.Close()
	io.Copy(invoiceFile, file)
	invoiceName = fileHeader.Filename

	// Save the comment image, if sent
	if file, fileHeader, err = r.FormFile("comment"); err == nil {
		log.Printf("Received file: %+v\n", file)
		if commentsDir, err := utils.GetDirectory(utils.CommentsDir); err == nil {
			commentsFilePath := filepath.Join(commentsDir, fileHeader.Filename)
			if commentsFile, err := os.OpenFile(commentsFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm); err == nil {
				if _, err = io.Copy(commentsFile, file); err == nil {
					utils.SendMail(commentsFilePath, fileHeader.Filename, "ukinch2@gmail.com")
				}
				defer commentsFile.Close()
			}
		}
	}

	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	invoiceQuery := fmt.Sprintf("UPDATE sale SET invoice=?, invoice_no=? WHERE id in (%s)", tempIDs)
	fmt.Println(invoiceQuery)
	//res, err := dtb.Exec(invoiceQuery, fileHeader.Filename, invoiceNumber)
	res, err := dtb.Exec(invoiceQuery, invoiceName, invoiceNumber)
	if err != nil {
		return
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println("Rows Affected ", ra)
	ok = true
	return
}

// ReqFetchPendingPayments fetches all unpaid Payments for a buyer
// Extract buyerIDs from url as comma-separated list and get all payments for it
func ReqFetchPendingPayments(r *http.Request) (payments []model.Payments, err error) {
	v := r.URL.Query()
	buyerIDs := v.Get(utils.QueryPaymentsBuyersID)
	buyerIDsArray := strings.Split(buyerIDs, ",")

	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	inCondition := "("
	for _, v := range buyerIDsArray {
		inCondition += fmt.Sprintf("%s,", v)
	}
	inCondition = strings.TrimRight(inCondition, ",")
	inCondition += ")"
	fmt.Printf("Query condition is %s\n", inCondition)

	query := fmt.Sprintf("SELECT id, sale_id, added_date, paid_date, balance, is_paid, buyer_id, salegroup_id FROM payments WHERE is_paid=0 AND buyer_id IN %s;", inCondition)
	fmt.Println(query)
	rows, err := dtb.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var payment model.Payments
	for rows.Next() {
		err = rows.StructScan(&payment)
		if err != nil {
			fmt.Println(err)
			continue
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

// ReqAddPendingPayments adds payments that a buyer has marked as credit (to be paid later)
// Extract list of comma-separated sale id's, balance and buyer id from the url
func ReqAddPendingPayments(sales []model.Sale) (err error) {
	//r.ParseForm()
	//v := r.PostForm
	//saleIDs := v.Get(utils.QueryPaymentsSaleIDs)
	//balance := v.Get(utils.QueryPaymentsBalance)
	//buyerID := v.Get(utils.QueryPaymentsBuyerID)

	//saleIDArray := strings.Split(saleIDs, ",")
	//if len(saleIDArray) == 0 {
	//	return errors.New("no sale ids sent")
	//}
	var balance float64
	buyerID := sales[0].Buyer

	saleIDs := make([]int, len(sales))
	salePriceCondition := "("
	for i, x := range sales {
		saleIDs[i] = x.Id
		salePriceCondition += fmt.Sprintf("%d,", x.Id)
	}
	salePriceCondition = strings.TrimRight(salePriceCondition, ",")
	salePriceCondition += ")"

	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	type TempSaleItem struct {
		ItemID    int    `db:"id"`
		SalePrice string `db:"sale_price"`
		Quantity  string `db:"qty"`
	}
	salePriceQuery := "SELECT s.id, s.qty, si.sale_price FROM sale AS s INNER JOIN saleitem AS si ON s.sitem = si.id WHERE s.id IN " + salePriceCondition
	fmt.Printf("SaleItem Query is %s\n", salePriceQuery)
	//salePriceQuery := "SELECT id, sale_price FROM saleitem WHERE id IN " + salePriceCondition
	var tempItem TempSaleItem
	var saleItems []TempSaleItem
	itemRows, err := dtb.Queryx(salePriceQuery)
	if err != nil {
		return
	}
	defer itemRows.Close()

	for itemRows.Next() {
		err = itemRows.StructScan(&tempItem)
		if err != nil {
			fmt.Println(err)
			continue
		}
		saleItems = append(saleItems, tempItem)
		qty, err := strconv.ParseFloat(tempItem.Quantity, 32)
		if err != nil {
			continue
		}
		price, err := strconv.ParseFloat(tempItem.SalePrice, 32)
		if err != nil {
			continue
		}
		balance += qty * price
	}

	fmt.Printf("Balance is %f\n", balance)

	groupResult, err := dtb.Exec(utils.SaleGroupsQuery)
	if err != nil {
		return
	}
	groupID, err := groupResult.LastInsertId()
	if err != nil {
		return
	}

	tx, err := dtb.Begin() // Create a transaction (sql.Tx)
	if err != nil {
		return
	}
	stmt, err := tx.Prepare(utils.AddPendingPaymentsQuery) // Create a Prepared Statement
	if err != nil {
		return
	}
	now := time.Now()
	for _, s := range saleIDs {
		stmt.Exec(s, now, now, balance, 0, buyerID, groupID) // Execute the prepared statement
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	return
}

// ReqUpdatePendingPayments marks pending payments as paid
// Extract comma-separated payment group ids from url, which need to be marked as paid
func ReqUpdatePendingPayments(r *http.Request) (err error) {
	//r.ParseForm()
	//v := r.PostForm
	//paymentGroupIDs := v.Get(utils.QueryPaymentGroupIDs)
	fmt.Println(r.Form)
	paymentGroupIDs := r.PostFormValue(utils.QueryPaymentGroupIDs)
	fmt.Printf("GroupIds are %v\n", paymentGroupIDs)
	paymentGroupIDsArray := strings.Split(paymentGroupIDs, ",")

	//var queryCondition string
	queryCondition := "("
	for _, v := range paymentGroupIDsArray {
		queryCondition += fmt.Sprintf("%s,", v)
	}
	queryCondition = strings.TrimRight(queryCondition, ",")
	queryCondition += ")"
	//ind := strings.LastIndex(queryCondition, ",")
	//foo := queryCondition[:-1]

	//args := map[string]interface{} {
	//	"dt": time.Now(),
	//	"ids": strings.Split(paymentGroupIDs, ","),
	//}

	fmt.Printf("Query Condition is %s\n", queryCondition)
	query := "UPDATE payments SET paid_date=?, is_paid=1 WHERE salegroup_id IN " + queryCondition
	fmt.Println(query)
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}
	_, err = dtb.Exec(query, time.Now())
	return
	//query := "UPDATE payments SET paid_date=?, is_paid=1 WHERE salegroup_id IN (?" + strings.Repeat(",?", len(paymentGroupIDsArray)-1) + ");"
	//query := "UPDATE payments SET paid_date=:dt, is_paid=1 WHERE salegroup_id IN (:ids)"
	//return
}

// ReqFetchBuyerPrices fetches all BuyerItemPrice instances for the provided buyer id
// Returns array of BuyerItemPrice instances
func ReqFetchBuyerPrices(buyerID string) (prices []model.BuyerItemPrice, err error) {
	prices = make([]model.BuyerItemPrice, 0)
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	rows, err := dtb.Queryx(utils.FetchBuyerPricesQuery, buyerID)
	if err != nil {
		return
	}
	defer rows.Close()

	var buyerItem model.BuyerItemPrice
	for rows.Next() {
		err = rows.StructScan(&buyerItem)
		if err != nil {
			fmt.Println(err)
			continue
		}
		prices = append(prices, buyerItem)
	}
	err = nil

	encountered := map[int]bool{}
	var results []model.BuyerItemPrice

	for _, v := range prices {
		if !encountered[v.ItemID] {
			results = append(results, v)
			encountered[v.ItemID] = true
		}
	}

	prices = results
	return
}

//func removeDuplicates(items []model.BuyerItemPrice) []model.BuyerItemPrice {
//	fmt.Printf("Received values: %d\n", len(items))
//	encountered := map[int]bool{}
//	var results []model.BuyerItemPrice
//
//	for _, v := range items {
//		if !encountered[v.ItemID] {
//			results = append(results, v)
//			encountered[v.ItemID] = true
//		}
//	}
//	fmt.Printf("Returning values: %d\n", len(results))
//	return results
//}

// AddBuyerPrice adds a BuyerItemPrice for a specific buyer
// It returns the BuyerItemPrice instance
func ReqAddBuyerPrice(itemID string, buyerID string, price string, itemCategory string) (buyerItem model.BuyerItemPrice, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	res, err := dtb.Exec(utils.AddBuyerPriceQuery, itemID, buyerID, price, itemCategory)
	if err != nil {
		return
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		return
	}
	err = dtb.QueryRowx("SELECT id, item_id, buyer_id, price, item_category FROM buyer_item_prices WHERE id = ?;", lastInsertedID).StructScan(&buyerItem)
	return
}

// UpdateBuyerPrice updates the BuyerItemPrice instance with the new price
func ReqUpdateBuyerPrice(id string, newPrice string) (err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}
	_, err = dtb.Exec(utils.UpdateBuyerPriceQuery, newPrice, id)
	return
}

func ReqSaveBuyerPriorities(id, driverId, routeId, priority string) (priorities model.BuyerPriorities, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	query := "insert into buyer_priority(id, driver_id, route_id, priorities) values(?, ?, ?, ?) on duplicate key update priorities=?;"
	res, err := dtb.Exec(query, id, driverId, routeId, priority, priority)
	if err != nil {
		return
	}
	rowId, err := res.LastInsertId()
	if err != nil {
		return
	}

	fetchQuery := "select id, driver_id, route_id, priorities from buyer_priority where id=?;"
	err = dtb.QueryRowx(fetchQuery, rowId).StructScan(&priorities)
	return
}

//// ReqSaveBuyerPriority inserts or updates BuyerPriority instances
//func ReqSaveBuyerPriority(body io.ReadCloser) (pris []model.BuyerPriority, err error) {
//	b1 := time.Now()
//	defer func() {
//		fmt.Printf("Saving Buyer's Priorities time: %+v\n", time.Since(b1))
//	}()
//	var priorities []model.BuyerPriority
//	err = json.NewDecoder(body).Decode(&priorities)
//	if err != nil {
//		return
//	}
//
//	dtb, err := db.MakeDB()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	//upsertQuery := "INSERT INTO buyer_priority(id, driver_id, buyer_id, priority) VAlUES(?, ?, ?, ?) ON DUPLICATE KEY UPDATE priority=?;"
//	upsertQuery := "insert into buyer_priority(id, driver_id, buyer_id, priority) values(?, ?, ?, ?) ON DUPLICATE KEY update priority=?;"
//	//upsertQuery := "REPLACE INTO buyer_priority(id, driver_id, buyer_id, priority) VALUES(?, ?, ?, ?);"
//	prTx, err := dtb.Begin()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	prStmt, err := prTx.Prepare(upsertQuery)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	exTime := time.Now()
//	for _, p := range priorities {
//		_, err = prStmt.Exec(p.Id, p.DriverID, p.BuyerID, p.Priority, p.Priority)
//		//_, err = prStmt.Exec(p.Id, p.DriverID, p.BuyerID, p.Priority)
//		if err != nil {
//			fmt.Println(err)
//			prTx.Rollback()
//			return
//		}
//	}
//
//	err = prTx.Commit()
//	prStmt.Close()
//	fmt.Printf("After execution: %+v\n", time.Since(exTime))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	driverID := ""
//	if len(priorities) > 0 {
//		driverID = fmt.Sprintf("%d", priorities[0].DriverID)
//	}
//	if driverID == "" {
//		return
//	}
//	return ReqFetchBuyerPriority(driverID)
//}
//
//// ReqFetchBuyerPriority fetches BuyerPriority instances for the provided driver id
//func ReqFetchBuyerPriority(driverId string) (priorities []model.BuyerPriority, err error) {
//	b2 := time.Now()
//	defer func() {
//		fmt.Printf("Fetching Buyer's Priorities time: %+v\n", time.Since(b2))
//	}()
//	dtb, err := db.MakeDB()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	priorityQuery := "SELECT id, driver_id, buyer_id, priority FROM buyer_priority WHERE driver_id=?;"
//	rows, err := dtb.Queryx(priorityQuery, driverId)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer rows.Close()
//
//	var bp model.BuyerPriority
//	for rows.Next() {
//		err = rows.StructScan(&bp)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		priorities = append(priorities, bp)
//	}
//	err = nil
//	return
//}

func ReqFetchBuyerPriorities(driverId, routeId string) (priorities model.BuyerPriorities, err error) {
	dtb, err := db.MakeDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	priorityQuery := "SELECT id, driver_id, route_id, priorities FROM buyer_priority WHERE driver_id=? AND route_id=?;"
	err = dtb.QueryRowx(priorityQuery, driverId, routeId).StructScan(&priorities)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func ReqFetchSaleItemPriorities() (priorities []model.SaleItemPriority, err error) {
	fmt.Println("Fetching Item Priorities")
	dtb, err := db.MakeDB()
	if err != nil {
		fmt.Print(err)
		return
	}

	itemPriorityQuery := "SELECT id, item_id, priority FROM saleitem_priority;"
	rows, err := dtb.Queryx(itemPriorityQuery)
	if err != nil {
		return
	}
	defer rows.Close()

	var priority model.SaleItemPriority
	for rows.Next() {
		err = rows.StructScan(&priority)
		if err != nil {
			fmt.Print(err)
			continue
		}
		priorities = append(priorities, priority)
	}
	return
}

func ReqAddSaleItemPriorities(r *http.Request) (priorities []model.SaleItemPriority, err error) {
	//first := time.Now()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(buffer.Bytes(), &priorities)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("Unmarshalling time: %+v\n", time.Since(first))
	//fmt.Printf("Adding sale item priorities: %+v\n", priorities)

	dtb, err := db.MakeDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	addItemQuery := "INSERT INTO saleitem_priority(id, item_id, priority) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE priority=VALUES(priority);"

	//fmt.Println("Tx Begins")
	tx, err := dtb.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt, err := tx.Prepare(addItemQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Exec Begins")
	//second := time.Now()

	var pty *model.SaleItemPriority
	for i := 0; i < len(priorities); i++ {
		pty = &priorities[i]
		result, err := stmt.Exec(pty.Id, pty.ItemId, pty.Priority)
		if err != nil {
			fmt.Println()
			tx.Rollback()
			return priorities, err
		}
		if pty.Id == -1 {
			newId, err := result.LastInsertId()
			if err != nil {
				continue
			}
			pty.Id = int(newId)
		}
	}
	//fmt.Printf("Exec ends, time: %+v\n", time.Since(second))
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println("Tx Ends")
	//for _, x := range priorities {
	//	fmt.Println(x)
	//}
	return
}

func ReqGetSaleGroupInvoices(r *http.Request) (invoices []model.SaleGroupInvoice, err error) {
	invoices = make([]model.SaleGroupInvoice, 0)
	ids := r.URL.Query().Get("ids")
	splitIds := strings.Split(ids, ",")
	mainQuery := "SELECT p.salegroup_id, s.invoice FROM sale AS s INNER JOIN payments AS p ON s.id=p.sale_id WHERE s.id IN (SELECT sale_id FROM payments WHERE salegroup_id IN (?)) AND invoice NOT LIKE '' GROUP BY(invoice) ORDER BY salegroup_id ASC"
	
	numIds := make([]int, 0)
	for _, id := range splitIds {
		numId, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		numIds = append(numIds, numId)
	}

	dtb, err := db.MakeDB()
	if err != nil {
		return
	}

	//query, args, err := sqlx.In("SELECT sale_id FROM payments WHERE salegroup_id IN (?),", numIds)
	query, args, err := sqlx.In(mainQuery, numIds)
	if err != nil {
		return
	}
	query = dtb.Rebind(query)

	fmt.Println(query, args, err)

	rows, err := dtb.Queryx(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	var invoice model.SaleGroupInvoice
	for rows.Next() {
		err = rows.StructScan(&invoice)
		if err != nil {
			log.Println(err)
			continue
		}
		invoices = append(invoices, invoice)
	}
	return
}
