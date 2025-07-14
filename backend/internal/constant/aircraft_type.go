package constant

import "airline-voucher-seat-assignment/pkg/model"

var AircraftTypes = []model.AircraftType{
	{
		Name:      "ATR",
		RowStart:  1,
		RowEnd:    18,
		TotalSeat: 4,
		Seats:     []string{"A", "B", "C", "D"},
	},
	{
		Name:      "Airbus 320",
		RowStart:  1,
		RowEnd:    32,
		TotalSeat: 6,
		Seats:     []string{"A", "B", "C", "D", "E", "F"},
	},
	{
		Name:      "Boeing 737 Max",
		RowStart:  1,
		RowEnd:    32,
		TotalSeat: 6,
		Seats:     []string{"A", "B", "C", "D", "E", "F"},
	},
}
