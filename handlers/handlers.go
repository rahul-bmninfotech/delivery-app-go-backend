package handlers

import (
	"DeliveryApp/requests"
	"DeliveryApp/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// DriverLoginHandler handles the login process
// Method: POST
// Params: username and password
// Returns: Driver instance if details valid, else error
func DriverLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue(utils.QueryUsername)
	password := r.PostFormValue(utils.QueryPassword)

	driver, err := requests.DriverLoginVerify(username, password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(driver)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// DriverVehicleCommentHandler handles the driver's comments about a vehicle
// Method: GET
// Params: vehicleID
// Returns: 200 if saved, else 404
func DriverVehicleCommentHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	vehicleID := values.Get(utils.QueryVehicleId)
	comment := values.Get(utils.QueryVehicleComment)
	ok, err := requests.DriverPostComment(vehicleID, comment)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !ok {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

// DriverFetchBuyersHandler fetches all buyers for the provided routeId
// Method: GET
// Params: routeID
// Return: array of Buyer instances, if found
func DriverFetchBuyersHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	routeID := values.Get(utils.QueryRouteId)

	buyers, err := requests.FetchDriverBuyers(routeID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(buyers)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// DriverFetchAllItemsHandler fetches all items in our DB
// Method: GET
// Params: none
// Return: array of SaleItem instances
func DriverFetchAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, err := requests.FetchAllItems()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// DriverFetchDemandHandler fetches all demands for the vehicleID provided
// Method: GET
// Params: vehicleID
// Return: array of Demand instances
func DriverFetchDemandHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	vehicleID := values.Get(utils.QueryVehicleId)

	demands, err := requests.DriverFetchDemands(vehicleID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(demands)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

/************************************************************************
	NEW METHODS
************************************************************************/

// DriverFetchVehiclesHandler fetches all vehicles
// Method: GET
// Params: none
// Return: array of Vehicle instances
func DriverFetchVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	vehicles, err := requests.DriverFetchVehicles()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(vehicles)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

// DriverMarkVehicleTakenHandler marks the vehicle with the provided vehicleID as taken
// Method: GET
// Params: vehicleID
// Return: 200 if successful, else 404
func DriverMarkVehicleTakenHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := requests.MarkVehicleTaken(r)
	if !ok || err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

// DriverFetchRoutesHandler fetches all routes
// Method: GET
// Params: none
// Return: array of Route instances
func DriverFetchRoutesHandler(w http.ResponseWriter, r *http.Request) {
	routes, err := requests.DriverFetchRoutes()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(routes)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

// DriverMarkRouteTakenHandler marks a route as taken
func DriverMarkRouteTakenHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := requests.MarkRouteTaken(r)
	if !ok || err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

// DriverAddSaleHandler adds a sale for a driver
// ******************  Ripe for Deletion ***********
func DriverAddSaleHandler(w http.ResponseWriter, r *http.Request) {
	sale, err := requests.AddSale(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(sale)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// DriverDeleteSaleHandler deletes a sale
func DriverDeleteSaleHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := requests.DeleteSale(r)
	if !ok || err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// DriverMarkVehicleFreeHandler marks a vehicle free
func DriverMarkVehicleFreeHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := requests.MarkVehicleFree(r)
	if !ok || err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

// DriverMarkRouteFreeHandler marks a route free
func DriverMarkRouteFreeHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := requests.MarkRouteFree(r)
	if !ok || err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

}

// DriverMarkRouteAndVehicleFreeHandler marks both route and vehicle free
func DriverMarkRouteAndVehicleFreeHandler(w http.ResponseWriter, r *http.Request) {
	ok, err := requests.MarkRouteAndVehicleFree(r)
	if !ok || err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

func DriverFetchAllSales(w http.ResponseWriter, r *http.Request) {}

// DriverFetchCategoriesHandler fetches all categories
func DriverFetchCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := requests.FetchCategories()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// DriverAddSalesHandler adds sales for a driver
func DriverAddSalesHandler(w http.ResponseWriter, r *http.Request) {
	sales, err := requests.ReqAddSales(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(sales)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

}

// DriverServeMedia serves the file with the `filename` passed in
func DriverServeMedia(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	fPath, err := utils.MediaPath(filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", filename))
	//w.Header().Set("filename", filename)
	http.ServeFile(w, r, fPath)
}

// DriverAddInvoiceHandler saves the invoice to disk, and adds the invoice name to the relevant Sales
// Method: POST
// Params: Comma separated list of Sale ids, and PDF invoice in the request body
// Returns: 200 if successful, else 400 (Bad Request)
func DriverAddInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	successful, err := requests.ReqSaveInvoice(r)
	if err != nil || !successful {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// DriverFetchPendingPaymentsHandler fetches all payments for a buyer that are pending
// Method: GET
// Params: none
// Returns: 200 if successful, else 404 (Not Found)
func DriverFetchPendingPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	payments, err := requests.ReqFetchPendingPayments(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(payments)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// DriverAddPendingPaymentsHandler adds pending payments for a buyer
// Method: POST
// Params: none
// Returns: 200 if successful, else 400
func DriverAddPendingPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	//err := requests.ReqAddPendingPayments(r)
	//if err != nil {
	//	fmt.Println(err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
}

// DriverUpdatePendingPaymentsHandler marks the pending payments as paid for a buyer
// Method: POST
// Params: none
// Returns: 200 if successful, else 400
func DriverUpdatePendingPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	err := requests.ReqUpdatePendingPayments(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// DriverFetchBuyerItems fetches all BuyerItemPrice instances for a buyer
func DriverFetchBuyerItems(w http.ResponseWriter, r *http.Request) {
	buyerID := r.URL.Query().Get(utils.QueryBuyerPriceIDs)
	fmt.Printf("Buyer id is %s\n", buyerID)
	buyerItems, err := requests.ReqFetchBuyerPrices(buyerID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(buyerItems)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// DriverAddBuyerItem adds a BuyerItemPrice instance
func DriverAddBuyerItem(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	itemID := v.Get(utils.QueryBuyerPriceItemID)
	buyerID := v.Get(utils.QueryBuyerPriceBuyerID)
	price := v.Get(utils.QueryBuyerPriceValue)
	category := v.Get(utils.QueryBuyerPriceCategory)

	buyerItem, err := requests.ReqAddBuyerPrice(itemID, buyerID, price, category)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(buyerItem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// DriverUpdateBuyerItem updates an existing BuyerItemPrice instance with a new price value
func DriverUpdateBuyerItem(w http.ResponseWriter, r *http.Request) {
	//v := r.URL.Query()
	itemID := r.PostFormValue(utils.QueryBuyerPriceID)
	newPrice := r.PostFormValue(utils.QueryBuyerPriceNewValue)

	err := requests.ReqUpdateBuyerPrice(itemID, newPrice)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

//SaveBuyerPriorityHandler saves the buyer priorities for a certain Driver
func SaveBuyerPriorityHandler(w http.ResponseWriter, r *http.Request) {
	//priorities, err := requests.ReqSaveBuyerPriority(r.Body)
	r.ParseForm()
	q := r.PostForm
	id := q.Get("id")
	dId := q.Get("driver_id")
	rId := q.Get("route_id")
	pr := q.Get("priorities")

	fmt.Println(id, dId, rId, pr)

	//priorities, err := requests.ReqSaveBuyerPriorities(r.Body)
	priorities, err := requests.ReqSaveBuyerPriorities(id, dId, rId, pr)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(priorities); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

//GetBuyerPriorityHandler fetches the buyer priorities for a certain Driver
func GetBuyerPriorityHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	driverID := q.Get(utils.QueryDriverId)
	routeId := q.Get(utils.QueryRouteId)
	priorities, err := requests.ReqFetchBuyerPriorities(driverID, routeId)
	//priorities, err := requests.ReqFetchBuyerPriority(driverID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err = json.NewEncoder(w).Encode(priorities); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func GetSaleItemPrioritiesHandler(w http.ResponseWriter, r *http.Request) {
	itemsPriorities, err := requests.ReqFetchSaleItemPriorities()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err = json.NewEncoder(w).Encode(itemsPriorities); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func AddSaleItemPrioritiesHandler(w http.ResponseWriter, r *http.Request) {
	addedPriorities, err := requests.ReqAddSaleItemPriorities(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err = json.NewEncoder(w).Encode(addedPriorities); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func GetSaleGroupInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	invoices, err := requests.ReqGetSaleGroupInvoices(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err = json.NewEncoder(w).Encode(invoices); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

}

// PingHandler returns 200 if service is up
func PingHandler(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Ping")
}
