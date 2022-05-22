package main
import ( 
	"fmt" // for file io
	"encoding/csv" // for processing csv
	"os" // for os interaction
	"strconv" // for float conversion
	"strings" // for substring creation
)
func readCsvFile(filePath string) ([][]string, error) {
    f, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Unable to read input file " + filePath, err)
		return nil, err
    }

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        fmt.Println("Unable to parse file as CSV for " + filePath, err)
		return nil, err
    }

	f.Close() // closes file
    return records, nil
}

func stringArrayToFloat (records [][] string) ([][]float64, error) {
	var err error = nil // just in case error needs to be returned.

	list := make([][]float64, len(records)) // dynamic memory allocation creates 2d array

	for i := range list {
    	list[i] = make([]float64, 2) // creates sub arrays, of length 2.
	}
	j := 0
	// loops through each row
	for i, record := range records {
		// loops through each column
		for _, cell := range record {
			// splits the text that is in the columns, so it can be stored in the float64 columns
			parts := strings.Split(cell, "\t")
			// loops through each part that is created to parse it
			for _, element := range parts {
				//fmt.Println(strconv.ParseFloat(element, 64));
				list[i][j], err = strconv.ParseFloat(element, 64) // parse array and store in list index.
				j++
				if err != nil {
					return list, err
				}
			}
		}
		j = 0
	}
	return list, nil
}

func printNewRecords (newRecords [][] float64) {
	// loops through indexes i with item row
	for i, row := range newRecords {
		fmt.Print("row is : ", i, " , ")
		// loops through column #j with value column
		for j, col := range row {
			fmt.Print("element ", j, " is ", col, " , ")
		}
		fmt.Println()
	}
}

func averageWordsPerMinute (arr [][] float64) int {
	var sum float64 = 0 
	var length float64 = 0
	for _, row := range arr {
		length += row[0]
		sum += row[1] * row[0]
	}
	return int(sum/length) // change to get words and time,  
}

func writeToCsv (column1 string, column2 string) (bool, error) {
	f, err := os.OpenFile("../test.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return false, err
	}
	arr := [] string {column1, column2}
	w := csv.NewWriter(f)
	w.Write(arr)
	w.Flush()
	return true, nil
}

//func main() {
    /*records, err := readCsvFile("./test.csv")
	if err != nil {
		os.Exit(1)
	}

	newRecords, err := stringArrayToFloat(records)
	
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(newRecords)
	printNewRecords(newRecords)
	writeToCsv("1","2")*/
//}