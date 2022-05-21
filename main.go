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
    	list[i] = make([]float64, len(records[i]))
	}
	i := 0
	for _, record := range records {
		for _, cell := range record {
			parts := strings.Split(cell, "\t")
			for index, element := range parts {
				// index is the index where we are
    			// element is the element from someSlice for where we are
				fmt.Println(strconv.ParseFloat(element, 64))
				list[i][index], err = strconv.ParseFloat(element, 64)
				if err != nil {
					fmt.Println("Unable to parse");
				}
			}
		}
		i++
	}
	return list, nil
}

func main() {
    records, err := readCsvFile("./test.csv")
	if err != nil {
		os.Exit(1)
	}
    fmt.Println(records)
	
}