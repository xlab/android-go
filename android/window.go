package android

import "unsafe"

func (w *NativeWindow) Ptr() uintptr {
	return uintptr(unsafe.Pointer(w))
}
