package model

type Vehicle struct {
	Id            int    `db:"id" json:"id"`
	VehicleNumber string `db:"vehicle_no" json:"vehicle_no"`
	Comments      string `db:"comments" json:"comments"`
	IsTaken       int    `db:"is_taken" json:"is_taken"`
}

//type DriverAssignedVehicle struct {
//	Id        int `db:"id" json:"id"`
//	VehicleId int `db:"vehicle_id" json:"vehicle_id"`
//	DriverId  int `db:"driver_id" json:"driver_id"`
//}
