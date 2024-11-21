package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"activity-reporter-cli/controller"
	"activity-reporter-cli/utils"
	"activity-reporter-cli/variable"
	"activity-reporter-cli/view/components"
)

func Setup(socialGraph *controller.Activity) {
	finishSession := false
	scanner := bufio.NewScanner(os.Stdin)
	for !finishSession {
		input := promptInput(scanner, "Setup social graph: ")
		inputToArray := strings.Split(input, " ")
		if len(inputToArray) != 3 {
			utils.PrintError(variable.ErrorInvalidKeyword)
			finishSession = true
			break
		}
		switch inputToArray[1] {
		case "follows":
			errMsg := socialGraph.FollowUser(inputToArray[0], inputToArray[2])
			if errMsg != nil {
				components.PrintErrorFromController(errMsg, 1)
				finishSession = true
				break
			}
			fmt.Println()
			finishSession = true
		default:
			utils.PrintError(variable.ErrorInvalidKeyword)
			finishSession = true
		}
	}
}
