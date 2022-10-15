package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
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
	stop := make(chan bool)

	a := app.New()
	w := a.NewWindow("Jiggle")

	status := binding.NewString()
	status.Set("Stopped")
	statusLbl := widget.NewLabelWithData(binding.NewSprintf("Status: %s", status))
	statusLbl.Alignment = fyne.TextAlignCenter
	statusLbl.TextStyle.Bold = true

	instruct := binding.NewString()
	instructLbl := widget.NewLabelWithData(instruct)
	instructLbl.Alignment = fyne.TextAlignCenter

	var startBtn *widget.Button
	startBtn = widget.NewButton("Start", func() {
		startBtn.Disable()
		status.Set("Started")
		instruct.Set("Press any key once to stop")

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

	w.SetContent(container.NewVBox(
		layout.NewSpacer(),
		statusLbl,
		layout.NewSpacer(),
		instructLbl,
		startBtn,
	))

	// Do this after the content setup. For some reason setting the size before
	// setting the content messes up the layout.
	w.Resize(fyne.NewSize(300, 150))
	w.SetFixedSize(true)

	// TODO: Fix multiple key press bug.
	go func() {
		hook.Register(hook.KeyDown, []string{"*"}, func(e hook.Event) {
			stop <- true

			startBtn.Enable()
			status.Set("Stopped")
			instruct.Set("")
		})

		<-hook.Process(hook.Start())
	}()

	w.SetOnClosed(hook.End)
	w.ShowAndRun()
}
