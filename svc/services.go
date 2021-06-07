package svc

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	//. "github.com/ahmetb/go-linq/v
	"github.com/darshanmarathe/goStocks/cli"
	"github.com/darshanmarathe/goStocks/files"
	"github.com/darshanmarathe/goStocks/models"
	"github.com/lensesio/tableprinter"
	"github.com/olekukonko/tablewriter"
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

	nStock, index, err := GetStock(stocks, symbol)
	if err != nil {
		stock := &models.Stock{
			Qty:      qty,
			Symbol:   symbol,
			BuyPrice: buyPrice,
			BuyValue: buyPrice * float64(qty),
			BuyDate:  time.Now(),
		}

		stock.AuditTrail = append(stock.AuditTrail, *stock)
		stocks = append(stocks, *stock)
		fmt.Println("symbol is", symbol, " added.")

	} else {
		stock := &models.Stock{
			Qty:      qty,
			Symbol:   symbol,
			BuyPrice: buyPrice,
			BuyValue: buyPrice * float64(qty),
			BuyDate:  time.Now(),
		}

		nStock.AuditTrail = append(nStock.AuditTrail, *stock)
		nStock.Qty = nStock.Qty + qty
		nStock.BuyValue = nStock.BuyValue + (buyPrice * float64(qty))
		nStock.BuyPrice = nStock.BuyValue / float64(nStock.Qty)

		//cli.Stringfy(nStock)
		stocks[index] = nStock
	}

	defer files.WriteFile(stocks, datafile)
	return stocks
}

func SellStock(stocks []models.Stock) []models.Stock {
	symbol := cli.ReadString("Enter the stock symbol to Sell")
	qty := cli.ReadInt("Enter the Qty:")
	sellPrice := cli.ReadFloat("Sell Price:")

	fmt.Println("symbol", symbol, qty, sellPrice)

	//TODO :: Remove the logic
	return stocks

}

func PrintTranctions(stocks []models.Stock) {
	symbol := cli.ReadString("Enter the stock symbol to add:")
	nStock, _, err := GetStock(stocks, symbol)
	if err != nil {
		fmt.Println("symbol not found %q", symbol)
		return
	}
	t := []models.Stock{nStock}
	PrintStocks(t)
	cli.ReadKey()
	PrintStocks(nStock.AuditTrail)

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
}

func GetStock(stocks []models.Stock, symbol string) (models.Stock, int, error) {

	var stToRet models.Stock
	for i, st := range stocks {
		if strings.ToLower(st.Symbol) == strings.ToLower(symbol) {
			return st, i, nil
		}

	}
	return stToRet, 0, errors.New("Can not find " + symbol)
}
