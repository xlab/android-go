package main

import (
	"log"
	"runtime"
	"time"

	"github.com/xlab/android-go/android"
	"github.com/xlab/android-go/app"
	"github.com/xlab/android-go/egl"
	gl "github.com/xlab/android-go/gles"
)

func init() {
	app.SetLogTag("EGLActivity")
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	nativeWindowEvents := make(chan app.NativeWindowEvent, 1)
	windowFocusEvents := make(chan app.WindowFocusEvent, 1)
	inputQueueEvents := make(chan app.InputQueueEvent, 1)
	inputQueueChan := make(chan *android.InputQueue, 1)
	var displayHandle *DisplayHandle
	var windowFocused bool

	type vec3 struct {
		X, Y, Z float32
	}
	sensorEvents := make(chan vec3, 100)
	sensorMan := NewSensorMan(20*time.Millisecond, func(event *android.SensorEvent) {
		x, y, z := event.Acceleration()
		select {
		case sensorEvents <- vec3{x, y, z}:
		default:
		}
	})

	stateX := float32(0.5)
	stateY := float32(0.5)
	stateZ := float32(0.5)
	stateRiseY := true

	app.Main(func(a app.NativeActivity) {
		a.HandleNativeWindowEvents(nativeWindowEvents)
		a.HandleWindowFocusEvents(windowFocusEvents)
		a.HandleInputQueueEvents(inputQueueEvents)
		go app.HandleInputQueues(inputQueueChan, func() {
			a.InputQueueHandled()
		}, app.LogInputEvents)
		a.InitDone()
		for {
			select {
			case vec := <-sensorEvents:
				// log.Printf("accelerometer x=%0.3f y=%0.3f z=%0.3f", vec.X, vec.Y, vec.Z)
				stateX = 0.5 + vec.X/10.0
				stateZ = 0.5 + vec.Y/10.0
				if stateRiseY {
					stateY += 0.01
					if stateY >= 1 {
						stateRiseY = false
					}
				} else {
					stateY -= 0.01
					if stateY <= 0 {
						stateRiseY = true
					}
				}
				draw(displayHandle, stateX, stateY, stateZ)
			case <-a.LifecycleEvents():
			case event := <-windowFocusEvents:
				if event.HasFocus && !windowFocused {
					windowFocused = true
					sensorMan.Start()
				}
				if !event.HasFocus && windowFocused {
					windowFocused = false
					sensorMan.Stop()
				}
				draw(displayHandle, stateX, stateY, stateZ)
			case event := <-inputQueueEvents:
				switch event.Kind {
				case app.QueueCreated:
					inputQueueChan <- event.Queue
				case app.QueueDestroyed:
					inputQueueChan <- nil
				}
			case event := <-nativeWindowEvents:
				switch event.Kind {
				case app.NativeWindowRedrawNeeded:
					draw(displayHandle, stateX, stateY, stateZ)
					a.NativeWindowRedrawDone()
				case app.NativeWindowCreated:
					if handle, err := NewDisplayHandle(event.Activity, event.Window); err != nil {
						log.Fatalln("EGL error:", err)
					} else {
						displayHandle = handle
						log.Printf("EGL display res: %dx%d", handle.Width, handle.Height)
					}
					initGL()
				case app.NativeWindowDestroyed:
					displayHandle.Destroy()
				}
			}
		}
	})
}

func initGL() {
	gl.Hint(gl.PerspectiveCorrectionHint, gl.Fastest)
	gl.Enable(gl.CullFace)
	gl.Disable(gl.DepthTest)
	gl.SetShadeModel(gl.Smooth)
}

func draw(handle *DisplayHandle, x, y, z float32) {
	gl.ClearColor(x, y, z, 1)
	gl.ClearFunc(gl.ColorBufferBit)
	egl.SwapBuffers(handle.display, handle.surface)
}
