package uibuilder

import (
	"fmt"
	"git-tool/brbtn"
	"git-tool/ui_binds"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const winInitialWidth, winInitialHeight float32 = 500, 500

func BuildWindow(currBranch string, pastBranches []string) fyne.Window {
	myApp := app.New()
	w := myApp.NewWindow("Entry Widget")
	w.Resize(fyne.Size{Width: winInitialWidth, Height: winInitialHeight})

	lbl := widget.NewLabel("Local Branches ")

	err := ui_binds.CurrentBranch.Set(fmt.Sprintf("(current: %s)", currBranch))
	if err != nil {
		log.Println("Error updating current branch in ui", err.Error())
	}
	lbl2 := widget.NewLabelWithData(ui_binds.CurrentBranch)

	result := widget.NewMultiLineEntry()

	var wgtPastBranches []fyne.CanvasObject

	for _, br := range pastBranches {
		btn := brbtn.NewBranchBtn(" > ", br)

		brRow := [3]fyne.CanvasObject{
			btn,
			widget.NewLabel(br),
			widget.NewLabel("Branch info will go here"),
		}
		wgtPastBranches = append(wgtPastBranches, brRow[:]...)
	}

	brGrid := container.New(layout.NewGridLayout(3), wgtPastBranches...)

	vscon := container.NewVScroll(brGrid)
	vscon.SetMinSize(fyne.Size{Width: winInitialWidth, Height: winInitialHeight * 0.75})
	mcon := container.NewVBox(
		lbl, lbl2,
		widget.NewSeparator(),
		vscon,
		layout.NewSpacer(),
		result,
	)

	w.SetContent(mcon)
	return w
}
