package features

import (
	"git-tool/brops"
	"git-tool/system"
	"git-tool/system/util"
	"strings"
)

// GetBranchesInfo returns pastBranches in most recent order
func GetBranchesInfo() (validPastBranches []brops.BranchListItem, currBranch string, err error) {
	validPastBranches = make([]brops.BranchListItem, 0, 32) // len and capacity specified
	brList := make([]brops.BranchListItem, 0, 32)
	pastBranches := make([]string, 0, 32) // this, based on reflog, will establish the order

	brListGIO := &system.GoRoutineIO{Fn: brops.GetBranchList}
	pastBrsGIO := &system.GoRoutineIO{Fn: brops.GetPastBranches}
	currBrGIO := &system.GoRoutineIO{Fn: brops.GetCurrentBranch}

	system.LaunchGoRoutines(brListGIO, pastBrsGIO, currBrGIO) // fire these in parallel

	if brListGIO.Err == nil {
		brList = brListGIO.ResultData.([]brops.BranchListItem)
	}

	if pastBrsGIO.Err == nil {
		pastBranches = pastBrsGIO.ResultData.([]string)
	}

	if currBrGIO.Err == nil {
		currBranch = currBrGIO.ResultData.(string)
	}

	// Collect only valid past branches into a map
	validBranchInfos := make(map[string]string, 32)

	for _, brItem := range brList {
		if brItem.BranchName == currBranch {
			println("*skipping current branch")
			continue
		}
		validBranchInfos[brItem.BranchName] = brItem.BranchDetails
	}

	// Order as in reflog, but use only valid branches
	for _, pbr := range pastBranches {
		if details, ok := validBranchInfos[pbr]; ok {
			validPastBranches = append(validPastBranches,
				brops.BranchListItem{
					BranchName: pbr,
					BranchDetails: util.TruncateString(
						strings.TrimSpace(details), 100, true)})
		}
	}

	return
}
