package features

import (
	"git-tool/brops"
	"git-tool/brops/brdata"
	"git-tool/system"
)

// GetBranchesInfo returns pastBranches in most recent order
func GetBranchesInfo() (validPastBranches []brdata.Branch, currBranch string, err error) {
	validPastBranches = make([]brdata.Branch, 0, 32) // len and capacity specified
	brList := make([]brdata.Branch, 0, 32)
	pastBranches := make([]string, 0, 32) // this, based on reflog, will establish the order

	currBrsGIO := &system.GoRoutineIO{Fn: brops.GetCurrentBranches}
	pastBrsGIO := &system.GoRoutineIO{Fn: brops.GetPastBranches}
	// currBrGIO := &system.GoRoutineIO{Fn: brops.GetCurrentBranch}

	system.LaunchGoRoutines(currBrsGIO, pastBrsGIO) // , currBrGIO fire these in parallel

	if currBrsGIO.Err == nil {
		brList = currBrsGIO.ResultData.([]brdata.Branch)
	}

	if pastBrsGIO.Err == nil {
		pastBranches = pastBrsGIO.ResultData.([]string)
	}

	// Collect only valid past branches into a map
	validBranchInfos := make(map[string]brdata.Branch, 32)

	for _, brItem := range brList {
		// if brItem.BranchName == currBranch {
		// 	println("*skipping current branch")
		// 	continue
		// }
		validBranchInfos[brItem.Name] = brItem
	}

	// Order as in reflog, but use only valid branches
	for _, pbr := range pastBranches {
		if details, ok := validBranchInfos[pbr]; ok {
			validPastBranches = append(validPastBranches, details)
			// brdata.Branch{
			// 	Name: pbr,
			// 	BranchDetails: util.TruncateString(
			// 		strings.TrimSpace(details), 100, true)})
		}
	}

	return
}
