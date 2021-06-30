package files

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/darshanmarathe/goStocks/cli"
	"github.com/darshanmarathe/goStocks/models"
)

func CheckAndCreateFile(filename string) {
	_, err := os.Open(filename)
	if err != nil {
		data := make([]models.Stock, 0)
		file, _ := json.MarshalIndent(data, "", " ")
		_ = ioutil.WriteFile(filename, file, 0644)
	}
}

func ReadFile(filename string) []models.Stock {
	jsonFile, err := os.Open(filename)
	var stocks []models.Stock
	defer jsonFile.Close()
	if err == nil {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &stocks)
		return stocks
	} else {
		fmt.Println("Error in reading file ", filename)
	}
	return nil
}

func WriteFile(data []models.Stock, filename string) {
	fmt.Println(data)
	cli.ReadKey()
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("failed to parse the stocks", err)
		panic(err)
	}
	_ = ioutil.WriteFile(filename, file, fs.ModeAppend)
}
