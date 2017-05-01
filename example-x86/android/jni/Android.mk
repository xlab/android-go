LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)

LOCAL_MODULE    := example
LOCAL_SRC_FILES := lib/libexample.so
LOCAL_LDLIBS    := -llog -landroid
APP_ABI := x86

include $(PREBUILT_SHARED_LIBRARY)
