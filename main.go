package main

import (
	customWidget "SSD-Go/widget"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"image/color"
)

func buildServerList(size int) fyne.CanvasObject {
	list := make([]fyne.CanvasObject, 0)
	itemHeight := 64
	for i := 0; i < size; i++ {
		selectionIndicator := customWidget.NewColoredBox()
		selectionIndicator.Resize(fyne.Size{Width: 8, Height: itemHeight})
		selectionIndicator.SetBackgroundColor(red)
		listItemContent := customWidget.NewTwoLineListItem("profile.Name", "SSOMONDFOJNDOFUENW")
		list = append(list, fyne.NewContainerWithLayout(customWidget.NewStackLayout(), selectionIndicator, listItemContent))
	}
	return widget.NewScrollContainer(widget.NewVBox(list...))
}

var red = color.RGBA{R: 255, G: 0, B: 0, A: 255}

func main() {
	application := app.New()

	window := application.NewWindow("SSD Go")
	window.Resize(fyne.Size{
		Width:  1200,
		Height: 800,
	})
	window.CenterOnScreen()

	// At least 1 tab is required, or index out of range is thrown
	subscriptionsTabs := widget.NewTabContainer(
		widget.NewTabItem("First", buildServerList(1)),
		widget.NewTabItem("Second", buildServerList(160)))
	subscriptionsTabs.SetTabLocation(widget.TabLocationLeading)
	window.SetContent(subscriptionsTabs)

	window.ShowAndRun()
}
