package beep

import "github.com/gen2brain/beeep"

func Alert(title string, message string) error {
	return beeep.Alert(title, message, "assets/warning.png")
}

func Warn(title string, message string) error {
	return beeep.Notify(title, message, "assets/information.png")
}