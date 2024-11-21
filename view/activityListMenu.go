package view

import (
	"bufio"
	"fmt"
	"os"

	"activity-reporter-cli/controller"
	"activity-reporter-cli/utils"
)

func ActivityList(socialGraph *controller.Activity) {
	finishSession := false
	scanner := bufio.NewScanner(os.Stdin)
	for !finishSession {
		input := promptInput(scanner, "Display activity for: ")
		result, errMsg := socialGraph.ActivityUser(input)
		if errMsg != nil {
			utils.PrintError(errMsg)
			finishSession = true
			break
		}
		fmt.Println()
		fmt.Printf("%v activities:\n", input)
		for _, list := range result {
			fmt.Println(list)
		}
		fmt.Println()
		finishSession = true
	}
}
