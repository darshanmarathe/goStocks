package cli

import (
	"bufio"
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
	text := ReadString(msg)
	i, _ := strconv.Atoi(text)
	return i

}

func ReadFloat(msg string) float64 {
	text := ReadString(msg)
	i, _ := strconv.ParseFloat(text, 64)
	return i
}

func ReadKey() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}
