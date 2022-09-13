package model

type Route struct {
	Id               int    `db:"id" json:"id"`
	RouteName        string `db:"routno" json:"routno"`
	RouteDescription string `db:"descr" json:"descr"`
	IsTaken          int    `db:"is_taken" json:"is_taken"`
}

//type DriverAssignedRoute struct {
//	Id       int `db:"id" json:"id"`
//	RouteId  int `db:"route_id" json:"route_id"`
//	DriverId int `db:"driver_id" json:"driver_id"`
//}
