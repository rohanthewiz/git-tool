package uibuilder

import (
	"git-tool/brbtn"
	"git-tool/data_bindings"
	"git-tool/features"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/rohanthewiz/rerr"
)

const winInitialWidth, winInitialHeight float32 = 500, 500

func BuildWindow() (fyne.Window, error) {
	myApp := app.New()
	w := myApp.NewWindow("Entry Widget")
	w.Resize(fyne.Size{Width: winInitialWidth, Height: winInitialHeight})

	pastBranches, currBranch, err := features.GetBranchesInfo()
	if err != nil {
		return w, rerr.Wrap(err)
	}

	conBranchTitle, err := buildBranchListTitle(currBranch)
	if err != nil {
		return w, rerr.Wrap(err)
	}

	result := widget.NewLabelWithData(data_bindings.CmdResp)

	var wgtPastBranches []fyne.CanvasObject

	for _, br := range pastBranches {
		btn := brbtn.NewBranchBtn(br.BranchName, br.BranchName)

		brRow := [2]fyne.CanvasObject{
			btn,
			widget.NewLabel(br.BranchDetails),
		}
		wgtPastBranches = append(wgtPastBranches, brRow[:]...)
	}

	brGrid := container.New(layout.NewGridLayout(2), wgtPastBranches...)

	vScroll := container.NewVScroll(brGrid)
	vScroll.SetMinSize(fyne.Size{Width: winInitialWidth, Height: winInitialHeight * 0.75})
	mainVB := container.NewVBox(
		conBranchTitle,
		widget.NewSeparator(),
		vScroll,
		layout.NewSpacer(),
		result,
	)

	w.SetContent(mainVB)
	return w, nil
}
