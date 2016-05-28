ANDROID_TOOLCHAIN_DIR ?= $(shell pwd)/toolchain
ANDROID_PLATFORM ?= android-21
ANDROID_SYSROOT = $(NDK)/platforms/$(ANDROID_PLATFORM)/arch-arm

all: toolchain build apk

toolchain:
	$(NDK)/build/tools/make-standalone-toolchain.sh \
		--platform=$(ANDROID_PLATFORM) --install-dir=$(ANDROID_TOOLCHAIN_DIR) \
		--arch=arm

build:
	mkdir -p android/jni/lib
	CC="$(ANDROID_TOOLCHAIN_DIR)/bin/arm-linux-androideabi-gcc" \
	CXX="$(ANDROID_TOOLCHAIN_DIR)/bin/arm-linux-androideabi-g++" \
	CGO_CFLAGS="-march=armv7-a" \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
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
	adb logcat *:S EGLActivity
	