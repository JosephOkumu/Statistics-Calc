package stats

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ReadDataFromFile reads data from the provided filename and returns a slice of float64 values.
func ReadDataFromFile(filename string) ([]float64, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading from file: %v", err)
	}
	defer file.Close() // Close the file when the function returns

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		// Remove trailing commas and other non-numeric characters
		line = strings.Trim(line, ", ")

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Warning: skipping line with invalid number: %s\n", line)
			continue
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("error: file is empty or contains no valid data")
	}

	return data, nil
}

// Average calculates the mean of a slice of float64 values.
func Average(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Median calculates the median of a slice of float64 values.
func Median(data []float64) float64 {
	sort.Float64s(data) // Sort the data in ascending order
	n := len(data)
	if n%2 == 0 {
		// If the number of data points is even, the median is the average of the two middle values
		return (data[n/2-1] + data[n/2]) / 2
	}
	// If the number of data points is odd, the median is the middle value
	return data[n/2]
}

// Variance calculates the variance of a slice of float64 values.
func Variance(data []float64) float64 {
	mean := Average(data)
	var sumSquaredDiffs float64
	for _, value := range data {
		diff := value - mean
		sumSquaredDiffs += diff * diff
	}
	return sumSquaredDiffs / float64(len(data))
}

// StdDev calculates the standard deviation of a slice of float64 values.
func StdDev(data []float64) float64 {
	return math.Sqrt(Variance(data))
}
