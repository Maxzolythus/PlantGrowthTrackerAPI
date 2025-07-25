package types

import "time"

// DataPoint is a representation of the data captured by our peripherals and stored in the DB
type DataPoint struct {
	Id           int       `validate:"required"`
	Timestamp    time.Time `validate:"required,datetime"`
	SoilMoisture int
	PH           float64
	Fertilizer   float64
	Tempurature  float64
}
