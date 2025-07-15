package entity

import (
	"time"
)

type Audit struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Voucher struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	CrewID       string `gorm:"type:text;not null"`
	CrewName     string `gorm:"type:text;not null"`
	FlightNumber string `gorm:"type:text;not null;index:idx_flight_number_date,unique"`
	FlightDate   string `gorm:"type:text;not null;index:idx_flight_number_date,unique"`
	AircraftType string `gorm:"type:text;not null"`
	Seat         string `gorm:"type:text;not null"`
	Audit
}
