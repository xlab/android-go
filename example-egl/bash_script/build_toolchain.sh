#!/bin/bash
#Working directory is <project root>/android
set -e

ABIS=($(echo $1 | sed 's/,/ /g'))
ANDROID_API=$2
printf "Preparing toolchains for ABIs: %s Android API: %s\n" "${ABIS[*]}" $ANDROID_API


TOOLCHAIN_ROOT_DIR=build_go/toolchain
printf "Cleaning toolchains root directory: %s\n" "$TOOLCHAIN_ROOT_DIR"
rm -rf "$TOOLCHAIN_ROOT_DIR"
mkdir -p "$TOOLCHAIN_ROOT_DIR"

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
printf "Toolchain will be created for arch: %s (ABI: %s)\n" $ARCH $ABI

ARCHS[$ARCH]=1
done

for ARCH in ${!ARCHS[*]}
do

TOOLCHAIN_DIR="$TOOLCHAIN_ROOT_DIR/$ARCH"
printf "Making standalone toolchain for arch: %s\n" $ARCH
set -x
"$ANDROID_HOME"/ndk-bundle/build/tools/make_standalone_toolchain.py \
		--api=$ANDROID_API --install-dir=$TOOLCHAIN_DIR \
		--arch=$ARCH --stl libc++
# The command above includes the wrong headers due to a bug in the android sdk and/or cmake. The following two commands fix this issue.
rm -rf $TOOLCHAIN_DIR/sysroot/usr
cp -r  "$ANDROID_HOME/ndk-bundle/platforms/android-$ANDROID_API/arch-$ARCH/usr" "$TOOLCHAIN_DIR/sysroot/usr"
set +x
printf "Standalone toolchain ready\n"
done