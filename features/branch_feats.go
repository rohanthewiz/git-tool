package features

import (
	"fmt"
	"git-tool/brops"
	"git-tool/data_bindings"
	"log"

	"github.com/rohanthewiz/rerr"
)

func CheckoutBranch(br string) {
	if br == "" {
		log.Println("no branch specified")
		return
	}
	brops.CheckoutBranch(br)
	brCurr, err := brops.GetCurrentBranch()
	if err != nil {
		log.Println(rerr.StringFromErr(err))
	} else {
		br = brCurr.(string)
	}
	err = data_bindings.CurrentBranch.Set(fmt.Sprintf("(current: %s)", br))
	if err != nil {
		log.Println("Error updating current branch in ui", err.Error())
	}
}
