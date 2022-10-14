package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	robot "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var (
	sx, sy = robot.GetScreenSize()
	startX = sx/2 + 100
	startY = sy / 2
)

func main() {
	a := app.New()
	w := a.NewWindow("Jiggle")

	stop := make(chan bool)

	var startBtn *widget.Button
	startBtn = widget.NewButton("Start", func() {
		startBtn.Disable()
		robot.Move(startX, startY)

		go func() {
			for {
				select {
				case <-stop:
					return
				default:
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

	go func() {
		hook.Register(hook.KeyDown, []string{"*"}, func(e hook.Event) {
			startBtn.Enable()
			stop <- true
		})

		<-hook.Process(hook.Start())
	}()

	w.Resize(fyne.NewSize(300, 150))
	w.SetContent(startBtn)
	w.ShowAndRun()

	hook.End()
}
