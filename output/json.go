package output

import (
	"fmt"
	"os"
	"strconv"
)

// WriteToJSON ...
func WriteToJSON(page int, data []byte) error {

	fmt.Println("writeToJSON ... page:" + strconv.Itoa(page))

	filename := "page-" + strconv.Itoa(page) + ".json"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return file.Sync()
}
