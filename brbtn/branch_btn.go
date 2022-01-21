package brbtn

import (
	"fmt"
	"git-tool/brops"
	"git-tool/ui_binds"
	"log"

	"fyne.io/fyne/v2/widget"
	"github.com/rohanthewiz/rerr"
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
		brops.CheckoutBranch(abtn.branch)
		brCurr, err := brops.GetCurrentBranch()
		if err != nil {
			log.Println(rerr.StringFromErr(err))
		} else {
			br = brCurr
		}
		err = ui_binds.CurrentBranch.Set(fmt.Sprintf("(current: %s)", br))
		if err != nil {
			log.Println("Error updating current branch in ui", err.Error())
		}
	}
	return abtn
}
