package main

import (
	"io"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	binContentChan := make(chan []byte)

	a := app.New()
	w := a.NewWindow("File Browser")
	w.Resize(fyne.NewSize(500, 500))

	fileLabel := widget.NewLabel("File:")
	selectorButton := widget.NewButton("Select File", func() {
		selector := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil || uc == nil {
				return
			}
			content, err := io.ReadAll(uc)
			if err != nil {
				log.Println("读取文件失败:", err)
			}
			binContentChan <- content
		}, w)
		selector.Show()
	})

	fileWidget := widget.NewMultiLineEntry()
	fileWidget.SetMinRowsVisible(10)

	content := container.NewHBox(
		fileLabel,
		container.NewGridWrap(
			fyne.NewSize(80, 30),
			selectorButton,
		),
	)
	content = container.NewVBox(content, container.NewMax(fileWidget))
	w.SetContent(content)

	go func() {
		content := <-binContentChan
		fileWidget.SetText(string(content))
	}()

	// w.SetFixedSize(true)
	w.ShowAndRun()
}
