package brdata

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/rohanthewiz/rtable"
)

var BranchBindings []binding.DataMap

type Branch struct {
	Selected bool // Selected in UI
	// Current bool // currently checked out branch
	Name       string
	Upstream   string
	CommitHash string
	CommitMsg  string
}

var BranchDataCols = []rtable.ColAttr{
	{ColName: "Selected", Header: "Select", WidthPercent: 64},
	{ColName: "Name", Header: "Name", WidthPercent: 100},
	{ColName: "CommitHash", Header: "Head", WidthPercent: 100},
	{ColName: "CommitMsg", Header: "CommitMsg", WidthPercent: 100},
	{ColName: "Upstream", Header: "Upstream", WidthPercent: 100},
}
