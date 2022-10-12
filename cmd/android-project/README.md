## android-project

Android Project tool is a simple replacement for infamous `android` util from Android SDK, prior to Android SDK Tools [Revision 25.3.0 (March 2017) release](https://developer.android.com/studio/releases/sdk-tools.html) when they dropped that util abruptly. Realising their mistake later that March, they got it back, so it actually forwards the commands for `avd`, `target`, and `device` to the underlying layers properly, but still lacks of `android project update` capability.

I used `android project update` to simply create Ant build scripts for the projects with guarantees, that platform exists and is supported. So this tool does the same thing and it's the only feature implemented now.

### Installation

```
go install github.com/xlab/android-go/cmd/android-project@latest
```

Also you must set `$ANDROID_HOME` to your Android SDK location, e.g.

```
export ANDROID_HOME=/Users/xlab/Library/Android/sdk
```

### Usage

```
Usage: android-project update [--sdk] [--target] [--force] --name --path

Updates an Android project (must already have an AndroidManifest.xml)
```

[GolangExample](https://github.com/xlab/android-go/blob/master/example/android/Makefile)

```
$ android-project update --target android-23 --name GolangExample --path .
.
├── Environment
│   ├── [/Users/xlab/Library/Android/sdk]  Android SDK location
│   ├── [/Users/xlab/Library/Android/sdk/ndk-bundle]  Android NDK location
│   ├── [/Users/xlab/Library/Android/sdk/platforms]  Android Platforms location
│   └── [[android-23 android-N]]  Android Platforms available
├── Project
│   ├── [GolangExample]  Project name
│   ├── [android-23]  Project target
│   └── [.]  Project location
└── Files updated
    ├── build.xml
    ├── local.properties
    ├── project.properties
    └── proguard-project.txt
```

As you can see, the effect of running this command is very similar to `android project update`. At least is used to be.

```
$ android project update
*************************************************************************
The "android" command is deprecated.
For manual SDK, AVD, and project management, please use Android Studio.
For command-line tools, use tools/bin/sdkmanager and tools/bin/avdmanager
*************************************************************************
Invalid or unsupported command "project update"
```

### License

MIT
