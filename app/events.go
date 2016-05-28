// +build android

package app

// #cgo LDFLAGS: -landroid
//
// #include <android/input.h>
// #include <android/native_activity.h>
// #include <android/native_window.h>
import "C"

import (
	"time"
	"unsafe"

	"github.com/xlab/android-go/android"
)

type LifecycleEvent struct {
	Activity *android.NativeActivity
	Kind     LifecycleEventKind
}

type LifecycleEventKind string

const (
	OnCreate  LifecycleEventKind = "onCreate"
	OnDestroy LifecycleEventKind = "onDestroy"
	OnStart   LifecycleEventKind = "onStart"
	OnStop    LifecycleEventKind = "onStop"
	OnPause   LifecycleEventKind = "onPause"
	OnResume  LifecycleEventKind = "onResume"
)

//export onCreate
func onCreate(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(activity),
		Kind:     OnCreate,
	}
	defaultApp.lifecycleEvents <- event
}

//export onDestroy
func onDestroy(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(activity),
		Kind:     OnDestroy,
	}
	defaultApp.lifecycleEvents <- event
}

//export onStart
func onStart(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(activity),
		Kind:     OnStart,
	}
	defaultApp.lifecycleEvents <- event
}

//export onStop
func onStop(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(activity),
		Kind:     OnStop,
	}
	defaultApp.lifecycleEvents <- event
}

//export onPause
func onPause(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(activity),
		Kind:     OnPause,
	}
	defaultApp.lifecycleEvents <- event
}

//export onResume
func onResume(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(activity),
		Kind:     OnResume,
	}
	defaultApp.lifecycleEvents <- event
}

type SaveStateFunc func(activity *android.NativeActivity, size *uint32) unsafe.Pointer

//export onSaveInstanceState
func onSaveInstanceState(activity *C.ANativeActivity, outSize *C.size_t) unsafe.Pointer {
	defaultApp.initWG.Wait()

	// https://developer.android.com/training/basics/activity-lifecycle/recreating.html
	fn := defaultApp.getSaveInstanceStateFunc()
	if fn == nil {
		return nil
	}
	activityRef := android.NewNativeActivityRef(activity)
	result := fn(activityRef, (*uint32)(outSize))
	return result
}

type WindowFocusEvent struct {
	Activity *android.NativeActivity
	HasFocus bool
}

//export onWindowFocusChanged
func onWindowFocusChanged(activity *C.ANativeActivity, hasFocus int) {
	defaultApp.initWG.Wait()

	out := defaultApp.getWindowFocusEventsOut()
	if out == nil {
		return
	}
	event := WindowFocusEvent{
		Activity: android.NewNativeActivityRef(activity),
		HasFocus: hasFocus > 0,
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		// timed out
	}
}

type NativeWindowEvent struct {
	Activity *android.NativeActivity
	Window   *android.NativeWindow
	Kind     NativeWindowEventKind
}

type NativeWindowEventKind string

const (
	NativeWindowCreated      NativeWindowEventKind = "nativeWindowCreated"
	NativeWindowRedrawNeeded NativeWindowEventKind = "nativeWindowRedrawNeeded"
	NativeWindowDestroyed    NativeWindowEventKind = "nativeWindowDestroyed"
)

//export onNativeWindowCreated
func onNativeWindowCreated(activity *C.ANativeActivity, window *C.ANativeWindow) {
	defaultApp.initWG.Wait()

	out := defaultApp.getNativeWindowEventsOut()
	if out == nil {
		return
	}
	event := NativeWindowEvent{
		Activity: android.NewNativeActivityRef(activity),
		Window:   (*android.NativeWindow)(window),
		Kind:     NativeWindowCreated,
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		// timed out
	}
}

