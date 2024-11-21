package view

import (
	"fmt"

	"activity-reporter-cli/controller"
)

func TrendingMenu(socialGraph *controller.Activity) {
	finishSession := false
	for !finishSession {
		fmt.Println()
		fmt.Println("Trending photos:")
		result := socialGraph.TrendingPhotos()
		for _, list := range result {
			fmt.Println(list)
		}
		fmt.Println()
		finishSession = true
	}
}
