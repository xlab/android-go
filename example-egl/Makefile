NDK ?= $(ANDROID_HOME)
ANDROID_TOOLCHAIN_DIR ?= $(shell pwd)/toolchain
ANDROID_API ?= 26
ANDROID_SYSROOT = $(NDK)/platforms/android-$(ANDROID_API)/arch-arm

all: toolchain build apk

toolchain:
	$(NDK)/build/tools/make_standalone_toolchain.py \
		--api=$(ANDROID_API) --install-dir=$(ANDROID_TOOLCHAIN_DIR) \
		--arch=arm --stl libc++

build:
	mkdir -p android/app/src/main/jniLibs/armeabi-v7a
	mkdir -p android/app/src/build/go-output/armeabi-v7a
	CC="$(ANDROID_TOOLCHAIN_DIR)/bin/arm-linux-androideabi-gcc" \
	CXX="$(ANDROID_TOOLCHAIN_DIR)/bin/arm-linux-androideabi-g++" \
	CGO_CFLAGS="-march=armv7-a" \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	CGO_ENABLED=1 \
	go build -pkgdir android/app/src/build/go-output/armeabi-v7a \
	-buildmode=c-shared -o android/app/src/main/jniLibs/armeabi-v7a
 
apk:
	cd android && make

clean:
	cd android && make clean

install:
	cd android && make install

listen:
	adb logcat -c
	adb logcat *:S EGLActivity
