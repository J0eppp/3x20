package main

import (
	"fmt"
	"github.com/J0eppp/3x20/pkg/beep"
	"math"
	"time"
)

func main() {
	fmt.Println("3x20 is now running!")
	// err := beep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away")
	// //err := beeep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away", "assets/warning.png")
	// if err != nil {
	//	panic(err)
	// }
	endTime := time.Now().Add(20 * time.Minute)
	for true {
		end := endTime.Unix()
		now := time.Now().Unix()
		diff := end - now
		minutes := int(math.Floor(float64(diff / 60)))
		seconds := diff % 60
		fmt.Printf("Time left: %d:%d           \r", minutes, seconds)
		// now := time.Now()
		// fmt.Printf("Time left: %d:%d  \r", endTime.Minute()-now.Minute()-1, 60+(endTime.Second()-now.Second()))
		if time.Now().After(endTime) {
			// 20 minutes have passed
			fmt.Printf("Look away :)           \r")
			err := beep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away")
			// err := beeep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away", "assets/warning.png")
			if err != nil {
				panic(err)
			}

			// Wait for 20 seconds
			newEndTime := time.Now().Add(20 * time.Second)
			for true {
				if time.Now().After(newEndTime) {
					err = beep.Warn("3x20", "You can continue working now!")
					// err = beeep.Notify("3x20", "You can continue working now!", "assets/information.png")
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
