package android

import "unsafe"

func (e SensorEvent) Acceleration() (x, y, z float32) {
	ref := e.Ref()
	if ref == nil {
		return
	}
	tuple := *(*[3]float32)(unsafe.Pointer(&ref.anon0))
	x = tuple[0]
	y = tuple[1]
	z = tuple[2]
	return
}
