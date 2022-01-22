package main

import (
	"git-tool/uibuilder"
	"log"

	"github.com/rohanthewiz/rerr"
)

func main() {
	w, err := uibuilder.BuildWindow()
	if err != nil {
		log.Println(rerr.StringFromErr(err))
		return
	}
	w.ShowAndRun()
}
