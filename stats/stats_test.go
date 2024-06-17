package stats_test

import (
	"path/filepath"
	"testing"

	"statistics/stats"
)

func TestStats(t *testing.T) {
	data := []float64{100.0, 200.0, 300.0, 400.0, 500.0}
	// Test ReadDataFromFile function
	t.Run("ReadDataFromFile", func(t *testing.T) {
		// Provide a sample test file containing the values in 'data'
		dataFilePath := filepath.Join("..", "data.txt")
		result, err := stats.ReadDataFromFile(dataFilePath)
		if err != nil {
			t.Errorf("ReadDataFromFile returned an error: %v", err)
		}

		// Compare the result with expected data

		for i := range data {
			if result[i] != data[i] {
				t.Errorf("ReadDataFromFile returned incorrect data. Expected: %v, Got: %v", data[i], result[i])
			}
		}
	})

	// Test Average function
	t.Run("Average", func(t *testing.T) {
		expected := 300.0
		result := stats.Average(data)
		if result != expected {
			t.Errorf("Average returned incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})

	// Test Median function
	t.Run("Median", func(t *testing.T) {
		expected := 300.0
		result := stats.Median(data)
		if result != expected {
			t.Errorf("Median returned incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})

	// Test Variance function
	t.Run("Variance", func(t *testing.T) {
		expected := 20000.0
		result := stats.Variance(data)
		if result != expected {
			t.Errorf("Variance returned incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})

	// Test StdDev function
	t.Run("StdDev", func(t *testing.T) {
		expected := 141.4213562373095
		result := stats.StdDev(data)
		if result != expected {
			t.Errorf("StdDev returned incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})
}

// Replace with your actual package import path
