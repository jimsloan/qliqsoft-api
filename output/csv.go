package output

import (
	"encoding/csv"
	"fmt"
	"os"
)

// WriteToCsv ...
func WriteToCsv(rows [][]string) {

	path := "./data.csv"

	// remove file if it exist
	if _, err := os.Stat(path); err == nil {
		err := os.Remove(path)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("writeToCsv ... ")

	csvFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	for _, r := range rows {
		writer.Write(r)
	}
	writer.Flush()
}
