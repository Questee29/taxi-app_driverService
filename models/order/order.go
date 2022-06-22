package order

type UserRequest struct {
	ID       int32  `json:"id"`
	TaxiType string `json:"taxi_type"`
}
type DriverResponse struct {
	Driverid int32  `json:"Driver_id"`
	TaxType  string `json:"taxi_type"`
}
