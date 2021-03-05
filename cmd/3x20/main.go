//package main
//
//import (
//	"fmt"
//	"fyne.io/fyne/v2/app"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/widget"
//	"github.com/J0eppp/3x20/pkg/beep"
//	"time"
//)
//
//func main() {
//	a := app.New()
//	w := a.NewWindow("Hello")
//
//	hello := widget.NewLabel("Hello Fyne!")
//	w.SetContent(container.NewVBox(
//		hello,
//		widget.NewButton("Hi!", func() {
//			hello.SetText("Welcome :)")
//		}),
//	))
//
//	w.ShowAndRun()
//	fmt.Println("3x20 is now running!")
//	err := beep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away")
//	//err := beeep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away", "assets/warning.png")
//	if err != nil {
//		panic(err)
//	}
//	endTime := time.Now().Add(1 * time.Minute)
//	for true {
//		if time.Now().After(endTime) {
//			// 20 minutes have passed
//			fmt.Println("Look away :)")
//			err := beep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away")
//			//err := beeep.Alert("3x20", "You have looked for 20 minutes at your screen, now please look away for 20 seconds 20 meters away", "assets/warning.png")
//			if err != nil {
//				panic(err)
//			}
//
//			// Wait for 20 seconds
//			newEndTime := time.Now().Add(20 * time.Second)
//			for true {
//				if time.Now().After(newEndTime) {
//					err = beep.Warn("3x20", "You can continue working now!")
//					//err = beeep.Notify("3x20", "You can continue working now!", "assets/information.png")
//					if err != nil {
//						panic(err)
//					}
//					endTime = time.Now().Add(1 * time.Minute)
//					break
//				}
//			}
//		}
//	}
//}

package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(1080, 600)
	window.SetWindowTitle("Hello World")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	// Add the close button
	button := widgets.NewQPushButton2("Close", nil)
	button.ConnectClicked(func(bool) {
		window.Close()
		app.CloseAllWindows()
	})
	widget.Layout().AddWidget(button)

	window.Show()

	app.Exec()
}