EGLActivity
===========

This app leverages all three packages together: **android**, **egl**
and of course **gles** to create quite a rich app that animates its color based
on the accelerometer values. It also reads input events such as key events and
multitouch motion events (with pressure, if supported by the device), you can
check these events in the ADB Shell. Check out the video of the expected
behaviour:

[![Golang + EGL/GLES App on Android](https://img.youtube.com/vi/H2cafzATUEw/0.jpg)](https://www.youtube.com/watch?v=H2cafzATUEw)

### Prerequisites

* Install the latest Android SDK or extract the Android SDK command line tools into your preferred Android SDK root (available as archive at https://developer.android.com/studio/index.html, see bottom of page).

* Make sure you have the `$ANDROID_HOME` environment variable set to the Android SDK root (default is `$HOME/android-sdk`).

[example]: https://github.com/xlab/android-go/tree/master/example

### Building

```
$ ./gradlew build
```

### Deploying

```
$ ./gradlew installDebug
```
Use adb to make sure your phone attached correctly:
```
$ adb devices
```

### Cleanup

To remove standalone toolchains you can execute:
```
$ ./gradlew cleanToolchain
```
Toolchains will be regenerated on the next build.