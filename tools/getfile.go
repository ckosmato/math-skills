package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// GetFile reads integers from a specified file and returns a list of integers
func GetFile(inputFiles string) ([]float64, bool) {
	fileData, err := os.Open(inputFiles)
	empty := false
	if err != nil {
		fmt.Println("File doesn't exist!")
		os.Exit(0)

	}
	data := []string{}
	numList := []float64{}
	defer fileData.Close()
	file := bufio.NewScanner(fileData)
	for file.Scan() {
		data = append(data, file.Text())
		data = append(data, "\n")
	}
	for _, n := range data {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		numList = append(numList, float64(num))
	}
	if len(numList) == 0 {
		empty = true
	}
	return numList, empty
}
