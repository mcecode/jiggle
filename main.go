package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	robot "github.com/go-vgo/robotgo"
)

func main() {
	sx, sy := robot.GetScreenSize()
	startX := sx/2 + 100
	centerY := sy / 2

	a := app.New()
	w := a.NewWindow("Jiggle")

	w.Resize(fyne.NewSize(300, 150))

	// TODO: Improve UX.
	// Disable 'startBtn' while the mouse is moving.
	var startBtn *widget.Button
	// Remove 'stopBtn'.
	var stopBtn *widget.Button

	stop := make(chan bool)

	startBtn = widget.NewButton("Start", func() {
		w.SetContent(stopBtn)

		robot.Move(startX, centerY)

		go func() {
			for {
				select {
				case <-stop:
					return
				default:
					// TODO: Stop mouse movement with a keyboard event instead.
					if _, y := robot.GetMousePos(); y != centerY {
						w.SetContent(startBtn)
						stop <- true
					}

					// TODO: Control speed and distace traveled using user input.
					// TODO: Move mouse throughout the screen, not just in one line.
					robot.MoveSmoothRelative(-200, 0)
					robot.MilliSleep(100)
					robot.MoveSmoothRelative(200, 0)
					robot.MilliSleep(100)
				}
			}
		}()
	})

	stopBtn = widget.NewButton("Stop", func() {
		w.SetContent(startBtn)
		stop <- true
	})

	w.SetContent(startBtn)
	w.ShowAndRun()
}
