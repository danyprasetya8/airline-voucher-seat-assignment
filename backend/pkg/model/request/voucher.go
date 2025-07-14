package request

type CheckVoucher struct {
	FlightNumber string `json:"flightNumber"`
	Date         string `json:"date"`
}

type GenerateVoucher struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	FlightNumber string `json:"flightNumber"`
	Date         string `json:"date"`
	Aircraft     string `json:"aircraft"`
}
