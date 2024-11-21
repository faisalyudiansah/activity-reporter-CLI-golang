package main

import (
	"bufio"
	"fmt"
	"os"

	"activity-reporter-cli/controller"
	"activity-reporter-cli/utils"
	"activity-reporter-cli/variable"
	"activity-reporter-cli/view"
)

func main() {
	RunCli()
}

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func RunCli() {
	socialGraph := controller.NewSocialGraph()
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	for !exit {
		view.MainMenu()
		input := promptInput(scanner, "Enter menu: ")
		switch input {
		case "1":
			view.Setup(socialGraph)
		case "2":
			view.Action(socialGraph)
		case "3":
			view.ActivityList(socialGraph)
		case "4":
			view.TrendingMenu(socialGraph)
		case "5":
			fmt.Println("Good bye!")
			exit = true
		default:
			utils.PrintError(variable.ErrorInvalidMenu)
		}
	}
}
