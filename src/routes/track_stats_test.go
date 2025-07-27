package routes

import (
	mockUtils "main/mocks/utils"
	"main/src/types"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTrackStats(t *testing.T) {
	// Test Happy Path 1 - All things filled out
	t.Run("TestTrackStatsAll", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp:    time.Now().Format(time.DateTime),
			Height:       12.7,
			SoilMoisture: 57,
			PH:           4.7,
			Fertilizer:   125,
			Tempurature:  87,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(data).Return(nil)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	// Test Happy Path 2 - Only Height
	t.Run("TestTrackStatsOnlyHeight", func(t *testing.T) {

	})

	// Test Happy Path 3 - Only Soil Moisture
	t.Run("TestTrackStatsOnlySoilMoisture", func(t *testing.T) {

	})

	// Test Happy Path 4 - Only Fertilizer
	t.Run("TestTrackStatsOnlyFetilizer", func(t *testing.T) {

	})

	// Test Happy Path 5 - Only Temp
	t.Run("TestTrackStatsOnlyTemp", func(t *testing.T) {

	})

	// Test Happy Path - Only ph
	t.Run("TestTrackStatsOnlyPH", func(t *testing.T) {

	})

	// Test Sad Path 1 - Wrong Type
	t.Run("TestTrackStatsWrongType", func(t *testing.T) {

	})

	// Test Sad Path 2 - Only time stamp
	t.Run("TestTrackStatsOnlyTimeStamp", func(t *testing.T) {

	})

	// Test Sad Path 3 - No time stamp
	t.Run("TestTrackStatsNoTimeStamp", func(t *testing.T) {

	})

	// Test Sad Path 3 - Invalid time stamp
	t.Run("TestTrackStatsInvalidTimeStamp", func(t *testing.T) {

	})

	// Test Sad Path 2 - Invalid pH
	t.Run("TestTrackStatsInvalidPH", func(t *testing.T) {

	})

	// Test Sad Path 3 - All nil
	t.Run("TestTrackStatsAllNil", func(t *testing.T) {

	})

	// Test Sad Path 4 - Maliscious Text
	t.Run("TestTrackStatsAllNil", func(t *testing.T) {

	})

	// Test Route
	t.Run("TestTrackStatsRoute", func(t *testing.T) {

	})
}
