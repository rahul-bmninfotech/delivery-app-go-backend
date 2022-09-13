package router

import (
	"DeliveryApp/handlers"
	"net/http"
)

//Route identifies a single endpoint
type Route struct {
	Name        string
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

//Routes is a typealias for an array of HTTPRoute objects
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Driver Login",
		Path:        "/DriverLogin",
		Method:      "POST",
		HandlerFunc: handlers.DriverLoginHandler,
	},

	Route{
		Name:        "DriverFetchVehicles",
		Path:        "/DriverFetchVehicles",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchVehiclesHandler,
	},

	Route{
		Name:        "DriverMarkVehicleTaken",
		Path:        "/DriverMarkVehicleTaken",
		Method:      "GET",
		HandlerFunc: handlers.DriverMarkVehicleTakenHandler,
	},

	Route{
		Name:        "DriverMarkVehicleFree",
		Path:        "/DriverMarkVehicleFree",
		Method:      "GET",
		HandlerFunc: handlers.DriverMarkVehicleFreeHandler,
	},

	Route{
		Name:        "DriverMarkVehicleAndRouteFree",
		Path:        "/DriverMarkVehicleAndRouteFree",
		Method:      "GET",
		HandlerFunc: handlers.DriverMarkRouteAndVehicleFreeHandler,
	},

	Route{
		Name:        "DriverPostVehicleComment",
		Path:        "/DriverPostVehicleComment",
		Method:      "GET",
		HandlerFunc: handlers.DriverVehicleCommentHandler,
	},

	Route{
		Name:        "DriverFetchRoutes",
		Path:        "/DriverFetchRoutes",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchRoutesHandler,
	},

	Route{
		Name:        "DriverMarkRouteTaken",
		Path:        "/DriverMarkRouteTaken",
		Method:      "GET",
		HandlerFunc: handlers.DriverMarkRouteTakenHandler,
	},

	Route{
		Name:        "DriverFetchCategories",
		Path:        "/FetchCategories",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchCategoriesHandler,
	},

	Route{
		Name:        "DriverMarkRouteFree",
		Path:        "/DriverMarkRouteFree",
		Method:      "GET",
		HandlerFunc: handlers.DriverMarkRouteFreeHandler,
	},

	Route{
		Name:        "DriverFetchAllItems",
		Path:        "/DriverFetchAllItems",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchAllItemsHandler,
	},

	Route{
		Name:        "DriverFetchDemand",
		Path:        "/DriverFetchDemand",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchDemandHandler,
	},

	Route{
		Name:        "DriverFetchBuyers",
		Path:        "/DriverFetchBuyers",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchBuyersHandler,
	},
	Route{
		Name:        "DriverAddSales",
		Path:        "/DriverAddSales",
		Method:      "POST",
		HandlerFunc: handlers.DriverAddSalesHandler,
	},
	Route{
		Name:        "DriverFetchAllSales",
		Path:        "/DriverFetchAllSales",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchAllSales,
	},
	Route{
		Name:        "Ping",
		Path:        "/Ping",
		Method:      "GET",
		HandlerFunc: handlers.PingHandler,
	},

	Route{
		Name:        "Payments",
		Path:        "/Payments",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchPendingPaymentsHandler,
	},

	Route{
		Name:        "AddPayments",
		Path:        "/AddPayments",
		Method:      "POST",
		HandlerFunc: handlers.DriverAddPendingPaymentsHandler,
	},

	Route{
		Name:        "UpdatePayments",
		Path:        "/UpdatePayments",
		Method:      "POST",
		HandlerFunc: handlers.DriverUpdatePendingPaymentsHandler,
	},

	Route{
		Name:        "FetchBuyerPrices",
		Path:        "/BuyerPrices",
		Method:      "GET",
		HandlerFunc: handlers.DriverFetchBuyerItems,
	},

	Route{
		Name:        "AddBuyerPrice",
		Path:        "/AddBuyerPrice",
		Method:      "GET",
		HandlerFunc: handlers.DriverAddBuyerItem,
	},

	Route{
		Name:        "UpdateBuyerPrice",
		Path:        "/UpdateBuyerPrice",
		Method:      "POST",
		HandlerFunc: handlers.DriverUpdateBuyerItem,
	},

	Route{
		Name:        "Fetch SaleItemPriority",
		Path:        "/sale_item_priority",
		Method:      "GET",
		HandlerFunc: handlers.GetSaleItemPrioritiesHandler,
	},

	Route{
		Name:        "Add SaleItemPriority",
		Path:        "/add_sale_item_priority",
		Method:      "POST",
		HandlerFunc: handlers.AddSaleItemPrioritiesHandler,
	},

	Route{
		Name:        "Fetch Invoice for SaleGroup ids",
		Path:        "/salegroup_invoices",
		Method:      "GET",
		HandlerFunc: handlers.GetSaleGroupInvoicesHandler,
	},

	Route{
		Name:        "ServeMedia",
		Path:        "/Media",
		Method:      "GET",
		HandlerFunc: handlers.DriverServeMedia,
	},

	Route{
		Name:        "Driver Add Invoice",
		Path:        "/AddInvoice",
		Method:      "POST",
		HandlerFunc: handlers.DriverAddInvoiceHandler,
	},
	Route{
		Name:        "Save Buyer Priority",
		Path:        "/buyer_priority",
		Method:      "POST",
		HandlerFunc: handlers.SaveBuyerPriorityHandler,
	},
	Route{
		Name:        "Get Buyer Priority",
		Path:        "/buyer_priority",
		Method:      "GET",
		HandlerFunc: handlers.GetBuyerPriorityHandler,
	},
}
