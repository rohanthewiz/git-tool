package main

import (
	"git-tool/brops"
	"git-tool/uibuilder"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	pastBranches := make([]string, 0, 32) // len and capacity specified
	currBranch := ""
	var err error

	wg.Add(1) // checkout a goroutine
	go func() {
		defer wg.Done()
		pastBranches, err = brops.GetPastBranches(pastBranches)
		if err != nil {
			log.Println(err.Error()) // todo use rlog
			return
		}
	}()

	wg.Add(1) // checkout a goroutine
	go func() {
		defer wg.Done()
		currBranch, err = brops.GetCurrentBranch()
		if err != nil {
			log.Println(err.Error()) // todo use rlog
			return
		}
	}()
	wg.Wait()

	w := uibuilder.BuildWindow(currBranch, pastBranches)
	w.ShowAndRun()
}
