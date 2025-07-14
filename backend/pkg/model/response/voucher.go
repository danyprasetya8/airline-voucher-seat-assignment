package response

type GetVoucher struct {
	ID           int      `json:"id"`
	CrewID       string   `json:"crewId"`
	CrewName     string   `json:"crewName"`
	FlightNumber string   `json:"flightNumber"`
	FlightDate   string   `json:"flightDate"`
	AircraftType string   `json:"aircraftType"`
	Seat         []string `json:"seat"`
	CreatedAt    string   `json:"createdAt"`
}

type CheckVoucher struct {
	Exists bool `json:"exists"`
}

type GenerateVoucher struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}
