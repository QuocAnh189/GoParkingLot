package dto

type CreateCardRequest struct {
	Rfid         string `json:"rfid"`
	OwnerName    string `json:"owner_name"`
	CardType     string `json:"card_type"`
	VehicleType  string `json:"vehicle_type"`
	LicensePlate string `json:"license_plate"`
	ExpiredDate  string `json:"expired_date"`
}
