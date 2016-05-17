all:
	CPP=arm-linux-androideabi-cpp cgogen -ccdefs=true android.yml

clean:
	rm -f android/cgo_helpers.go android/cgo_helpers.h android/cgo_helpers.c
	rm -f android/doc.go android/types.go android/const.go
	rm -f android/android.go

test:
	cd android && go build
