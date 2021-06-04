package svc

import (
	"fmt"
	"os"
	"sort"
	"time"

	//. "github.com/ahmetb/go-linq/v
	"github.com/darshanmarathe/goStocks/cli"
	"github.com/darshanmarathe/goStocks/files"
	"github.com/darshanmarathe/goStocks/models"
	"github.com/kataras/tablewriter"
	"github.com/lensesio/tableprinter"
)

const (
	datafile = "data.json"
)

func ReadStocks() []models.Stock {
	fmt.Println("Reading the stock")
	//check file
	files.CheckAndCreateFile(datafile)

	return files.ReadFile(datafile)
}

func AddStock(stocks []models.Stock) []models.Stock {
	symbol := cli.ReadString("Enter the stock symbol to add:")
	qty := cli.ReadInt("Enter the Qty:")
	buyPrice := cli.ReadFloat("Buying Price:")

	fmt.Println("symbol", symbol, "qty", qty, "buyPrice", buyPrice)
	stock := &models.Stock{
		Qty:      qty,
		Symbol:   symbol,
		BuyPrice: buyPrice,
		BuyValue: buyPrice * float64(qty),
		BuyDate:  time.Now(),
	}

	stocks = append(stocks, *stock)
	fmt.Println("symbol is", symbol, " added.")
	defer files.WriteFile(stocks, datafile)
	return stocks
}

func SellStock(stocks []models.Stock) []models.Stock {
	symbol := cli.ReadString("Enter the stock symbol to Sell")
	qty := cli.ReadInt("Enter the Qty:")
	sellPrice := cli.ReadFloat("Sell Price:")

	// currentStock := From(stocks).Where(func(c interface{}) boo {
	// 	return c(models.Stock).Symbol == symol
	// }).Firs()

	fmt.Println("symbol", symbol, qty, sellPrice)

	//TODO :: Remove the logic
	return stocks
}

func PrintStocks(stocks []models.Stock) {
	printer := tableprinter.New(os.Stdout)

	stocks_sorted := make(models.StocksType, 0, len(stocks))
	for _, d := range stocks {
		stocks_sorted = append(stocks_sorted, d)
	}

	sort.Sort(stocks_sorted)

	//Optionally, customize the table, import of the underline 'tablewriter' package is required for that.
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor
	printer.Print(stocks_sorted)
	//Print the slice of structs as table, as shown above.
	//tableprinter.Print(os.Stdout, stocks)
}
