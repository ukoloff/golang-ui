package main

import (
	"github.com/zserge/webview"
)

func main() {
	webview.Open("Minimal webview example",
		"https://en.m.wikipedia.org/wiki/Main_Page", 800, 600, true)
}
