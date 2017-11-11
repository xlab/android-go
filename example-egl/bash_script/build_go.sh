#!/bin/bash
#Working directory is <project root>/android
set -e

if ! [ -d "$ANDROID_HOME" ];
then
printf "ANDROID_HOME does not point to any directory. Please set ANDROID_HOME variable\n"
exit 1
fi

ABIS=($(echo $1 | sed 's/,/ /g'))
ANDROID_API=$2
printf "Build Go sources using ABIs: %s Android API: %s\n" "${ABIS[*]}" "$ANDROID_API"

TOOLCHAIN_ROOT_DIR=build_go/toolchain
OUTPUT_ROOT_DIR=build_go/output
printf "Cleaning output dir %s\n" "$OUTPUT_ROOT_DIR"
rm -rf "$OUTPUT_ROOT_DIR"

for ABI in ${ABIS[*]}
do
GOARCH=
GOARM=
CC=
CXX=
CGO_CFLAGS=

case $ABI in
armeabi-v7a)
GOARCH="arm"
GOARM=7
CC="$TOOLCHAIN_ROOT_DIR/arm/bin/arm-linux-androideabi-gcc"
CXX="$TOOLCHAIN_ROOT_DIR/arm/bin/arm-linux-androideabi-g++"
CGO_CFLAGS="-march=armv7-a"
;;
x86)
GOARCH="386"
GOARM=
CC="$TOOLCHAIN_ROOT_DIR/x86/bin/i686-linux-android-gcc"
CXX="$TOOLCHAIN_ROOT_DIR/x86/bin/i686-linux-android-g++"
CGO_CFLAGS=
;;
*)
continue
;;
esac


OUTPUT_DIR="$OUTPUT_ROOT_DIR/$ABI"
mkdir -p "$OUTPUT_DIR"

cd ..
set -x
CURRENT_DIR=$(pwd)
CC="$CURRENT_DIR/android/$CC" CXX="$CURRENT_DIR/android/$CXX" CGO_ENABLED=1 \
CGO_CFLAGS="$CGO_CFLAGS" GOOS=android GOARCH="$GOARCH" GOARM="$GOARM" \
go build -i -pkgdir "$CURRENT_DIR/android/$OUTPUT_DIR" -buildmode=c-shared -o "android/src/main/jniLibs/$ABI/libgomain.so"
cd android
set +x
done