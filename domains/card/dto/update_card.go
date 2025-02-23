package dto

type UpdateCardRequest struct {
	ID              string  `json:"id" validate:"required"`
	Rfid            string  `json:"rfid" validate:"required"`
	OwnerName       string  `json:"owner_name" validate:"required"`
	CardType        string  `json:"card_type" validate:"required"`
	VehicleType     string  `json:"vehicle_type" validate:"required"`
	LicensePlate    string  `json:"license_plate" validate:"required"`
	ExpiredDate     string  `json:"expired_date" validate:"required"`
	LastIOHistoryID *string `json:"last_io_history_id"`
}
