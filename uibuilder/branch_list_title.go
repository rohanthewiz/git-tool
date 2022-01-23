package uibuilder

import (
	"fmt"
	"git-tool/data_bindings"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rohanthewiz/rerr"
)

func buildBranchListTitle(currBranch string) (*fyne.Container, error) {
	lbl := widget.NewLabel("Local Branches ")

	err := data_bindings.CurrentBranch.Set(fmt.Sprintf("current: %s", currBranch))
	if err != nil {
		log.Println("Error setting data-binding for current br", err.Error())
		return nil, rerr.Wrap(err)
	}

	lbl2 := widget.NewLabelWithData(data_bindings.CurrentBranch)

	conBranchTitle := container.NewHBox(lbl, lbl2)
	return conBranchTitle, nil
}
