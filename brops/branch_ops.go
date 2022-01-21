package brops

import (
	"bufio"
	"bytes"
	"fmt"
	"git-tool/command"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/rohanthewiz/rerr"
)

func CheckoutBranch(val string) (out string) {
	cmd := exec.Command("git", "checkout", val)
	by, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Checked out branch", val, "Cmdline resp:", string(by))
	out = "checked out branch " + val + "\n" + string(by)
	return
}

func GetPastBranches(branches []string) (brchs []string, err error) {
	cmd := exec.Command("git", "reflog", "-100")

	// Run
	bytOut, err := cmd.Output()
	if err != nil {
		return brchs, err
	}
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return
	}
	fmt.Println("Reflog ->", string(bytOut))

	uniqBranches := make(map[string]struct{}, 16)

	scnr := bufio.NewScanner(bytes.NewReader(bytOut))
	for scnr.Scan() {
		// Parse output
		rg, err := regexp.Compile(`checkout: moving from (.+?) to`)
		if err != nil {
			return brchs, errors.Wrap(err, "failed to compile regex")
		}
		matches := rg.FindStringSubmatch(scnr.Text())
		if len(matches) > 1 {
			// fmt.Printf("matches %#v\n", matches)
			br := matches[1] // the capture group
			// Validations
			// Do we already have this branch in the map
			if _, ok := uniqBranches[br]; ok {
				continue
			}
			// TODO - Skip current and non-existent branches

			uniqBranches[br] = struct{}{} // track
			branches = append(branches, br)
		}
	}

	// fmt.Printf("branches->%q\n", branches) // debug
	return branches, nil
}

func GetCurrentBranch() (br string, err error) {
	byts, err := command.ExecCmd("git", "symbolic-ref", "HEAD", "--short")
	if err != nil {
		return br, rerr.Wrap(err)
	}
	return strings.TrimSpace(string(byts)), nil
}
