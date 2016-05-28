CPP = $(shell ndk-which cpp)

all: all-egl all-gles all-gles2
	CPP="$(CPP)" cgogen -ccdefs=true android.yml

all-egl:
	cgogen egl.yml

all-gles:
	cgogen gles.yml

all-gles2:
	cgogen gles2.yml

clean: clean-egl clean-gles clean-gles2
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

test:
	cd android && go build

test-egl:
	cd egl && go build

test-gles:
	cd gles && go build

test-gles2:
	cd gles2 && go build
