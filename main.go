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
	list := make([][]float64, len(records))
	for i := range list {
    	list[i] = make([]float64, len(records[i]))
	}
	i := 0
	j := 0
	for _, record := range records {
		for _, cell := range record {
			parts := strings.Split(cell)
			
			for index, element := range parts {
				// index is the index where we are
    			// element is the element from someSlice for where we are
				
			}

			list[i][j], err = strconv.ParseFloat(cell, 64)
			if err != nil {
				//fmt.Println("Unable to parse")
				//os.Exit(1)
			}
			//fmt.Print(cell, " ")
			j++
		}
		i++
		j = 0
		//fmt.Println()
	}
	return list
}

func main() {
    records := readCsvFile("./test.csv")
    fmt.Println(records)
}