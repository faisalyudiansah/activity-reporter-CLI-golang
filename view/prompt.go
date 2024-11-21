package view

import (
	"bufio"
	"fmt"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}
