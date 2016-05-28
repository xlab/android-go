package app

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/xlab/android-go/android"
)

var (
	logTag     string
	logTagFunc func() string
	logTagMux  sync.Mutex
)

func SetLogTag(tag string) {
	logTagMux.Lock()
	logTag = tag + "\x00"
	logTagFunc = func() string {
		return logTag
	}
	logTagMux.Unlock()
}

type logger struct{}

func (logger) Write(p []byte) (n int, err error) {
	android.LogWrite(int32(android.LogInfo), logTagFunc(), string(p)+"\x00")
	return len(p), nil
}

func init() {
	SetLogTag("GoApp")

	log.SetOutput(logger{})
	// android logcat includes all of log.LstdFlags
	log.SetFlags(log.Flags() &^ log.LstdFlags)

	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stderr = w
	go lineLog(r, android.LogError)

	r, w, err = os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	go lineLog(r, android.LogInfo)
}

func lineLog(f *os.File, priority android.LogPriority) {
	const logSize = 1024 // matches android/log.h.
	r := bufio.NewReaderSize(f, logSize)
	for {
		line, _, err := r.ReadLine()
		str := string(line)
		if err != nil {
			str += " " + err.Error()
		}
		android.LogWrite(int32(priority), logTagFunc(), str+"\x00")
		if err != nil {
			break
		}
	}
}
