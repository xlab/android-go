package android

type SoftKeyboardState int32

const (
	SoftKeyboardHidden SoftKeyboardState = iota
	SoftKeyboardVisible
)

func SetSoftKeyboardState(state SoftKeyboardState) {

}

func (a *NativeActivity) JNICall(fn func()) {
	a.Deref()
	// JavaVMAttachArgs
}
