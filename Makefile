CPP = $(shell ndk-which cpp)

all: gen-android gen-egl gen-gles gen-gles2 gen-gles3 gen-gles31

gen-android:
	CPP="$(CPP)" cgogen -ccdefs=true android.yml

gen-egl:
	cgogen egl.yml

gen-gles:
	cgogen gles.yml

gen-gles2:
	cgogen gles2.yml

gen-gles3:
	cgogen gles3.yml

gen-gles31:
	cgogen gles31.yml

clean: clean-egl clean-gles clean-gles2 clean-gles3 clean-gles31
	rm -f android/cgo_helpers.go android/cgo_helpers.h android/cgo_helpers.c
	rm -f android/doc.go android/types.go android/const.go
	rm -f android/android.go

clean-egl:
	rm -f egl/cgo_helpers.go egl/cgo_helpers.h egl/cgo_helpers.c
	rm -f egl/doc.go egl/types.go egl/const.go
	rm -f egl/egl.go

clean-gles:
	rm -f gles/cgo_helpers.go gles/cgo_helpers.h gles/cgo_helpers.c
	rm -f gles/doc.go gles/types.go gles/const.go
	rm -f gles/gles.go

clean-gles2:
	rm -f gles2/cgo_helpers.go gles2/cgo_helpers.h gles2/cgo_helpers.c
	rm -f gles2/doc.go gles2/types.go gles2/const.go
	rm -f gles2/gles2.go

clean-gles3:
	rm -f gles3/cgo_helpers.go gles3/cgo_helpers.h gles3/cgo_helpers.c
	rm -f gles3/doc.go gles3/types.go gles3/const.go
	rm -f gles3/gles3.go

clean-gles31:
	rm -f gles31/cgo_helpers.go gles31/cgo_helpers.h gles31/cgo_helpers.c
	rm -f gles31/doc.go gles31/types.go gles31/const.go
	rm -f gles31/gles31.go

test:
	cd android && go build

test-egl:
	cd egl && go build

test-gles:
	cd gles && go build

test-gles2:
	cd gles2 && go build

test-gles3:
	cd gles3 && go build

test-gles31:
	cd gles31 && go build
