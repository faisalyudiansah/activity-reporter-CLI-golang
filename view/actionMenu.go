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

func Action(socialGraph *controller.Activity) {
	finishSession := false
	scanner := bufio.NewScanner(os.Stdin)
	for !finishSession {
		input := promptInput(scanner, "Enter user Actions: ")
		inputToArray := strings.Split(input, " ")
		if len(inputToArray) < 3 || len(inputToArray) > 4 {
			utils.PrintError(variable.ErrorInvalidKeyword)
			finishSession = true
			break
		}
		switch inputToArray[1] {
		case "uploaded":
			if ok := isKeywordPhone(inputToArray, 2); !ok {
				continue
			}
			errMsg := socialGraph.UploadPhoto(inputToArray[0])
			if errMsg != nil {
				components.PrintErrorFromController(errMsg, 2)
				finishSession = true
				break
			}
			fmt.Println()
			finishSession = true
		case "likes":
			if ok := isKeywordPhone(inputToArray, 3); !ok {
				continue
			}
			errMsg := socialGraph.LikePhoto(inputToArray[0], inputToArray[2])
			if errMsg != nil {
				components.PrintErrorFromController(errMsg, 2)
				finishSession = true
				break
			}
			fmt.Println()
			finishSession = true
		default:
			utils.PrintError(variable.ErrorInvalidMenu)
			finishSession = true
		}
	}
}

func isKeywordPhone(input []string, idx int) bool {
	if input[idx] != "photo" {
		utils.PrintError(variable.ErrorInvalidKeyword)
		fmt.Println("Enter menu: 2")
		return false
	}
	return true
}
