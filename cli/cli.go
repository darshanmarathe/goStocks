package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadString(msg string) string {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	return text

}

func ReadInt(msg string) int {
reAsk:
	text := ReadString(msg)
	i, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("Given value is not a valid int %q ", text)
		goto reAsk
	}
	return i

}

func ReadFloat(msg string) float64 {
reAsk:
	text := ReadString(msg)
	i, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Printf("Given value is not a valid float %q ", text)
		goto reAsk
	}
	return i
}

func ReadKey() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func Stringfy(data interface{}) {
	str, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(str))
}
