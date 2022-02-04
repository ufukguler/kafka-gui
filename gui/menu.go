package gui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"kafka-gui/kafka"
	"strconv"
	"strings"
	"time"
)

func (c *Dashboard) loadMenu() {
	c.window.SetMainMenu(&fyne.MainMenu{
		Items: []*fyne.Menu{
			c.getMainMenu(),
			c.getConfigMenu(),
			c.getTestMenu(),
		},
	})
}

func (c *Dashboard) getMainMenu() *fyne.Menu {
	ipPort := fyne.NewMenuItem("Save IP:Port", func() {
		ip := strings.TrimSpace(c.ip.Text)
		if ip != "" {
			viperSetConfig("IP", ip)
			c.showPopUp("IP:Port saved for further usage")
			log.Infof("IP:Port [ %s ] saved for further usage", ip)
			return
		}
		c.showPopUp("IP:Port value is empty")
	})

	topic := fyne.NewMenuItem("Save Topic", func() {
		topic := strings.TrimSpace(c.topic.Text)
		if topic != "" {
			viperSetConfig("TOPIC", topic)
			c.showPopUp("Topic saved for further usage")
			log.Infof("Topic [ %s ] saved for further usage", topic)
			return
		}
		c.showPopUp("Topic value is empty")
	})

	headerKey := fyne.NewMenuItem("Save Header Key", func() {
		key := strings.TrimSpace(c.headerKey.Text)
		if key != "" {
			viperSetConfig("HEADER_KEY", key)
			c.showPopUp("Header key saved for further usage")
			log.Infof("Header key [ %s ] saved for further usage", key)
			return
		}
		c.showPopUp("Header key is empty")
	})

	headerValue := fyne.NewMenuItem("Save Header Value", func() {
		value := strings.TrimSpace(c.headerValue.Text)
		if value != "" {
			viperSetConfig("HEADER_VALUE", value)
			c.showPopUp("Header value saved for further usage")
			log.Infof("Header value [ %s ] saved for further usage", value)
			return
		}
		c.showPopUp("Header value is empty")
	})

	partition := fyne.NewMenuItem("Save Partition", func() {
		partition := strings.TrimSpace(c.partition.Text)
		if partition != "" {
			viperSetConfig("PARTITION", partition)
			c.showPopUp("Partition saved for further usage")
			log.Infof("Partition [ %s ] saved for further usage", partition)
			return
		}
		c.showPopUp("Partition value is empty")
	})
	return &fyne.Menu{
		Label: "Save",
		Items: []*fyne.MenuItem{ipPort, topic, headerKey, headerValue, partition},
	}
}

func (c *Dashboard) getConfigMenu() *fyne.Menu {
	width := fyne.NewMenuItem("Change Width", func() {
		myWindow := c.app.NewWindow("Set Width")
		myWindow.CenterOnScreen()
		myWindow.Resize(fyne.Size{
			Width:  300,
			Height: 100,
		})
		input := widget.NewEntry()
		input.SetText(viper.GetString("WIDTH"))
		input.Validator = func(s string) error {
			if strings.TrimSpace(s) == "" {
				return errors.New("Invalid width")
			}

			_, err := strconv.ParseFloat(input.Text, 32)
			if err != nil {
				return err
			}
			return nil
		}
		input.SetPlaceHolder("Enter width...")
		content := container.NewVBox(
			widget.NewLabel("Only Digits"),
			input,
			widget.NewButton("Save", func() {
				if input.Validate() != nil {
					c.showPopUp("invalid width value")
					return
				}
				viperSetConfig("WIDTH", input.Text)
				log.Infof("Width [ %s ] saved for further usage", input.Text)
				h, err := strconv.ParseFloat(input.Text, 32)
				if err == nil {
					c.window.Resize(fyne.Size{
						Width:  float32(h),
						Height: c.window.Canvas().Size().Height,
					})
				}
				myWindow.Close()
			}),
		)
		myWindow.SetContent(content)
		myWindow.Show()
	})

	height := fyne.NewMenuItem("Change Height", func() {
		myWindow := c.app.NewWindow("Set Height")
		myWindow.CenterOnScreen()
		myWindow.Resize(fyne.Size{
			Width:  300,
			Height: 100,
		})
		input := widget.NewEntry()
		input.SetText(viper.GetString("HEIGHT"))
		input.Validator = func(s string) error {
			if strings.TrimSpace(s) == "" {
				return errors.New("invalid height")
			}
			_, err := strconv.ParseFloat(input.Text, 32)
			if err != nil {
				return err
			}
			return nil
		}
		input.SetPlaceHolder("Enter height...")
		content := container.NewVBox(
			widget.NewLabel("Only Digits"),
			input,
			widget.NewButton("Save", func() {
				if input.Validate() != nil {
					c.showPopUp("Invalid height value")
					return
				}
				viperSetConfig("HEIGHT", input.Text)
				log.Infof("Height [ %s ] saved for further usage", input.Text)
				w, err := strconv.ParseFloat(input.Text, 32)
				if err == nil {
					c.window.Resize(fyne.Size{
						Height: float32(w),
						Width:  c.window.Canvas().Size().Width,
					})
				}
				myWindow.Close()
			}),
		)
		myWindow.SetContent(content)
		myWindow.Show()
	})

	return &fyne.Menu{
		Label: "Config",
		Items: []*fyne.MenuItem{width, height},
	}
}

func (c *Dashboard) getTestMenu() *fyne.Menu {
	test := fyne.NewMenuItem("Test Connection", func() {
		ip := strings.TrimSpace(c.ip.Text)
		if ip == "" {
			c.showPopUpDuration("Error: Invalid IP:Port", time.Second*3)
			return
		}
		progressBar := c.initInfiniteProgressBar("Testing Kafka")
		progressBar.CenterOnScreen()
		progressBar.Show()
		time.Sleep(time.Millisecond * 250)
		conn, err := kafka.Dial(ip)
		defer conn.Close()
		if err != nil {
			progressBar.Close()
			c.showPopUpDuration("Error: "+err.Error(), time.Second*3)
			return
		}
		progressBar.Close()
		log.Infof("Kafka [ %s:%d ] connection test was successful.", conn.Broker().Host, conn.Broker().Port)
		c.showPopUp("Test was successful.")
	})
	return &fyne.Menu{
		Label: "Kafka Test",
		Items: []*fyne.MenuItem{test},
	}
}
