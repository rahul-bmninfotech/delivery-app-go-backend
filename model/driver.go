package model

type Driver struct {
	Id            int    `db:"id" json:"id"`
	ContactNumber string `db:"contact_no" json:"contact_no"`
	Email         string `db:"email" json:"email"`
	Username      string `db:"username" json:"username"`
	Password      string `db:"password" json:"password"`
	Address       string `db:"address" json:"address"`
	Name          string `db:"name" json:"name"`
	DateOfJoining string `db:"doj" json:"doj"`
	DateOfBirth   string `db:"dob" json:"dob"`
	LicenseNumber string `db:"licenseno" json:"licenseno"`
}
