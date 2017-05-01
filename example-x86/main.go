package main

import (
	"log"

	"github.com/xlab/android-go/app"
)

func init() {
	app.SetLogTag("GolangExample")
}

func main() {
	log.Println("NativeActivity has started ^_^")
	nativeWindowEvents := make(chan app.NativeWindowEvent)

	app.Main(func(a app.NativeActivity) {
		a.HandleNativeWindowEvents(nativeWindowEvents)
		a.InitDone()
		for {
			select {
			case event := <-a.LifecycleEvents():
				switch event.Kind {
				case app.OnCreate:
					log.Println(event.Kind, "handled")
				default:
					log.Println(event.Kind, "event ignored")
				}
			case event := <-nativeWindowEvents:
				switch event.Kind {
				case app.NativeWindowRedrawNeeded:
					a.NativeWindowRedrawDone()
					log.Println(event.Kind, "handled")
				default:
					log.Println(event.Kind, "event ignored")
				}
			}
		}
	})
}
