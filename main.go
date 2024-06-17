package main

import (
	"fmt"
	"math"
	"os"
	"statistics/stats"

)

func main() {
	// Check if the user provided a filename as a command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}

	filename := os.Args[1]
	data, err := stats.ReadDataFromFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Calculate the average, median, variance, and standard deviation
	avg := int(math.Round(stats.Average(data)))
	med := int(math.Round(stats.Median(data)))
	varnc := int(math.Round(stats.Variance(data)))
	stddev := int(math.Round(stats.StdDev(data)))

	// Print the results
	fmt.Printf("Average: %d\n", avg)
	fmt.Printf("Median: %d\n", med)
	fmt.Printf("Variance: %d\n", varnc)
	fmt.Printf("Standard Deviation: %d\n", stddev)
}
