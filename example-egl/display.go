package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/xlab/android-go/android"
	"github.com/xlab/android-go/egl"
)

type DisplayHandle struct {
	display egl.Display
	surface egl.Surface
	context egl.Context

	Width  int
	Height int
}

func (d *DisplayHandle) Destroy() {
	if d == nil {
		return
	}
	egl.MakeCurrent(d.display, egl.NoSurface, egl.NoSurface, egl.NoContext)
	egl.DestroyContext(d.display, d.context)
	egl.DestroySurface(d.display, d.surface)
	egl.Terminate(d.display)
	if err := egl.Error(); err != nil {
		log.Println("EGL error:", err)
	}
}

// NewDisplay initializes EGL display/surface/context and returns a handle object or error.
func NewDisplayHandle(activity *android.NativeActivity,
	window *android.NativeWindow) (*DisplayHandle, error) {

	// Here specify the attributes of the desired configuration.
	// Below, we select an egl.Config with at least 8 bits per color
	// component compatible with on-screen windows
	attribs := []int32{
		egl.SurfaceType, egl.WindowBit,
		egl.BlueSize, 8,
		egl.GreenSize, 8,
		egl.RedSize, 8,
		egl.None,
	}
	display := egl.GetDisplay(egl.DefaultDisplay)
	if egl.Initialize(display, nil, nil) == egl.False {
		err := fmt.Errorf("eglInitialize failed: %v", egl.Error())
		return nil, err
	}

	// Here, the application chooses the configuration it desires. In this
	// sample, we have a very simplified selection process, where we pick
	// the first egl.Config that matches our criteria
	configs := make([]egl.Config, 1)
	var numConfigs int32
	if egl.ChooseConfig(display, attribs, configs, 1, &numConfigs) == egl.False {
		egl.Terminate(display)
		err := fmt.Errorf("eglChooseConfig failed: %v", egl.Error())
		return nil, err
	}

	// egl.NativeVisualId is an attribute of the egl.Config that is
	// guaranteed to be accepted by android.NativeWindowSetBuffersGeometry().
	// As soon as we picked a egl.Config, we can safely reconfigure the
	// NativeWindow buffers to match, using egl.NativeVisualId.
	var format int32
	if egl.GetConfigAttrib(display, configs[0], egl.NativeVisualId, &format) == egl.False {
		egl.Terminate(display)
		err := fmt.Errorf("eglGetConfigAttrib failed: %v", egl.Error())
		return nil, err
	}

	android.NativeWindowSetBuffersGeometry(window, 0, 0, int32(format))
	windowPtr := egl.NativeWindowType(unsafe.Pointer(window))
	surface := egl.CreateWindowSurface(display, configs[0], windowPtr, nil)
	context := egl.CreateContext(display, configs[0], egl.NoContext, nil)
	if egl.MakeCurrent(display, surface, surface, context) == egl.False {
		egl.DestroyContext(display, context)
		egl.DestroySurface(display, surface)
		egl.Terminate(display)
		err := fmt.Errorf("eglMakeCurrent failed: %v", egl.Error())
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
	egl.QuerySurface(d.display, d.surface, egl.Width, &width)
	egl.QuerySurface(d.display, d.surface, egl.Height, &height)
	d.Width = int(width)
	d.Height = int(height)
}
