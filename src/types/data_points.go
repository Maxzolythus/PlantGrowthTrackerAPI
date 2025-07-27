package types

// DataPoint is a representation of the data captured by our peripherals and stored in the DB
type DataPoint struct {
	Id           int     `json:"_id,omitempty" validate:"omitempty,mongodb"`
	Timestamp    string  `json:"timestamp" validate:"required,datetime"`
	Height       float64 `json:"height"`
	SoilMoisture int     `json:"soilMoisture"`
	PH           float64 `json:"ph" validate:"gte=0,lte=14"`
	Fertilizer   float64 `json:"fertilizer"`
	Tempurature  float64 `json:"tempurature"`
}
