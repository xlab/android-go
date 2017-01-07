package egl

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/xlab/android-go/android"
)

type DisplayHandle struct {
	display Display
	surface Surface
	context Context

	Width  int
	Height int
}

func (d *DisplayHandle) EGLDisplay() Display {
	return d.display
}

func (d *DisplayHandle) EGLSurface() Surface {
	return d.surface
}

func (d *DisplayHandle) EGLContext() Context {
	return d.context
}

func (d *DisplayHandle) Destroy() {
	if d == nil {
		return
	}
	MakeCurrent(d.display, NoSurface, NoSurface, NoContext)
	DestroyContext(d.display, d.context)
	DestroySurface(d.display, d.surface)
	Terminate(d.display)
	if err := Error(); err != nil {
		log.Println("EGL error:", err)
	}
}

// NewDisplayHandle initializes EGL display/surface/context and returns a handle object or error.
// Use expectedConfig to specify the desired EGL config constraints like:
//
//  map[int32]int32{
//  	egl.SurfaceType: egl.WindowBit,
//  	egl.RedSize:   8,
//  	egl.GreenSize: 8,
//  	egl.BlueSize:  8,
//  	egl.AlphaSize: 8,
//  	egl.DepthSize: 24,
//  }
func NewDisplayHandle(window *android.NativeWindow, expectedConfig map[int32]int32) (*DisplayHandle, error) {
	display := GetDisplay(DefaultDisplay)
	if Initialize(display, nil, nil) == False {
		err := fmt.Errorf("eglInitialize failed: %v", Error())
		return nil, err
	}

	var count int32
	GetConfigs(display, nil, 0, &count)
	configs := make([]Config, count)
	GetConfigs(display, configs, count, &count)
	var format int32
	var foundConfig = -1
	var v int32
	if expectedConfig == nil {
		expectedConfig = map[int32]int32{}
	}
	for i, cfg := range configs {
		GetConfigAttrib(display, cfg, SurfaceType, &v)
		if e := expectedConfig[SurfaceType]; e > 0 && v&e != e {
			continue
		}
		GetConfigAttrib(display, cfg, RedSize, &v)
		if e := expectedConfig[RedSize]; e > 0 && v != e {
			continue
		}
		GetConfigAttrib(display, cfg, GreenSize, &v)
		if e := expectedConfig[GreenSize]; e > 0 && v != e {
			continue
		}
		GetConfigAttrib(display, cfg, BlueSize, &v)
		if e := expectedConfig[BlueSize]; e > 0 && v != e {
			continue
		}
		GetConfigAttrib(display, cfg, AlphaSize, &v)
		if e := expectedConfig[AlphaSize]; e > 0 && v != e {
			continue
		}
		GetConfigAttrib(display, cfg, DepthSize, &v)
		if e := expectedConfig[DepthSize]; e > 0 && v != e {
			continue
		}
		// gotcha!
		foundConfig = i
		// NativeVisualId is an attribute of the Config that is
		// guaranteed to be accepted by android.NativeWindowSetBuffersGeometry().
		// As soon as we picked a Config, we can safely reconfigure the
		// NativeWindow buffers to match, using NativeVisualId.
		GetConfigAttrib(display, cfg, NativeVisualId, &format)
	}
	if foundConfig < 0 {
		Terminate(display)
		err := fmt.Errorf("failed to find EGL config for %#v", expectedConfig)
		return nil, err
	}
	android.NativeWindowSetBuffersGeometry(window, 0, 0, format)
	windowPtr := NativeWindowType(unsafe.Pointer(window))
	surface := CreateWindowSurface(display, configs[foundConfig], windowPtr, nil)
	var ctxRequest []int32
	if ctxVer := expectedConfig[ContextClientVersion]; ctxVer > 0 {
		ctxRequest = []int32{ContextClientVersion, ctxVer, None}
	}
	context := CreateContext(display, configs[foundConfig], NoContext, ctxRequest)
	if MakeCurrent(display, surface, surface, context) == False {
		DestroyContext(display, context)
		DestroySurface(display, surface)
		Terminate(display)
		err := fmt.Errorf("eglMakeCurrent failed: %v", Error())
		return nil, err
	}
	handle := &DisplayHandle{
		display: display,
		surface: surface,
		context: context,
	}
	handle.UpdateDimensions()
	return handle, nil
}

func (d *DisplayHandle) UpdateDimensions() {
	var width, height int32
	QuerySurface(d.display, d.surface, Width, &width)
	QuerySurface(d.display, d.surface, Height, &height)
	d.Width = int(width)
	d.Height = int(height)
}

func (d *DisplayHandle) SwapBuffers() {
	SwapBuffers(d.display, d.surface)
}
