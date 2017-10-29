#!/bin/bash
set -e

ABIS=($(echo $1 | sed 's/,/ /g'))
ANDROID_API=$2
printf "Architectures to build: %s\nAndroid API version %s\n" "${ABIS[*]}" $ANDROID_API


TOOLCHAIN_ROOT_DIR=build_go/toolchain
mkdir -p $TOOLCHAIN_ROOT_DIR
printf "Toolchain root directory: %s\n" $TOOLCHAIN_ROOT_DIR
for ABI in ${ABIS[*]}
do

case $ABI in
armeabi-v7a)
ARCH="arm"
;;
x86)
ARCH="x86"
;;
esac
printf "Architecture: %s Application binary interface: %s\n" $ABI $ARCH

TOOLCHAIN_DIR="$TOOLCHAIN_ROOT_DIR/$ABI"
printf "Toolchain dir: %s\n" "$TOOLCHAIN_DIR"
rm -rf "$TOOLCHAIN_DIR"

set -x
$ANDROID_HOME/ndk-bundle/build/tools/make_standalone_toolchain.py \
		--api=$ANDROID_API --install-dir=$TOOLCHAIN_DIR \
		--arch=$ARCH --stl libc++
set +x
done