package optionsContract

import (
	"encoding/json"
	"io"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/common"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/marketPrice"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/marketPrice/marketPriceMock"
)

var (
	TestDataFileName = "testdata/testdata.json"
	Tolerance        = 1e-9
)

func TestOptionSvc_Analysis(t *testing.T) {
	t.Parallel()

	optionsSvc := NewOptionsSvc(marketPriceMockForOptions(t))

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedOptionsAnalysis := &OptionsAnalysis{
			XYValues:        make([]XYValue, 4),
			MaxProfit:       11.00,
			MaxLoss:         39.04,
			BreakEvenPoints: make([]float64, 4),
		}

		expectedOptionsAnalysis.XYValues[0] = XYValue{X: 100, Y: -12.04}
		expectedOptionsAnalysis.XYValues[1] = XYValue{X: 100, Y: -14}
		expectedOptionsAnalysis.XYValues[2] = XYValue{X: 100, Y: 11}
		expectedOptionsAnalysis.XYValues[3] = XYValue{X: 100, Y: -13}

		expectedOptionsAnalysis.BreakEvenPoints = []float64{112.04, 116.5, 89, 87}

		analysis, err := optionsSvc.Analysis(testObjectsFromTestData(t))
		if err != nil {
			t.Fatal(err)
		}

		for i, v := range analysis.XYValues {
			t.Log("value of x = price = ", v.X, "\tvalue of y = risk or reward = ", v.Y)
			if expectedOptionsAnalysis.XYValues[i] != v {
				t.Fatal("FAIL: expected: ", expectedOptionsAnalysis.XYValues[i], "\tgot: ", v)
			}
		}

		t.Log("max profit: ", analysis.MaxProfit)
		if expectedOptionsAnalysis.MaxProfit != analysis.MaxProfit {
			t.Fatal("FAIL: expected: ", expectedOptionsAnalysis.MaxProfit, "\tgot: ", analysis.MaxProfit)
		}

		t.Log("max loss: ", analysis.MaxLoss)
		if expectedOptionsAnalysis.MaxLoss != analysis.MaxLoss {
			t.Fatal("FAIL: expected: ", expectedOptionsAnalysis.MaxLoss, "\tgot: ", analysis.MaxLoss)
		}

		for i, v := range analysis.BreakEvenPoints {
			t.Log("break-even point: ", v)
			if !(math.Abs(expectedOptionsAnalysis.BreakEvenPoints[i]-v) <= Tolerance) {
				t.Fatal("FAIL: expected: ", expectedOptionsAnalysis.BreakEvenPoints[i], "\tgot: ", v)
			}
		}
	})
}

func marketPriceMockForOptions(t *testing.T) marketPrice.Service {
	t.Helper()

	return marketPriceMock.New(
		marketPriceMock.WithGet(func() (*marketPrice.Token, error) {
			return &marketPrice.Token{
				Symbol:  "TEST",
				Chain:   "TESTCHAIN",
				Network: "TESTNETWORK",
				Price:   100,
			}, nil
		}),
	)
}

func testObjectsFromTestData(t *testing.T) []*Options {
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

	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatal("unable to read the file: ", err)
	}

	var options []*Options
	err = json.Unmarshal(bytes, &options)
	if err != nil {
		t.Fatal("unable to unmarshal the test data into model: ", err)
	}

	return options
}