//export onNativeWindowRedrawNeeded
func onNativeWindowRedrawNeeded(activity *C.ANativeActivity, window *C.ANativeWindow) {
	defaultApp.initWG.Wait()

	out := defaultApp.getNativeWindowEventsOut()
	if out == nil {
		return
	}
	event := NativeWindowEvent{
		Activity: android.NewNativeActivityRef(activity),
		Window:   (*android.NativeWindow)(window),
		Kind:     NativeWindowRedrawNeeded,
	}
	select {
	case <-defaultApp.nativeWindowRedrawDone:
	default:
		// skip check
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		return // timed out
	}
	// The drawing window for this native activity needs to be
	// redrawn. To avoid transient artifacts during screen changes
	// (such resizing after rotation), applications should not return
	// from this function until they have finished drawing their
	// window in its current state.
	//
	// Refer to
	// https://developer.android.com/ndk/reference/struct_a_native_activity_callbacks.html
	<-defaultApp.nativeWindowRedrawDone
}

//export onNativeWindowDestroyed
func onNativeWindowDestroyed(activity *C.ANativeActivity, window *C.ANativeWindow) {
	defaultApp.initWG.Wait()

	out := defaultApp.getNativeWindowEventsOut()
	if out == nil {
		return
	}
	event := NativeWindowEvent{
		Activity: android.NewNativeActivityRef(activity),
		Window:   (*android.NativeWindow)(window),
		Kind:     NativeWindowDestroyed,
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		// timed out
	}
}

type InputQueueEvent struct {
	Activity *android.NativeActivity
	Queue    *android.InputQueue
	Kind     InputQueueEventKind
}

type InputQueueEventKind string

const (
	QueueCreated   InputQueueEventKind = "queueCreated"
	QueueDestroyed InputQueueEventKind = "queueDestroyed"
)

//export onInputQueueCreated
func onInputQueueCreated(activity *C.ANativeActivity, queue *C.AInputQueue) {
	defaultApp.initWG.Wait()

	out := defaultApp.getInputQueueEventsOut()
	if out == nil {
		return
	}
	event := InputQueueEvent{
		Activity: android.NewNativeActivityRef(activity),
		Queue:    (*android.InputQueue)(queue),
		Kind:     QueueCreated,
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		return // timed out
	}

	<-defaultApp.inputQueueHandled
}

//export onInputQueueDestroyed
func onInputQueueDestroyed(activity *C.ANativeActivity, queue *C.AInputQueue) {
	defaultApp.initWG.Wait()

	out := defaultApp.getInputQueueEventsOut()
	if out == nil {
		return
	}
	event := InputQueueEvent{
		Activity: android.NewNativeActivityRef(activity),
		Queue:    (*android.InputQueue)(queue),
		Kind:     QueueDestroyed,
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		return // timed out
	}

	<-defaultApp.inputQueueHandled
}

type ContentRectEvent struct {
	Activity *android.NativeActivity
	Rect     *android.Rect
}

//export onContentRectChanged
func onContentRectChanged(activity *C.ANativeActivity, rect *C.ARect) {
	defaultApp.initWG.Wait()

	out := defaultApp.getContentRectEventsOut()
	if out == nil {
		return
	}
	event := ContentRectEvent{
		Activity: android.NewNativeActivityRef(activity),
		Rect:     android.NewRectRef(rect),
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		// timed out
	}
}

type ActivityEvent struct {
	Activity *android.NativeActivity
	Kind     ActivityEventKind
}

type ActivityEventKind string

const (
	OnConfigurationChanged ActivityEventKind = "onConfigurationChanged"
	OnLowMemory            ActivityEventKind = "onLowMemory"
)

//export onConfigurationChanged
func onConfigurationChanged(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	out := defaultApp.getActivityEventsOut()
	if out == nil {
		return
	}
	event := ActivityEvent{
		Activity: android.NewNativeActivityRef(activity),
		Kind:     OnConfigurationChanged,
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		// timed out
	}
}

//export onLowMemory
func onLowMemory(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	out := defaultApp.getActivityEventsOut()
	if out == nil {
		return
	}
	event := ActivityEvent{
		Activity: android.NewNativeActivityRef(activity),
		Kind:     OnLowMemory,
	}
	select {
	case out <- event:
		// dispatched
	case <-time.After(defaultApp.maxDispatchTime):
		// timed out
	}
}
