package main

import (
	"log"

	"github.com/wailsapp/wails/lib/renderer/webview"
)

func main() {
	w := webview.NewWebview(webview.Settings{Width: 400, Height: 972})

	w.SetTitle("Minimal webview example")
	file := w.DialogOpenMultiple(0, "*.nes", "", "")
	log.Println(file)
	for _, v := range file {
		log.Println("file ", v)
	}
	w.Run()
}
