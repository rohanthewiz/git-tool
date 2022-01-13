package main

import (
	"bufio"
	"bytes"
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
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
	    pastBranches, err = goGetPastBranches(pastBranches)
	    if err != nil {
	    	log.Println(err.Error())
	    	return
		}
	}()

	// Entry
	lbl := widget.NewLabel("Go back to branch")

	wg.Wait()
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

	w.SetContent(con)
	w.ShowAndRun()
}

func goGetPastBranches(branches []string) (brchs []string, err error) {
	cmd := exec.Command("git", "reflog", "-30")

	// Run
	bytOut, err := cmd.Output()
	if err != nil {
		return brchs, err
	}
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return
	}
	fmt.Println("->", string(bytOut))

	scnr := bufio.NewScanner(bytes.NewReader(bytOut))
	for scnr.Scan() {
		// Parse output
		rg, err := regexp.Compile(`checkout: moving from (.+?) to`)
		if err != nil {
			return brchs, errors.Wrap(err, "failed to compile regex")
		}
		matches := rg.FindStringSubmatch(scnr.Text())
		if len(matches) > 1 {
			// fmt.Printf("matches %#v\n", matches)
			branches = append(branches, matches[1])
		}
	}

	fmt.Printf("branches->%#v\n", branches)
	return branches, nil
}
