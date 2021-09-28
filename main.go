package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/darshanmarathe/goStocks/cli"
	"github.com/darshanmarathe/goStocks/models"
	"github.com/darshanmarathe/goStocks/svc"
)

var stocks models.StocksType

func main() {
	Boot()
}

func Boot() {
	fmt.Println("Running stock app.")
	fmt.Println("=====================")
	stocks = svc.ReadStocks()
	fmt.Println("(", len(stocks), ") stocks")

reAsk:
	command := Menu()
	switch command {
	case 1:
		cli.Clear()
		stocks = svc.AddStock(stocks)
		Boot()
	case 2:
		stocks = svc.SellStock(stocks)
		Boot()
	case 3:
		svc.PrintTranctions(stocks)
		Boot()
	case 4:
		cli.Clear()
		svc.PrintStocks(stocks)
		cli.ReadKey()
		Boot()
	case 5:
		//cli.Clear()
		svc.ShowPnL()
		//cli.ReadKey()
		Boot()
	case 6:
		fmt.Println("*******************")
		fmt.Println("Shutting down.....")
	case 0:
		cli.Clear()
		fmt.Println("Input not understood try again.")
		goto reAsk
	}

}

func Menu() int {

	fmt.Println("Press 1 to record purches of stock")
	fmt.Println("Press 2 to record sell of stock")
	fmt.Println("Press 3 to print a stock and show all the transactions")
	fmt.Println("Press 4 to list all stocks")
	fmt.Println("Press 5 to do the P&L calc")
	fmt.Println("Press 6 to exit")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	i, _ := strconv.Atoi(text)

	return i
}
