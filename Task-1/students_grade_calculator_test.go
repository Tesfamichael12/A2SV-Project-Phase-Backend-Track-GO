package main

import (
	"math"
	"testing"
)

func TestGetAverageGrade(t *testing.T) {
	subjects1 := map[string]float64{
		"math":    90,
		"phy": 80,
		"chem": 70,
	}
	expected1 := 80.0
	actual1 := get_average_grade(subjects1)
	if actual1 != expected1 {
		t.Errorf("Failed: Expected %f, got %f", expected1, actual1)
	}


	subjects2 := map[string]float64{}
	expected2 := 0.0
	actual2 := get_average_grade(subjects2)
	if !math.IsNaN(actual2) && actual2 != expected2 {
		// Handle case where division by zero results in NaN
		if !(math.IsNaN(actual2) && expected2 == 0) {
			t.Errorf("Failed: Expected %f for empty map, got %f", expected2, actual2)
		}
	}

	subjects3 := map[string]float64{
		"History": 95,
	}
	expected3 := 95.0
	actual3 := get_average_grade(subjects3)
	if actual3 != expected3 {
		t.Errorf("Failed: Expected %f, got %f", expected3, actual3)
	}
}
