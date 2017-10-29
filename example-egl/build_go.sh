#!/bin/bash
#Working directory is <project root>/android
set -ex

ABIS=($(echo $1 | sed 's/,/ /g'))
ANDROID_API=$2
printf "Build Go sources using ABIs: %s Android API: %s Mode: %s\n" "${ABIS[*]}" $ANDROID_API "nomode"

TOOLCHAIN_ROOT_DIR=build_go/toolchain
OUTPUT_ROOT_DIR=build_go/output
printf "Cleaning output dir %s\n" "$OUTPUT_ROOT_DIR"
rm -rf "$OUTPUT_ROOT_DIR"

for ABI in ${ABIS[*]}
do

TOOLCHAIN_DIR="$TOOLCHAIN_ROOT_DIR/$ABI"

case $ABI in
armeabi-v7a)
ARCH="arm"
GOARCH=arm
GOARM=7
CC="$TOOLCHAIN_DIR/bin/arm-linux-androideabi-gcc"
CXX="$TOOLCHAIN_DIR/bin/arm-linux-androideabi-g++"
CGO_CFLAGS="-march=armv7-a"
;;
x86)
ARCH="x86"
;;
esac


OUTPUT_DIR=$OUTPUT_ROOT_DIR/$ABI
mkdir -p $OUTPUT_DIR

cd ..
CC=$(pwd)/android/$CC CXX=$(pwd)/android/$CXX CGO_ENABLED=1 CGO_CFLAGS=$CGO_CFLAGS GOOS=android GOARCH=$GOARCH GOARM=$GOARM \
go build -i -pkgdir $(pwd)/android/$OUTPUT_DIR -buildmode=c-shared -o android/src/main/jniLibs/$ABI/libgomain.so
cd android

done