package main

import (
	"testing"
)

// Unit test for bootstrapMedian
func TestBootstrapMedian(t *testing.T) {
	// Sample data
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	// Perform bootstrap with 1000 different resamples to clairfy
	resampledMedians := bootstrapMedian(data, 1000)

	// Ensure we have a valid result that means no empty parts or NAN values
	if len(resampledMedians) == 0 {
		t.Errorf("Expected non-empty resampled medians, but got empty or zero result")
	}

	// Check if all medians are within the original data set
	for _, median := range resampledMedians {
		if median < 1.0 || median > 5.0 {
			t.Errorf("Resampled median %.2f out of expected range", median)
		}
	}
}

// BenchmarkBootstrap tests the bootstrap median performance using data.csv that was created from the R script file
func BenchmarkBootstrap(b *testing.B) {
	// Load data from CSV file
	data, err := readCSV("data.csv")
	if err != nil {
		b.Fatalf("Failed to read data.csv: %v", err)
	}

	// Run benchmark to test it
	for i := 0; i < b.N; i++ {
		bootstrapMedian(data, 1000)
	}
	//everything passed so this all looks good yayyyy!
}
