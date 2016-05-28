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

There is no additional prerequisites, the project fully inherits the same
structure as the [example] app provides, make sure you were able to run it
first. Just to make sure that everything works smoothly for your OS, environment
setup and the device itself.

[example]: https://github.com/xlab/android-go/tree/master/example

### Structure

```
$ tree .
.
├── Makefile
├── android
│   ├── AndroidManifest.xml
│   ├── Makefile
│   ├── jni
│   │   ├── Android.mk
│   │   └── Application.mk
│   └── res
├── display.go
├── main.go
└── sensor.go

3 directories, 9 files
```

### Running

```
$ make
$ make install
$ make listen
```
