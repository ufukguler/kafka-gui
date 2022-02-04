// auto-generated

package gui

import (
	"fyne.io/fyne"
	"os"
)

var icon = &fyne.StaticResource{
	StaticName:    "Icon.png",
	StaticContent: getIcon(),
}

func getIcon() []byte {
	dat, err := os.ReadFile("Icon.png")
	if err != nil {
		panic(err)
	}
	return dat
}
