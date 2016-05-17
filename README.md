## Android ABI bindings for Go

Because I can.

### Example

Somewhere in the code of a native Go app for Android:

```
// +build android

import "github.com/android-go/android"


// #include <android/native_activity.h>
// #include <android/native_window.h>
import "C"

//export onNativeWindowCreated
func onNativeWindowCreated(activity *C.ANativeActivity, window *C.ANativeWindow) {
    win := (*android.NativeWindow)(window)
    log.Println("[WHOOP WHOOP!] Window width:", android.NativeWindowGetWidth(win))
}
```

Results in:

```
 2016/05/17 04:41:12 android.go:142: [WHOOP WHOOP!] Window width: 1920
```
