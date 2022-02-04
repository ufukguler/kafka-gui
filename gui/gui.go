package gui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

type Dashboard struct {
	app            fyne.App
	window         fyne.Window
	ip             *widget.Entry
	topic          *widget.Entry
	partition      *widget.Entry
	headerKey      *widget.Entry
	headerValue    *widget.Entry
	data           *widget.Entry
	button         *widget.Button
	topicListBtn   *widget.Button
	topicCreateBtn *widget.Button
}

func Run() {
	c := &Dashboard{}
	c.app = app.New()
	c.app.Settings().SetTheme(theme.DefaultTheme())
	c.app.SetIcon(icon)
	c.loadUI()
	c.app.Run()
}

func (c *Dashboard) loadUI() {
	log.Info("Loading UI")
	c.window = c.app.NewWindow("Kafka Producer")
	c.window.SetPadded(true)
	c.window.CenterOnScreen()
	c.setUiSize()
	c.initializeWidgets()
	c.setLayoutContent()
	c.setFieldsOnLoad()
	c.window.Show()
}

func (c *Dashboard) setLayoutContent() {
	top := widget.NewForm(
		widget.NewFormItem("IP:Port", c.ip),
		widget.NewFormItem("Topic", c.topic),
		widget.NewFormItem("Header Key", c.headerKey),
		widget.NewFormItem("Header Value", c.headerValue),
		widget.NewFormItem("Partition", c.partition),
	)
	c.loadMenu()
	form := container.NewGridWrap(
		fyne.Size{
			Width:  600,
			Height: 100,
		},
		top,
	)
	form.Move(fyne.Position{
		X: 0,
		Y: 0,
	})

	label := widget.NewLabel("Data")
	label.TextStyle.Bold = true
	dataLabel := container.NewWithoutLayout(label)
	dataLabel.Move(fyne.Position{
		X: c.ip.Position().X - 60,
		Y: 200,
	})

	data := container.NewWithoutLayout(c.data)
	c.data.Resize(fyne.Size{
		Width:  c.ip.Size().Width,
		Height: 500,
	})
	data.Move(fyne.Position{
		X: c.ip.Position().X,
		Y: 200,
	})

	button := container.NewWithoutLayout(c.button)
	c.button.Resize(fyne.Size{
		Width:  c.ip.Size().Width,
		Height: 30,
	})
	button.Move(fyne.Position{
		X: c.ip.Position().X,
		Y: 710,
	})

	topicListBtn := container.NewWithoutLayout(c.topicListBtn)
	c.topicListBtn.Resize(fyne.Size{
		Width:  c.ip.Size().Width / 2,
		Height: 30,
	})
	topicListBtn.Move(fyne.Position{
		X: c.ip.Position().X + c.ip.Size().Width + 50,
		Y: 10,
	})

	topicCreateBtn := container.NewWithoutLayout(c.topicCreateBtn)
	c.topicCreateBtn.Resize(fyne.Size{
		Width:  c.ip.Size().Width / 2,
		Height: 30,
	})
	topicCreateBtn.Move(fyne.Position{
		X: c.ip.Position().X + c.ip.Size().Width + 50,
		Y: 50,
	})

	layout := container.NewWithoutLayout(form, dataLabel, data, button, topicListBtn, topicCreateBtn)
	c.window.SetContent(layout)
	c.window.SetMaster()
}

func (c *Dashboard) setUiSize() {
	width, err := strconv.ParseFloat(viper.GetString("WIDTH"), 32)
	if err != nil {
		width = 1366
	}
	height, err := strconv.ParseFloat(viper.GetString("HEIGHT"), 32)
	if err != nil {
		height = 800
	}

	log.Infof("Setting window size to [ %dx%d ]", int(width), int(height))
	size := fyne.Size{
		Width:  float32(width),
		Height: float32(height),
	}
	c.window.Resize(size)
}

func (c *Dashboard) initializeWidgets() {
	c.ip = widget.NewEntry()
	c.topic = widget.NewEntry()

	c.headerKey = widget.NewEntry()
	c.headerKey.SetPlaceHolder("optional")

	c.headerValue = widget.NewEntry()
	c.headerValue.SetPlaceHolder("optional")

	c.data = widget.NewMultiLineEntry()
	c.partition = widget.NewEntry()
	c.partition.Validator = func(s string) error {
		_, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return errors.New("not a number")
		}
		return nil
	}

	c.button = widget.NewButton("Produce", func() { produceBtnFunc(c) })
	c.topicListBtn = widget.NewButton("List Topics", func() { topicListBtnFunc(c) })
	c.topicCreateBtn = widget.NewButton("Create New Topic", func() { topicCreateBtnFunc(c) })
}

func (c *Dashboard) setFieldsOnLoad() {
	log.Info("Reloading saved fields")
	ip := viper.GetString("IP")
	if ip != "" {
		c.ip.SetText(ip)
	}
	topic := viper.GetString("TOPIC")
	if topic != "" {
		c.topic.SetText(topic)
	}
	key := viper.GetString("HEADER_KEY")
	if key != "" {
		c.headerKey.SetText(key)
	}
	value := viper.GetString("HEADER_VALUE")
	if value != "" {
		c.headerValue.SetText(value)
	}
	partition := viper.GetString("PARTITION")
	if partition != "" {
		c.partition.SetText(partition)
	}
}
