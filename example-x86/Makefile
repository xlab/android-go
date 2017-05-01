ANDROID_TOOLCHAIN_DIR ?= $(shell pwd)/toolchain
ANDROID_API ?= 23
ANDROID_SYSROOT = $(NDK)/platforms/android-$(ANDROID_API)/arch-x86

all: toolchain build apk

toolchain:
	$(NDK)/build/tools/make_standalone_toolchain.py \
		--api=$(ANDROID_API) --install-dir=$(ANDROID_TOOLCHAIN_DIR) \
		--arch=x86 --stl libc++

build:
	mkdir -p android/jni/lib
	CC="$(ANDROID_TOOLCHAIN_DIR)/bin/i686-linux-android-gcc" \
	CXX="$(ANDROID_TOOLCHAIN_DIR)/bin/i686-linux-android-g++" \
	CGO_CFLAGS="" \
	GOARCH=386 \
	GOOS=android \
	CGO_ENABLED=1 \
	go build -buildmode=c-shared -o android/jni/lib/libexample.so
 
apk:
	cd android && make

clean:
	cd android && make clean

install:
	cd android && make install

listen:
	adb logcat -c
	adb logcat *:S GolangExample
