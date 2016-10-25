package main

import (
	"gae-test/view"

	"google.golang.org/appengine"
)

func main() {
	view.Main()
	appengine.Main()
}
