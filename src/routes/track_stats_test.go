package routes

import (
	"bytes"
	"errors"
	"io"
	"log"
	mockUtils "main/mocks/utils"
	"main/src/types"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTrackStats(t *testing.T) {
	height := 12.7
	sm := 57
	ph := 4.7
	fert := 125.0
	temp := 87.0
	fullyFilledData := types.DataPoint{
		Timestamp:    time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Height:       &height,
		SoilMoisture: &sm,
		PH:           &ph,
		Fertilizer:   &fert,
		Tempurature:  &temp,
	}

	// Test Happy Path - All things filled out
	t.Run("TestTrackStatsAll", func(t *testing.T) {
		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(fullyFilledData).Return(nil)
		statusCode, _, err := trackStats(mockMongo, fullyFilledData)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	// Test Happy Path - Only Height
	t.Run("TestTrackStatsOnlyHeight", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Height:    &height,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(data).Return(nil)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	// Test Happy Path - Only Soil Moisture
	t.Run("TestTrackStatsOnlySoilMoisture", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp:    time.Now().Format("2006-01-02T15:04:05Z07:00"),
			SoilMoisture: &sm,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(data).Return(nil)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	// Test Happy Path - Only Fertilizer
	t.Run("TestTrackStatsOnlyFetilizer", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp:  time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Fertilizer: &fert,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(data).Return(nil)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	// Test Happy Path - Only ph
	t.Run("TestTrackStatsOnlyPH", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
			PH:        &ph,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(data).Return(nil)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	// Test Happy Path - Only Temp
	t.Run("TestTrackStatsOnlyTemp", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Tempurature: &temp,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(data).Return(nil)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	// Test Happy Path - Mixed Nil Stats
	t.Run("TestTrackStatsMixedNil", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp:    time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Height:       &height,
			SoilMoisture: &sm,
			PH:           nil,
			Fertilizer:   &fert,
			Tempurature:  &temp,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	// Test Sad Path 2 - Only timestamp
	t.Run("TestTrackStatsOnlyTimeStamp", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.Error(t, err)
	})

	// Test Sad Path - No time stamp
	t.Run("TestTrackStatsNoTimeStamp", func(t *testing.T) {
		data := types.DataPoint{
			Height:       &height,
			SoilMoisture: &sm,
			PH:           &ph,
			Fertilizer:   &fert,
			Tempurature:  &temp,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.Error(t, err)
	})

	// Test Sad Path - Invalid pH
	t.Run("TestTrackStatsInvalidPH", func(t *testing.T) {
		invalidPh := 1400.0
		data := types.DataPoint{
			Timestamp:    time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Height:       &height,
			SoilMoisture: &sm,
			PH:           &invalidPh,
			Fertilizer:   &fert,
			Tempurature:  &temp,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.Error(t, err)
	})

	// Test Sad Path - Empty Struct
	t.Run("TestTrackStatsEmpty", func(t *testing.T) {
		data := types.DataPoint{}
		mockMongo := mockUtils.NewMockMongoClient(t)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.Error(t, err)
	})

	// Test Sad Path - DB Error
	t.Run("TestTrackStatsDBError", func(t *testing.T) {
		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(fullyFilledData).Return(errors.New("Mongo Error"))
		statusCode, _, err := trackStats(mockMongo, fullyFilledData)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.Error(t, err)
	})

	// Test Sad Path - All Nil Stats
	t.Run("TestTrackStatsAllNil", func(t *testing.T) {
		data := types.DataPoint{
			Timestamp:    time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Height:       nil,
			SoilMoisture: nil,
			PH:           nil,
			Fertilizer:   nil,
			Tempurature:  nil,
		}
		mockMongo := mockUtils.NewMockMongoClient(t)
		statusCode, _, err := trackStats(mockMongo, data)
		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.Error(t, err)

	})

	// Test Route - Valid All
	t.Run("TestTrackStatsRouteValid", func(t *testing.T) {
		//The response recorder used to record HTTP responses
		respRec := httptest.NewRecorder()
		path := filepath.Join("..", "..", "test", "mock", "validTrackStatsBody.json")

		body, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer body.Close()

		req, err := http.NewRequest("POST", "/stats", body)
		if err != nil {
			log.Fatal(err)
		}

		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(mock.Anything).Return(nil) // TODO: Test Value
		router := SetupRouter(nil, mockMongo)

		router.ServeHTTP(respRec, req)

		assert.Equal(t, http.StatusOK, respRec.Code)
	})

	// Test Route - Optional Valid
	t.Run("TestTrackStatsRouteOptionalValid", func(t *testing.T) {
		//The response recorder used to record HTTP responses
		respRec := httptest.NewRecorder()

		path := filepath.Join("..", "..", "test", "mock", "validTrackStatsOptionalBody.json")

		body, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer body.Close()

		req, err := http.NewRequest("POST", "/stats", body)
		if err != nil {
			log.Fatal(err)
		}

		mockMongo := mockUtils.NewMockMongoClient(t)
		mockMongo.EXPECT().InsertStat(mock.Anything).Return(nil) // TODO: Test Value
		router := SetupRouter(nil, mockMongo)

		router.ServeHTTP(respRec, req)

		assert.Equal(t, http.StatusOK, respRec.Code)
	})

	// Test Route - Invalid Empty Body
	t.Run("TestTrackStatsRouteEmptyBody", func(t *testing.T) {
		//The response recorder used to record HTTP responses
		respRec := httptest.NewRecorder()
		mockMongo := mockUtils.NewMockMongoClient(t)
		router := SetupRouter(nil, mockMongo)

		emptyJSONBody := io.NopCloser(bytes.NewReader([]byte("{}")))

		req, err := http.NewRequest("POST", "/stats", emptyJSONBody)
		if err != nil {
			log.Fatal(err)
		}

		router.ServeHTTP(respRec, req)

		assert.Equal(t, http.StatusBadRequest, respRec.Code)
	})

	// Test Route - Invalid Body
	t.Run("TestTrackStatsRouteInvalidBody", func(t *testing.T) {
		//The response recorder used to record HTTP responses
		respRec := httptest.NewRecorder()
		mockMongo := mockUtils.NewMockMongoClient(t)
		router := SetupRouter(nil, mockMongo)

		path := filepath.Join("..", "..", "test", "mock", "invalidTrackStatsBody.json")

		body, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer body.Close()

		req, err := http.NewRequest("POST", "/stats", body)
		if err != nil {
			log.Fatal(err)
		}

		router.ServeHTTP(respRec, req)

		assert.Equal(t, http.StatusBadRequest, respRec.Code)
	})

	// Test Route 3 - Invalid Nil Body
	t.Run("TestTrackStatsRouteNilBody", func(t *testing.T) {
		//The response recorder used to record HTTP responses
		respRec := httptest.NewRecorder()
		mockMongo := mockUtils.NewMockMongoClient(t)
		router := SetupRouter(nil, mockMongo)

		nilJSONBody := io.NopCloser(bytes.NewReader(nil))

		req, err := http.NewRequest("POST", "/stats", nilJSONBody)
		if err != nil {
			log.Fatal(err)
		}

		router.ServeHTTP(respRec, req)

		assert.Equal(t, http.StatusBadRequest, respRec.Code)
	})
}
