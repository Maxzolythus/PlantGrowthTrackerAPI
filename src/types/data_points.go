package types

import "time"

// DataPoint is a representation of the data captured by our peripherals and stored in the DB
type DataPoint struct {
	Id           int       `json:"_id" validate:"required,mongo"`
	Timestamp    time.Time `json:"timestamp" validate:"required,datetime"`
	SoilMoisture int       `json:"soilMoisture"`
	PH           float64   `json:"ph"`
	Fertilizer   float64   `json:"fertilizer"`
	Tempurature  float64   `json:"tempurature"`
}
