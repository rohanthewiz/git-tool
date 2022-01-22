package features

import (
	"git-tool/brops"
	"git-tool/system"
)

func GetBranchesInfo() (pastBranches []string, currBranch string, err error) {
	pastBranches = make([]string, 0, 32) // len and capacity specified

	pastBrsGIO := &system.GoRoutineIO{Fn: brops.GetPastBranches}
	currBrGIO := &system.GoRoutineIO{Fn: brops.GetCurrentBranch}
	system.LaunchGoRoutines(pastBrsGIO, currBrGIO)

	if pastBrsGIO.Err == nil {
		pastBranches = pastBrsGIO.ResultData.([]string)
	}

	if currBrGIO.Err == nil {
		currBranch = currBrGIO.ResultData.(string)
	}
	return
}
