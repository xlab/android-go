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
printf "Preparing toolchains for ABIs: %s Android API: %s\n" "${ABIS[*]}" $ANDROID_API


TOOLCHAIN_ROOT_DIR=build_go/toolchain

declare -A ARCHS
for ABI in ${ABIS[*]}
do
case $ABI in
armeabi-v7a)
ARCH="arm"
;;
x86)
ARCH="x86"
;;
*)
continue
;;
esac

ARCHS[$ARCH]=1
done

for ARCH in ${!ARCHS[*]}
do
TOOLCHAIN_DIR="$TOOLCHAIN_ROOT_DIR/$ARCH"

if [ -d "$TOOLCHAIN_DIR" ];
then
printf "Using existing standalone toolchain for arch: %s\n" $ARCH
continue
fi

printf "Making standalone toolchain for arch: %s\n" $ARCH
set -x
"$ANDROID_HOME"/ndk-bundle/build/tools/make_standalone_toolchain.py \
		--api=$ANDROID_API --install-dir=$TOOLCHAIN_DIR \
		--arch=$ARCH --stl libc++
set +x
printf "Standalone toolchain ready\n"
done
