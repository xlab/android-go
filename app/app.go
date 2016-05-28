// Package app implements a NativeActivity glue layer required to
// properly handle the startup process and the native activity events.
//
// Import this package into your Go application to make it Android-compatible.
package app

// #cgo LDFLAGS: -llog
//
// #include <stdlib.h>
// #include <time.h>
import "C"

import (
	"os"
	"sync"
	"time"
	"unsafe"

	"github.com/xlab/android-go/app/internal/callfn"
)

//export callMain
func callMain(mainPC uintptr) {
	// This function as well as the CallFn trampoline is borrowed from the
	// gomobile project.

	// N.B: the main.main must present in the actual app
	// that imports this package.

	for _, name := range []string{"TMPDIR", "PATH", "LD_LIBRARY_PATH"} {
		n := C.CString(name)
		os.Setenv(name, C.GoString(C.getenv(n)))
		C.free(unsafe.Pointer(n))
	}

	// Set timezone.
	//
	// Note that Android zoneinfo is stored in /system/usr/share/zoneinfo,
	// but it is in some kind of packed TZiff file that we do not support
	// yet. As a stopgap, we build a fixed zone using the tm_zone name.
	var curtime C.time_t
	var curtm C.struct_tm
	C.time(&curtime)
	C.localtime_r(&curtime, &curtm)
	tzOffset := int(curtm.tm_gmtoff)
	tz := C.GoString(curtm.tm_zone)
	time.Local = time.FixedZone(tz, tzOffset)

	go callfn.CallFn(mainPC)
}

type NativeActivity interface {
	InitDone()
	NativeWindowRedrawDone()
	InputQueueHandled()
	LifecycleEvents() <-chan LifecycleEvent
	HandleSaveInstanceState(fn SaveStateFunc)
	HandleWindowFocusEvents(out chan<- WindowFocusEvent)
	HandleNativeWindowEvents(out chan<- NativeWindowEvent)
	HandleInputQueueEvents(out chan<- InputQueueEvent)
	HandleContentRectEvents(out chan<- ContentRectEvent)
	HandleActivityEvents(out chan<- ActivityEvent)
}

var defaultApp = &nativeActivity{
	lifecycleEvents:        make(chan LifecycleEvent),
	nativeWindowRedrawDone: make(chan Signal, 1),
	inputQueueHandled:      make(chan Signal, 1),
	maxDispatchTime:        1 * time.Second,

	initWG: new(sync.WaitGroup),
	mux:    new(sync.RWMutex),
}

func init() {
	defaultApp.initWG.Add(1)
}

func Main(fn func(app NativeActivity)) {
	// runtime.LockOSThread()
	// defer runtime.UnlockOSThread()
	fn(defaultApp)
}

type nativeActivity struct {
	// lifecycleEvents must be handled in real-time.
	lifecycleEvents chan LifecycleEvent

	// maxDispatchTime sets the maximum time the send operation
	// allowed to wait while channel is blocked.
	maxDispatchTime time.Duration
	// channels below are optional and will be sent to only
	// if handled by an external client.

	windowFocusEvents  chan<- WindowFocusEvent
	nativeWindowEvents chan<- NativeWindowEvent
	inputQueueEvents   chan<- InputQueueEvent
	contentRectEvents  chan<- ContentRectEvent
	activityEvents     chan<- ActivityEvent

	saveInstanceStateFunc  SaveStateFunc
	nativeWindowRedrawDone chan Signal
	inputQueueHandled      chan Signal

	initWG *sync.WaitGroup
	mux    *sync.RWMutex
}

func (a *nativeActivity) InitDone() {
	a.initWG.Done()
}

func (a *nativeActivity) LifecycleEvents() <-chan LifecycleEvent {
	return a.lifecycleEvents
}

type Signal struct{}

// NativeWindowRedrawDone should be invoked as soon as WindowRedrawNeeded event has
// been processed and the redraw has been completed.
//
// Refer to https://developer.android.com/ndk/reference/struct_a_native_activity_callbacks.html
func (a *nativeActivity) NativeWindowRedrawDone() {
	select {
	case a.nativeWindowRedrawDone <- Signal{}:
	default:
	}
}

func (a *nativeActivity) InputQueueHandled() {
	select {
	case a.inputQueueHandled <- Signal{}:
	default:
	}
}

func (a *nativeActivity) HandleWindowFocusEvents(out chan<- WindowFocusEvent) {
	a.mux.Lock()
	a.windowFocusEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getWindowFocusEventsOut() chan<- WindowFocusEvent {
	a.mux.RLock()
	out := a.windowFocusEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleNativeWindowEvents(out chan<- NativeWindowEvent) {
	a.mux.Lock()
	a.nativeWindowEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getNativeWindowEventsOut() chan<- NativeWindowEvent {
	a.mux.RLock()
	out := a.nativeWindowEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleInputQueueEvents(out chan<- InputQueueEvent) {
	a.mux.Lock()
	a.inputQueueEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getInputQueueEventsOut() chan<- InputQueueEvent {
	a.mux.RLock()
	out := a.inputQueueEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleContentRectEvents(out chan<- ContentRectEvent) {
	a.mux.Lock()
	a.contentRectEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getContentRectEventsOut() chan<- ContentRectEvent {
	a.mux.RLock()
	out := a.contentRectEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleActivityEvents(out chan<- ActivityEvent) {
	a.mux.Lock()
	a.activityEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getActivityEventsOut() chan<- ActivityEvent {
	a.mux.RLock()
	out := a.activityEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleSaveInstanceState(fn SaveStateFunc) {
	a.mux.Lock()
	a.saveInstanceStateFunc = fn
	a.mux.Unlock()
}

func (a *nativeActivity) getSaveInstanceStateFunc() SaveStateFunc {
	a.mux.RLock()
	fn := a.saveInstanceStateFunc
	a.mux.RUnlock()
	return fn
}
