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
					XYValues:        nil,
					MaxProfit:       0,
					MaxLoss:         0,
					BreakEvenPoints: nil,
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
