package main

import (
	"encoding/csv"
	"fmt"
	"log" //import for logging
	"math"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // Import pprof for profiling
	"os"
	"sort"
	"strconv"

	"github.com/seehuhn/mt19937" //import for the suggested random sampling genorator from canvas
)

// Read dataset from CSV file (ignores header) so that x is not included from the header value
func readCSV(filename string) ([]float64, error) {
	log.Println("Starting to read CSV file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []float64
	for i, row := range rows {
		if i == 0 {
			continue // Skip the header row aka x
		}
		if len(row) > 0 {
			val, err := strconv.ParseFloat(row[0], 64)
			if err == nil {
				data = append(data, val)
			}
		}
	}
	log.Printf("Finished reading CSV file. Loaded %d data points.\n", len(data))
	return data, nil
}

// Compute the median of a dataset
func median(data []float64) float64 {
	sort.Float64s(data)
	mid := len(data) / 2
	if len(data)%2 == 0 {
		return (data[mid-1] + data[mid]) / 2.0
	}
	return data[mid]
}

// Perform bootstrap resampling with mt19937 PRNG which is a Merseenne twister from Canvas suggestion
func bootstrapMedian(data []float64, R int) []float64 {
	log.Println("Starting bootstrap resampling with", R, "resamples.")
	n := len(data)
	resampledMedians := make([]float64, R)

	// Use Mersenne Twister for deterministic results
	rng := rand.New(mt19937.New())
	rng.Seed(42)

	for i := 0; i < R; i++ {
		resample := make([]float64, n)
		for j := range resample {
			resample[j] = data[rng.Intn(n)]
		}
		resampledMedians[i] = median(resample)
	}
	log.Println("Bootstrap resampling completed.")
	return resampledMedians
}

// Compute standard error of the median
func standardError(medians []float64) float64 {
	sum, sumSq := 0.0, 0.0
	for _, v := range medians {
		sum += v
		sumSq += v * v
	}
	mean := sum / float64(len(medians))
	variance := (sumSq / float64(len(medians))) - (mean * mean)
	return math.Sqrt(variance) // Return the square root to get standard error
}

func main() {
	// Start profiling required from canvas assignment
	go func() {
		log.Println("Starting pprof HTTP server on localhost:6060")
		log.Fatal(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("Program started")

	// Load data from CSV file this is from the R script so that they are the same dataset so we can check for the same results
	data, err := readCSV("data.csv")
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	// Perform bootstrap
	bootstrappedMedians := bootstrapMedian(data, 1000)
	log.Printf("Performed bootstrap resampling. Number of medians: %d", len(bootstrappedMedians))

	// Compute and log the standard error logging was required from canvas
	se := standardError(bootstrappedMedians)
	log.Printf("Standard Error of Median: %.5f", se)

	// Print the result
	fmt.Printf("Bootstrap Standard Error of Median: %.5f\n", se)

	// Log completion
	log.Println("Program completed.")
}

// we got 0.03269 for the st error of median which is the same as R ya!
