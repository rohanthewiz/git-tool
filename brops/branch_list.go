package brops

import (
	"bufio"
	"bytes"
	"git-tool/command"
	"log"
	"strings"

	"github.com/rohanthewiz/rerr"
)

type BranchListItem struct {
	BranchName    string
	BranchDetails string
}

func GetBranchList() (data interface{}, err error) {
	var brList []BranchListItem

	bytOut, err := command.ExecCmd("git", "branch", "-vv")
	if err != nil {
		return data, rerr.Wrap(err)
	}
	// fmt.Println("Reflog ->", string(bytOut))

	uniqBranches := make(map[string]struct{}, 16)
	scnr := bufio.NewScanner(bytes.NewReader(bytOut))

	for scnr.Scan() { // each line
		line := strings.TrimSpace(scnr.Text())
		arr := strings.SplitN(line, " ", 2)
		if len(arr) != 2 {
			log.Println("This one is weird:", line)
			continue
		}
		br := arr[0]
		if _, ok := uniqBranches[br]; ok {
			continue
		}

		uniqBranches[br] = struct{}{} // track
		brList = append(brList, BranchListItem{BranchName: br, BranchDetails: arr[1]})
	}
	// fmt.Printf("branches->%q\n", brList) // debug
	return brList, err
}
