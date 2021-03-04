package main

import (
	"github.com/gen2brain/beeep"
	"time"
)

func main() {
	err := beeep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away", "assets/warning.png")
	if err != nil {
		panic(err)
	}
	endTime := time.Now().Add(20 * time.Minute)
	for true {
		if time.Now().After(endTime) {
			// 20 minutes have passed
			err := beeep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away", "assets/warning.png")
			if err != nil {
				panic(err)
			}

			// Wait for 20 seconds
			newEndTime := time.Now().Add(20 * time.Second)
			for true {
				if time.Now().After(newEndTime) {
					err = beeep.Notify("3x20", "You can continue working now!", "assets/information.png")
					if err != nil {
						panic(err)
					}
					endTime = time.Now().Add(20 * time.Minute)
					break
				}
			}
		}
	}
}