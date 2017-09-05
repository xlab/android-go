Test app
========

This app is used for testing. See `examples/minimal` more information on the build system.

#### Test

Execute `../build-android.sh && (adb uninstall com.go_android.tests; adb install android/app-debug.apk) && adb logcat -c && adb logcat | grep "android-go/tests:"` to build the apk and then install and test on a connected device.