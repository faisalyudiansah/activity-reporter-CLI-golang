package view

import (
	"fmt"
)

func MainMenu() {
	fmt.Println("Activity Reporter")
	fmt.Println()
	menu := "1. Setup\n" +
		"2. Action\n" +
		"3. Display\n" +
		"4. Trending\n" +
		"5. Exit\n"
	fmt.Println(menu)
}
