package brops

import (
	"bufio"
	"bytes"
	"fmt"
	"git-tool/brops/brdata"
	"git-tool/command"
	"regexp"
	"strings"

	"fyne.io/fyne/v2/data/binding"
	"github.com/rohanthewiz/rerr"
)

type BranchListItem struct {
	BranchName    string
	BranchDetails string
}

const regexBranchDetail = "^(\\*)? *(\\w+) +(\\w+) \\[(.*)\\] +(.*)$"
const lenBranchDetailMatches = 6

func GetCurrentBranches() (data interface{}, err error) {
	bytOut, err := command.ExecCmd("git", "branch", "-vv")
	if err != nil {
		return data, rerr.Wrap(err)
	}

	// uniqBranches := make(map[string]struct{}, 16)
	scnr := bufio.NewScanner(bytes.NewReader(bytOut))

	re, err := regexp.Compile(regexBranchDetail)
	if err != nil {
		return data, rerr.Wrap(err)
	}

	for scnr.Scan() { // each line
		line := strings.TrimSpace(scnr.Text())

		brAttrs := re.FindStringSubmatch(line)
		fmt.Printf("matches->%#v\n", brAttrs) // debug
		if len(brAttrs) != lenBranchDetailMatches {
			return data, rerr.New("Did not get the expected matches")
		}

		br := brdata.Branch{
			Name:       brAttrs[2],
			CommitHash: brAttrs[3],
			Upstream:   brAttrs[4],
			CommitMsg:  brAttrs[5],
		}
		/*		if brAttrs[1] == "*" {
					br.Current = true
				}
		*/
		brdata.BranchBindings = append(brdata.BranchBindings,
			binding.BindStruct(&br))

		// if _, ok := uniqBranches[br]; ok {
		// 	continue
		// }
		//
		// uniqBranches[br] = struct{}{} // track
		// brList = append(brList, BranchListItem{BranchName: br, BranchDetails: arr[1]})
	}
	// fmt.Printf("branches->%q\n", brList) // debug
	return
}
