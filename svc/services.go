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
	symbol = strings.ToUpper(symbol)
	nStock, index, err := GetStock(stocks, symbol)
	time := time.Now()
	if err != nil {
		stock := &models.Stock{
			Qty:        qty,
			Symbol:     symbol,
			BuyPrice:   buyPrice,
			BuyValue:   buyPrice * float64(qty),
			BuyDate:    time,
			BuyDateStr: time.Format("01 Jan 2006"),
			TransType:  "BUY",
		}

		nStock = CalculatePnL(nStock)
		stock.AuditTrail = append(stock.AuditTrail, *stock)
		stocks = append(stocks, *stock)
		fmt.Println("symbol is", symbol, " added.")

	} else {
		stock := &models.Stock{
			Qty:        qty,
			Symbol:     symbol,
			BuyPrice:   buyPrice,
			BuyValue:   buyPrice * float64(qty),
			BuyDate:    time,
			BuyDateStr: time.Format("01 Jan 2006"),
			TransType:  "BUY",
		}

		nStock.AuditTrail = append(nStock.AuditTrail, *stock)
		nStock.Qty = nStock.Qty + qty
		nStock.BuyValue = nStock.BuyValue + (buyPrice * float64(qty))
		nStock.BuyPrice = nStock.BuyValue / float64(nStock.Qty)
		nStock.BuyDate = stock.BuyDate
		nStock.BuyDateStr = stock.BuyDateStr
		nStock = CalculatePnL(nStock)
		//cli.Stringfy(nStock)
		stocks[index] = nStock
	}

	defer files.WriteFile(stocks, datafile)
	return stocks
}

func CalculatePnL(stock models.Stock) models.Stock {
	var buyToltal float64
	var SellToltal float64

	//calculate total buy and update
	for _, v := range stock.AuditTrail {
		if v.TransType == "BUY" {
			buyToltal += v.BuyValue
		}
	}

	//calculate total sell and update

	for _, v := range stock.AuditTrail {
		if v.TransType == "SELL" {
			SellToltal += v.SellValue
		}
	}

	//calc pnl and update
	stock.Pnl = SellToltal - buyToltal

	return stock
}

func SellStock(stocks []models.Stock) []models.Stock {
	symbol := cli.ReadString("Enter the stock symbol to Sell")
	qty := cli.ReadInt("Enter the Qty:")
	sellPrice := cli.ReadFloat("Sell Price:")

	fmt.Println("symbol", symbol, qty, sellPrice)

	symbol = strings.ToUpper(symbol)
	nStock, index, err := GetStock(stocks, symbol)
	if err != nil {
		fmt.Printf("Sorry %q is not found in collections", symbol)
		return stocks
	} else {
		if nStock.Qty < qty {
			fmt.Printf("Stock is less than sell quntity %q for symbol %q avaible quntity is (%q)", qty, symbol, nStock.Qty)
			return stocks
		}
		_time := time.Now()
		stock := &models.Stock{
			Qty:         qty,
			Symbol:      symbol,
			SellPrice:   sellPrice,
			SellValue:   sellPrice * float64(qty),
			SellDate:    _time,
			SellDateStr: _time.Format("01 Jan 2006"),
			TransType:   "SELL",
		}

		nStock.AuditTrail = append(nStock.AuditTrail, *stock)
		nStock.Qty = nStock.Qty - qty
		nStock.BuyValue = float64(nStock.Qty) * nStock.BuyPrice
		nStock.SellDate = stock.SellDate
		nStock.SellDateStr = stock.SellDateStr
		nStock.SellPrice = stock.SellPrice
		nStock.SellValue = stock.SellValue

		nStock = CalculatePnL(nStock)
		//cli.Stringfy(nStock)
		stocks[index] = nStock
		fmt.Println(stocks, "Writing....")
		files.WriteFile(stocks, datafile)
	}

	return stocks

}

func PrintTranctions(stocks []models.Stock) {
	symbol := cli.ReadString("Enter the stock symbol to add:")
	nStock, _, err := GetStock(stocks, symbol)
	if err != nil {
		fmt.Printf("symbol not found %q \n", symbol)
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
		if strings.ToUpper(st.Symbol) == strings.ToUpper(symbol) {
			st = CalculatePnL(st)
			return st, i, nil
		}

	}
	return stToRet, 0, errors.New("Can not find " + symbol)
}
