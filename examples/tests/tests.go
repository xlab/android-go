package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/xlab/android-go/app"
)

var tests = []test{testAssetsConcurrent, testAssetsShared}

func runTests(a app.NativeActivity) {
	var wg sync.WaitGroup
	for _, t := range tests {
		wg.Add(1)
		go func(t test) {
			t.log("Started")
			t.run(t, a)
			t.log("Completed")
			wg.Done()
		}(t)
	}
	wg.Wait()
	log.Println("Completed all tests")
}

type test struct {
	name string
	run  func(test, app.NativeActivity)
}

func (t test) fail(v ...interface{}) {
	t.log("Failed -", fmt.Sprintln(v...))
	panic(fmt.Sprintln(v...))
}

func (t test) failIf(cond bool, v ...interface{}) {
	if cond {
		t.fail(v...)
	}
}

func (t test) log(v ...interface{}) {
	log.Print(t.name, ": ", fmt.Sprintln(v...))
}
