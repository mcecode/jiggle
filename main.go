package main

import (
	"crypto/rand"
	"math/big"
	"sync/atomic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	robot "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var sx, sy = robot.GetScreenSize()

func main() {
	// 0 means false and 1 means true.
	var stop uint32 = 1

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

		atomic.StoreUint32(&stop, 0)

		go func() {
			for {
				if atomic.LoadUint32(&stop) == 1 {
					status.Set("Stopped")
					startBtn.Enable()
					return
				}

				// TODO: Add error handling.
				x, _ := rand.Int(rand.Reader, big.NewInt(int64(sx)))
				y, _ := rand.Int(rand.Reader, big.NewInt(int64(sy)))

				robot.MoveSmooth(int(x.Int64()), int(y.Int64()))
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

	go func() {
		hook.Register(hook.KeyDown, []string{"*"}, func(e hook.Event) {
			if atomic.LoadUint32(&stop) == 0 {
				atomic.StoreUint32(&stop, 1)
				status.Set("Stopping")
				instruct.Set("")
			}
		})

		<-hook.Process(hook.Start())
	}()

	w.SetOnClosed(hook.End)
	w.ShowAndRun()
}
