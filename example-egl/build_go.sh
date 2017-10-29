#!/bin/bash
set -e

BUILD_MODE=$1
ARCHS=($(echo $2 | sed 's/,/ /g'))
ANDROID_API=$3
BUILD_DIR=build_go
printf "Build mode: %s\nArchitectures to build: %s\nAndroid API version %s\n" $BUILD_MODE "${ARCHS[*]}" $ANDROID_API

for ARCH in ${ARCHS[*]}
do
printf "Architecture: %s\n" $ARCH
done

mkdir -p $BUILD_DIR

#$ANDROID_HOME/ndk-bundle/build/tools/make_standalone_toolchain.py \
#		--api=$ANDROID_API --install-dir=$BUILD_DIR/toolchain \
#		--arch=arm --stl libc++
