package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/pkg/errors"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Entry Widget")
	w.Resize(fyne.Size{Width: 500, Height: 500})

	wg := sync.WaitGroup{}
	pastBranches := make([]string, 0, 16)

	wg.Add(1)
	go func() {
		defer wg.Done()
	    _, er := goGetPastBranches(pastBranches)
	    if er != nil {
	    	log.Println(er.Error())
	    	return
		}
	}()

	// Entry
	lbl := widget.NewLabel("Go back to branch")

	// in := widget.NewEntry() // Text input
	// in.SetPlaceHolder("Enter two decimal numbers...")
	sel := widget.NewSelect(pastBranches,
		func(val string) {
			println(val, "selected")
	})

	result := widget.NewEntry()

	// saveBtn := widget.NewButton("Save",
	// 	func() {
	// 		// Parse entry
	// 		var a, b float64
	// 		_, err := fmt.Sscanf(in.Text, "%f %f", &a, &b)
	// 		if err != nil {
	// 			log.Println("Error reading the input", err.Error())
	// 			return
	// 		}
	// 		// Return the result
	// 		sum := a + b
	// 		result.SetText("> " + fmt.Sprintf("%0.1f", sum))
	// 	})

	con := container.NewVBox(
		lbl, widget.NewSeparator(),
		sel,
		layout.NewSpacer(),
		widget.NewSeparator(), result)

	wg.Wait()
	w.SetContent(con)
	w.ShowAndRun()
}

func goGetPastBranches(branches []string) (brchs []string, err error) {
	cmd := exec.Command("git", "reflog", "-30")
	err = cmd.Start()
	if err != nil {
		err = errors.Wrap(err, "error executing git command")
		return
	}

	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)

	// Parse output
	bytOut, err := cmd.Output()
	if err != nil {
		return brchs, err
	}
	rg, err := regexp.Compile(`checkout: moving from (.+?) to`)
	if err != nil {
		return brchs, errors.Wrap(err, "failed to compile regex")
	}
	matches := rg.FindStringSubmatch(string(bytOut))
	if len(matches) > 0 {
		fmt.Println(matches)
	}
	return
}
