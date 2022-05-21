package main
import ( 
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)
func readCsvFile(filePath string) ([][]string, error) {
    f, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Unable to read input file " + filePath, err)
		os.Exit(1)
    }

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        fmt.Println("Unable to parse file as CSV for " + filePath, err)
		os.Exit(1)
    }

	f.Close() // closes file
    return records, nil
}

func stringArrayToFloat (records [][] string) ([][]float64, error) {
	var err error = nil
	list := make([][]float64, len(records))
	for i := range list {
    	list[i] = make([]float64, 2)
	}
	for i, record := range records {
		for _, cell := range record {
			parts := strings.Split(cell, "\t")
			for index, element := range parts {
				list[i][index], err = strconv.ParseFloat(element, 64)
				if err != nil {
					fmt.Println("Unable to parse");
				}
			}
		}
	}
	return list, nil
}

func main() {
    records, err := readCsvFile("./test.csv")
	if err != nil {
		os.Exit(1)
	}
    //fmt.Println(records)
	newRecords, err := stringArrayToFloat(records)
	fmt.Println(newRecords)

	printNewRecords(newRecords)
}

func printNewRecords (newRecords [][] float64) {
	for i, row := range newRecords {
		fmt.Print("row is : ", i, " , ")
		for j, col := range row {
			fmt.Print("element ", j, " is ", col, " , ")
		}
		fmt.Println()
	}
}