# Math Skills Application

This application performs statistical analysis on a list of numbers provided in a text file. It calculates the average, median, variance, standard deviation, and the largest integer whose square is less than or equal to the average of the numbers, and outputs the results in a formatted manner.

## Introduction

The Math Skills application allows you to analyze a series of numbers contained within a file. This tool is designed to help users easily perform basic statistical calculations.

## Usage

Follow these steps to run the application:

1. **Run the Application**  
   Execute the application by providing the file containing the list of numbers:
   
   ```
   go run math.go example.txt
   ```
2. **Check Results**  
   If everything goes well, you will see the calculated statistics:
   ```
   Average: <value>
   Median: <value>
   Variance: <value>
   Standard Deviation: <value>
   ```

## Author

- Christos Kosmatos @ckosmato

## Key Functions and Modules

- `tools.GetAverage`: Calculates and returns the average of a list of integers.
- `tools.GetMedian`: Computes the median of a sorted list of integers.
- `tools.GetVariance`: Calculates the variance of a list of integers.
- `tools.GetStandardDeviation`: Computes the standard deviation based on variance.
- `subtools.Rounded`: Rounds a number to the nearest integer.
- `subtools.GetSorted`: Sorts a slice of integers using a bubble sort algorithm.
- `subtools.FindSquare`: Finds the largest integer whose square is less than or equal to a given number.

## Libraries and Dependencies

- Go Standard Library:
  - `os`: For file operations.
  - `bufio`: For reading input files.
  - `strconv`: For converting strings to integers.
  - `fmt`: For formatted I/O.

## Data Flow Diagram

1. The program takes a file path as an argument containing a list of numbers for statistical analysis.
2. The file is read, and integers are extracted.
3. Statistical functions compute the average, median, variance, standard deviation, and the largest square.
4. The results are printed to the console.

## Performance Considerations

- Efficient parsing ensures quick file reading and processing.
- Robust error handling allows the program to manage unexpected inputs gracefully.
- Stateless processing enables the application to handle individual requests independently.

## Limitations

- Supports only integer inputs.
- Requires the input file to be formatted correctly without any non-numeric data.
