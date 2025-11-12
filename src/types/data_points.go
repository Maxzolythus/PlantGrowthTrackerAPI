package types

// DataPoint is a representation of the data captured by our peripherals and stored in the DB
type DataPoint struct {
	Id           int      `json:"_id,omitempty" validate:"omitempty,mongodb"`
	Timestamp    string   `json:"timestamp" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Height       *float64 `json:"height,omitempty" validate:"omitempty,gte=0"`
	SoilMoisture *int     `json:"soilMoisture,omitempty" validate:"omitempty,gte=0"`
	PH           *float64 `json:"pH,omitempty" validate:"omitempty,gte=0,lte=14"`
	Fertilizer   *float64 `json:"fertilizer,omitempty" validate:"omitempty,gte=0"`
	Tempurature  *float64 `json:"tempurature,omitempty" validate:"omitempty"`
}
