package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/segmentio/kafka-go/protocol"
	log "github.com/sirupsen/logrus"
	"kafka-gui/kafka"
	"strings"
	"time"
)

func produceBtnFunc(c *Dashboard) {
	header := protocol.Header{}
	if c.headerKey.Text != "" {
		header.Key = strings.TrimSpace(c.headerKey.Text)
	}
	if c.headerValue.Text != "" {
		header.Value = []byte(strings.TrimSpace(c.headerValue.Text))
	}

	topic := strings.TrimSpace(c.topic.Text)
	partition := strings.TrimSpace(c.partition.Text)
	ip := strings.TrimSpace(c.ip.Text)
	data := strings.TrimSpace(c.data.Text)
	if err := kafka.Produce(topic, partition, ip, data, header); err != nil {
		c.showPopUpDuration(err.Error(), time.Second*3)
		return
	}
	log.Infof("Message produced successfully. Topic: %s, Partition: %s, IP: %s, Header:[%v]:[%v]",
		topic, partition, ip, c.headerKey.Text, c.headerValue.Text)
	c.showPopUp("Successful.")
}

func topicCreateBtnFunc(c *Dashboard) {
	myWindow := c.app.NewWindow("Create Topic")
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.Size{
		Width:  300,
		Height: 100,
	})
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter topic name...")
	content := container.NewVBox(
		widget.NewLabel("KAFKA_AUTO_CREATE_TOPICS_ENABLE must be set 'true'"),
		input,
		widget.NewButton("Create", func() {
			progressBar := c.initInfiniteProgressBar("Creating Topic")
			progressBar.CenterOnScreen()
			progressBar.Show()
			time.Sleep(time.Millisecond * 250)
			conn, err := kafka.CreateTopic(input.Text, c.partition.Text, c.ip.Text)
			defer conn.Close()
			if err != nil {
				progressBar.Close()
				c.showPopUpDuration(err.Error(), time.Second*3)
				return
			}
			progressBar.Hide()
			myWindow.Close()
			time.Sleep(time.Millisecond * 250)
			log.Infof("Topic { %s } created successfully.", input.Text)
			c.showPopUp("Topic created successfully.")
		}),
	)
	myWindow.SetContent(content)
	myWindow.Show()
}

func topicListBtnFunc(c *Dashboard) {
	topics, err := kafka.GetTopics(c.ip.Text)
	if err != nil {
		c.showPopUpDuration(err.Error(), time.Second*3)
		return
	}
	c.showTopicList(topics)
}
