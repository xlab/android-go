package main

import (
	"log"
	"os"
	"runtime"

	"github.com/xlab/android-go/app"
)

func init() {
	app.SetLogTag("android-go/tests")
}

func main() {
	log.Printf("NativeActivity started. Platform: %s %s", runtime.GOOS, runtime.GOARCH)
	nativeWindowEvents := make(chan app.NativeWindowEvent)

	app.Main(func(a app.NativeActivity) {
		a.HandleNativeWindowEvents(nativeWindowEvents)
		a.InitDone()
		for {
			select {
			case event := <-a.LifecycleEvents():
				log.Println("Received lifecycle event:", event.Kind)
				if event.Kind == app.OnCreate {
					go func() {
						runTests(a)
						os.Exit(0)
					}()
				}
			case event := <-nativeWindowEvents:
				log.Println("Received window event:", event.Kind)
				if event.Kind == app.NativeWindowRedrawNeeded {
					a.NativeWindowRedrawDone()
				}
			}
		}
	})
}
