package main

import (
	"fmt"
	"os"

	"math-skills/subtools"
	"math-skills/tools"
)

// The program takes a file path as an argument, containing a list of numbers for statistical analysis.

func main() {
	os.Args[0] = "math.go"
	if len(os.Args) < 2 {
		fmt.Println("No file path in arguments!")
		os.Exit(0)
	}
	numList, empty := tools.GetFile(os.Args[1])
	if empty {
		fmt.Println("No numbers in file!")
		os.Exit(0)
	}
	averageTemp := tools.GetAverage(numList)
	average := subtools.Rounded(averageTemp)
	median := tools.GetMedian(numList)
	varianceTemp := tools.GetVariance(numList)
	variance := subtools.Rounded(varianceTemp)
	standardDevation := tools.GetStandardDeviation(numList)

	fmt.Printf("Average: %v\n", average)
	fmt.Printf("Median: %v\n", median)
	fmt.Printf("Variance: %v\n", variance)
	fmt.Printf("Standard Deviation: %v\n", standardDevation)
}
