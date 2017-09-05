Minimal Go app
==============

This is a simple Android Go app template.

### Prerequisites

* Install the latest Android SDK or extract the Android SDK command line tools into your preferred Android SDK root (available as archive at https://developer.android.com/studio/index.html, see bottom of page).

* Make sure you have the `$ANDROID_HOME` environment variable set to the Android SDK root (default is `$HOME/android-sdk`).

#### Android project

Directory `android` is a standard NDK project that has the `AndroidManifest.xml`
as well gradle build files and the shared library we will built as a `PREBUILT_SHARED_LIBRARY`.

### Build

Change to the `examples/minimal` directory and run
`ANDROID_HOME=path/to/sdk ../build-android.sh"`
to build an APK ready for deployment.

#### Test

Execute `../build-android.sh && (adb uninstall com.go_android.minimal; adb install android/app-debug.apk) && adb logcat -c && adb logcat | grep "GolangExample:"` to build the apk and then install and test on a connected device.

### Clean-up

Delete the gradle build folders `android/build`, `android/app/build` and `android/app/src/jniLibs`, which contains the compiled shared libraries.
