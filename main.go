package main

import "github.com/zserge/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Directory Service",
		"http://connectors.cbwpayments.com:7070/ds/#/user/list", 800, 600, true)
}
