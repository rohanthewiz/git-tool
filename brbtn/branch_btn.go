package brbtn

import (
	"git-tool/features"

	"fyne.io/fyne/v2/widget"
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
		features.CheckoutBranch(abtn.branch)
	}
	return abtn
}
