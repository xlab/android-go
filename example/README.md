GolangExample
=============

This is a simple Android Go app template.

![android go example gif](http://cl.ly/3n2Z0M0a3A1T/golangapp.gif)

### Prerequisites

* Install the latest Android SDK (UPD March 2017: supported SDK Tools versions prior to 25.3.0 as they [dropped android tool support](https://developer.android.com/studio/releases/sdk-tools.html) in that release).

* Install the desired Android NDK platforms.

* Install the [Apache Ant](http://ant.apache.org) tool (`$ brew install ant` on OS X).

* Make sure you have the `$NDK` environment variable set to the NDK root, an example root may look like:

```
$ ls $NDK

CHANGELOG.md       ndk-depends*       ndk-which*         prebuilt/          source.properties
build/             ndk-gdb*           package.xml        python-packages/   sources/
ndk-build*         ndk-stack*         platforms/         shader-tools/      toolchains/
```

* Make sure you have this dir added to the `$PATH` so `ndk-which` and `ndk-build` are available by their names.

* Make sure you have added the `/sdk/tools` dir to your `$PATH` too, so the `android` tool is available by its name.

* Make sure you have `adb` available by its name and it connects to the device with no issues.

Now you're done with initial setup.

### Structure

```
$ tree
.
├── Makefile
├── android
│   ├── AndroidManifest.xml
│   ├── Makefile
│   └── jni
│       ├── Android.mk
│       └── Application.mk
└── main.go

2 directories, 6 files
```

#### Makefile

Makefile specifies the top-level build steps for the app. You can just type
`make` or execute steps separately (recommended for the first use). Run `make
toolchain` to build a platform-oriented toolchain using the NDK scripts. The
process is described on
https://developer.android.com/ndk/guides/standalone_toolchain.html

Now you're ready to build the `main` package into a shared library, execute
`make build`. The main package must include the `app` package so it will be
compiled with all the glue needed.

All other steps are just proxies into the `android` project dir.

#### Android project

Directory `android` is a standard NDK project that has the `AndroidManifest.xml`
as well as build scripts that treat the shared library we just built as a
`PREBUILT_SHARED_LIBRARY`.

Change to the directory and run `make project` to bootstrap a full project that
can later be used by your standard ways of Android development. Finally, run
`make build` or `make install` to invoke the `ndk-build` process so it will take
care of the shared library and place it in the right places, then `ant` will
build an APK ready for deployment.

Feel free to observe the underlying commands in Makefiles, they are simple and
straightforward, don't forget to change versions so they match your platforms.

So, back into the root dir of `example`, whenever you want to re-build the
project, just run `make && make install` and it will re-do all the building-
repacking steps needed to get your new code delivered. The errors that Go will
yield are quite readable and maintainable. If you're experiencing something weird,
just report an issue.

Ah, and the ADB Shell shortcut is `make listen`.

### Running

```
$ make
$ make install
$ make listen
```

### Clean-up

`make clean` would clean all the mess created by the bootstrapping process, but
the generated `toolchain` will be left inact, feel free to delete it too if you
want.

