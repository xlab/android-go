// +build android

#include <android/log.h>
#include <dlfcn.h>
#include <fcntl.h>
#include <stdint.h>
#include <string.h>
#include "_cgo_export.h"

#define LOG_FATAL(...) __android_log_print(ANDROID_LOG_FATAL, "Go", __VA_ARGS__)

int main_running = 0;

// ANativeActivity_onCreate implements the application's entry point.
//
// This is the function that must be in the native code to instantiate
// the application's native activity. It is called with the activity
// instance; if the code is being instantiated from a
// previously saved instance, the savedState will be non-NULL and
// point to the saved data. You must make any copy of this data you
// need – it will be released after you return from this function.
//
// See https://developer.android.com/ndk/reference/group___native_activity.html
//
// The Activity may be created and destroyed multiple times throughout
// the life of a single process. Each time, onCreate is called.
void ANativeActivity_onCreate(ANativeActivity *activity, void* savedState, size_t savedStateSize) {
	if (!main_running) {
		// Call the Go main.main.
		uintptr_t mainPC = (uintptr_t)dlsym(RTLD_DEFAULT, "main.main");
		if (!mainPC) {
			LOG_FATAL("missing main.main");
		}
		callMain(mainPC);
		main_running = 1;
	}

	// Set the native activity callbacks, see:
	// https://developer.android.com/ndk/reference/struct_a_native_activity_callbacks.html
	activity->callbacks->onStart = onStart;
	activity->callbacks->onResume = onResume;
	activity->callbacks->onSaveInstanceState = onSaveInstanceState;
	activity->callbacks->onPause = onPause;
	activity->callbacks->onStop = onStop;
	activity->callbacks->onDestroy = onDestroy;
	activity->callbacks->onWindowFocusChanged = onWindowFocusChanged;
	activity->callbacks->onNativeWindowCreated = onNativeWindowCreated;
	activity->callbacks->onNativeWindowRedrawNeeded = onNativeWindowRedrawNeeded;
	activity->callbacks->onNativeWindowDestroyed = onNativeWindowDestroyed;
	activity->callbacks->onInputQueueCreated = onInputQueueCreated;
	activity->callbacks->onInputQueueDestroyed = onInputQueueDestroyed;
	activity->callbacks->onConfigurationChanged = onConfigurationChanged;
	activity->callbacks->onLowMemory = onLowMemory;

	onCreate(activity);
}
