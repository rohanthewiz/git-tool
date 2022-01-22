package system

import (
	"log"
	"sync"

	"github.com/rohanthewiz/rerr"
)

type GoRoutineIO struct {
	Fn         func() (interface{}, error)
	ResultData interface{}
	Err        error
}

// LaunchGoRoutines launches the given functions concurrently in goroutines
func LaunchGoRoutines(ios ...*GoRoutineIO) {
	var wg sync.WaitGroup

	for i := range ios {
		wg.Add(1) // track a goroutine
		go func(ii int) {
			defer wg.Done()
			ios[ii].ResultData, ios[ii].Err = ios[ii].Fn()
			if ios[ii].Err != nil {
				ios[ii].Err = rerr.Wrap(ios[ii].Err)
				log.Println(rerr.StringFromErr(ios[ii].Err))
			}
		}(i)
	}

	wg.Wait()
	return
}
