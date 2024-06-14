package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/common"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/optionsContract"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/optionsContract/optionsContractMock"
)

var (
	TestDataFileName = "testdata/testdata.json"
	Tolerance        = 1e-9
)

func TestAnalysisController(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		testOptsService := optionsContractMock.New(
			optionsContractMock.WithAnalysis(func(opts []*optionsContract.Options) (*optionsContract.OptionsAnalysis, error) {
				return &optionsContract.OptionsAnalysis{
					XYValues:        []optionsContract.XYValue{{X: 100, Y: 10}},
					MaxProfit:       10,
					MaxLoss:         10,
					BreakEvenPoints: []float64{1.0, 2.0},
				}, nil
			}),
		)

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(testOptsService)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze", testObjectsFromTestData(t))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusOK)
			t.Fatal("invalid response")
		}
	})

	t.Run("invalid request", func(t *testing.T) {
		t.Parallel()

		testOptsService := optionsContractMock.New(
			optionsContractMock.WithAnalysis(func(opts []*optionsContract.Options) (*optionsContract.OptionsAnalysis, error) {
				return nil, fmt.Errorf("some error")
			}),
		)

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(testOptsService)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze", testObjectsFromTestData(t))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusInternalServerError)
			t.Fatal("invalid response")
		}
	})

	t.Run("service error", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze", testObjectsFromTestData(t))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusInternalServerError)
			t.Fatal("invalid response")
		}
	})

	t.Run("bid > ask", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze",
			OptionsToBuffer(t, []*OptionsContract{{
				StrikePrice:    100,
				Type:           "Call",
				Bid:            12.05,
				Ask:            12.04,
				LongShort:      "long",
				ExpirationDate: time.Date(2025, 12, 17, 0, 0, 0, 0, time.UTC),
			}}))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusBadRequest)
			t.Fatal("invalid response")
		}
	})

	t.Run("Bid < 0 || Ask < 0", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze",
			OptionsToBuffer(t, []*OptionsContract{{
				StrikePrice:    100,
				Type:           "Call",
				Bid:            -10.05,
				Ask:            12.04,
				LongShort:      "long",
				ExpirationDate: time.Date(2025, 12, 17, 0, 0, 0, 0, time.UTC),
			}}))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusBadRequest)
			t.Fatal("invalid response")
		}
	})

	t.Run("Ask < 0", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze",
			OptionsToBuffer(t, []*OptionsContract{{
				StrikePrice:    100,
				Type:           "Call",
				Bid:            10.05,
				Ask:            -12.04,
				LongShort:      "long",
				ExpirationDate: time.Date(2025, 12, 17, 0, 0, 0, 0, time.UTC),
			}}))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusBadRequest)
			t.Fatal("invalid response")
		}
	})

	t.Run("ExpirationDate.Before(currentTime)", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze",
			OptionsToBuffer(t, []*OptionsContract{{
				StrikePrice:    100,
				Type:           "Call",
				Bid:            10.05,
				Ask:            12.04,
				LongShort:      "long",
				ExpirationDate: time.Date(2022, 12, 17, 0, 0, 0, 0, time.UTC),
			}}))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusBadRequest)
			t.Fatal("invalid response")
		}
	})

	t.Run("StrikePrice < 0", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze",
			OptionsToBuffer(t, []*OptionsContract{{
				StrikePrice:    -100,
				Type:           "Call",
				Bid:            10.05,
				Ask:            12.04,
				LongShort:      "long",
				ExpirationDate: time.Date(2025, 12, 17, 0, 0, 0, 0, time.UTC),
			}}))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusBadRequest)
			t.Fatal("invalid response")
		}
	})

	t.Run("incorrect position - long || short", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze",
			OptionsToBuffer(t, []*OptionsContract{{
				StrikePrice:    100,
				Type:           "Call",
				Bid:            10.05,
				Ask:            12.04,
				LongShort:      "someWrong",
				ExpirationDate: time.Date(2025, 12, 17, 0, 0, 0, 0, time.UTC),
			}}))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusBadRequest)
			t.Fatal("invalid response")
		}
	})

	t.Run("incorrect type - put || call", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze",
			OptionsToBuffer(t, []*OptionsContract{{
				StrikePrice:    100,
				Type:           "kol",
				Bid:            10.05,
				Ask:            12.04,
				LongShort:      "long",
				ExpirationDate: time.Date(2025, 12, 17, 0, 0, 0, 0, time.UTC),
			}}))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusBadRequest)
			t.Fatal("invalid response")
		}
	})

	t.Run("incorrect binding", func(t *testing.T) {
		t.Parallel()

		newServerContainer := NewTestServer(t, WithAnalysisController(NewAnalysisController(nil)))

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/analyze",
			bytes.NewBufferString(`invalid json`))
		if err != nil {
			t.Fatal("unable to create a new request: ", err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		newServerContainer.r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			fmt.Println("Actual Status Code:", rr.Code)
			fmt.Println("Expected Status Code:", http.StatusBadRequest)
			t.Fatal("invalid response")
		}
	})
}

func testObjectsFromTestData(t *testing.T) *bytes.Buffer {
	t.Helper()

	rootDir, err := common.FindProjectRoot()
	if err != nil {
		t.Fatal("unable to find root path: ", err)
	}

	filePath := filepath.Join(rootDir, TestDataFileName)
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal("unable to open the test data file: ", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatal("unable to read the file: ", err)
	}

	var options []*OptionsContract
	err = json.Unmarshal(data, &options)
	if err != nil {
		t.Fatal("unable to unmarshal the test data into model: ", err)
	}

	return bytes.NewBuffer(data)
}

func OptionsToBuffer(t *testing.T, options []*OptionsContract) *bytes.Buffer {
	t.Helper()

	jsonData, err := json.Marshal(options)
	if err != nil {
		t.Fatal("unable to marshal an object: ", err)
	}

	buffer := bytes.NewBuffer(jsonData)
	return buffer
}
