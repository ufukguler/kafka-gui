package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"time"
)

func (c *Dashboard) showPopUp(text string) {
	label := widget.NewLabel(text)
	popUp := widget.NewPopUp(label, c.window.Canvas())
	popUp.ShowAtPosition(fyne.Position{
		X: (c.window.Canvas().Size().Width / 2) - 50,
		Y: c.window.Canvas().Size().Height / 2,
	})
	popUp.Show()
	time.Sleep(time.Second * 1)
	popUp.Hide()
}

func (c *Dashboard) showPopUpDuration(text string, d time.Duration) {
	label := widget.NewLabel(text)
	popUp := widget.NewPopUp(label, c.window.Canvas())
	popUp.ShowAtPosition(fyne.Position{
		X: (c.window.Canvas().Size().Width / 2) - 50,
		Y: c.window.Canvas().Size().Height / 2,
	})
	popUp.Show()
	time.Sleep(d)
	popUp.Hide()
}

func (c *Dashboard) showProgressBar() {
	myWindow := c.app.NewWindow("ProgressBar Widget")
	progress := widget.NewProgressBar()
	go func() {
		for i := 0.0; i <= 1.0; i += 0.3 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
		time.Sleep(time.Second * 1)
		myWindow.Hide()
	}()
	myWindow.Resize(fyne.Size{
		Width:  300,
		Height: 100,
	})
	myWindow.CenterOnScreen()
	myWindow.SetContent(container.NewVBox(progress))
	myWindow.Show()
}

func (c *Dashboard) initInfiniteProgressBar(text string) fyne.Window {
	myWindow := c.app.NewWindow(text)
	progress := widget.NewProgressBarInfinite()
	myWindow.Resize(fyne.Size{
		Width:  350,
		Height: 50,
	})
	myWindow.CenterOnScreen()
	myWindow.SetContent(container.NewVBox(progress))
	return myWindow
}

func (c *Dashboard) showEntryWindow(title string, input *widget.Entry, btnFunc func()) string {
	myWindow := c.app.NewWindow(title)
	myWindow.Resize(fyne.Size{
		Width:  300,
		Height: 100,
	})
	input.SetPlaceHolder("Enter text...")

	content := container.NewVBox(input, widget.NewButton("Save", btnFunc))

	myWindow.SetContent(content)
	myWindow.Show()
	return input.Text
}

func (c *Dashboard) showDialog(title, message string, f func(b bool)) {
	dialog.NewConfirm(title, message, f, c.window).Show()
}

func (c *Dashboard) showTopicList(data []string) {
	myWindow := c.app.NewWindow("Topic List")
	myWindow.CenterOnScreen()
	table := widget.NewTable(func() (int, int) {
		return len(data), 1
	}, func() fyne.CanvasObject {
		return widget.NewLabel("Topics")
	}, func(id widget.TableCellID, object fyne.CanvasObject) {
		object.(*widget.Label).SetText("- " + data[id.Row])
	})
	h := float32(len(data) * 60)
	if h > 500 {
		h = 500
	}
	myWindow.Resize(fyne.Size{
		Width:  400,
		Height: h,
	})
	myWindow.SetContent(container.NewHScroll(table))
	myWindow.Show()
}
