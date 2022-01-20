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

type BranchBtn struct {
	branch string
	widget.Button
}

func NewBranchBtn(text, br string) *BranchBtn {
	abtn := &BranchBtn{branch: br}
	abtn.SetText(text)
	abtn.ExtendBaseWidget(abtn)
	abtn.OnTapped = func() {
		checkoutBranch(abtn.branch)
	}
	return abtn
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Entry Widget")
	w.Resize(fyne.Size{Width: 500, Height: 500})

	wg := sync.WaitGroup{}
	pastBranches := make([]string, 0, 16)
	var err error

	wg.Add(1) // checkout a goroutine
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
	result := widget.NewMultiLineEntry()

	wg.Wait()

	var wgtPastBranches []fyne.CanvasObject

	for _, br := range pastBranches {
		btn := NewBranchBtn("...", br)

		brRow := [3]fyne.CanvasObject{
			btn,
			widget.NewLabel(br),
			widget.NewLabel("Branch info will go here"),
		}
		wgtPastBranches = append(wgtPastBranches, brRow[:]...)
	}

	brGrid := container.New(layout.NewGridLayout(3), wgtPastBranches...)

	con := container.NewVBox(
		lbl, widget.NewSeparator(),
		brGrid,
		layout.NewSpacer(), result,
		widget.NewSeparator())

	w.SetContent(con)
	w.ShowAndRun()
}

func checkoutBranch(val string) (out string) {
	cmd := exec.Command("git", "checkout", val)
	by, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Checked out branch", val, "Cmdline resp:", string(by))
	out = "checked out branch " + val + "\n" + string(by)
	return
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
	fmt.Println("Reflog ->", string(bytOut))

	uniqBranches := make(map[string]struct{}, 16)

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
			br := matches[1] // the capture group
			// Validations
			// Do we already have this branch in the map
			if _, ok := uniqBranches[br]; ok {
				continue
			}
			// TODO - Skip current and non-existent branches

			uniqBranches[br] = struct{}{} // track
			branches = append(branches, br)
		}
	}

	// fmt.Printf("branches->%q\n", branches) // debug
	return branches, nil
}
