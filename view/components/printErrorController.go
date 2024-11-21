package components

import "fmt"

func PrintErrorFromController(errMsg error, menu int) {
	fmt.Println()
	fmt.Println(errMsg)
	fmt.Println()
}
