package main

import (
	"log"
	"runtime"
	"time"

	"github.com/xlab/android-go/android"
)

type SensorMan struct {
	handler func(event *android.SensorEvent)

	sensorManager    *android.SensorManager
	sensorEventQueue *android.SensorEventQueue

	refreshRate   time.Duration
	accelerometer *android.Sensor
	looper        *android.Looper

	startC chan struct{}
	stopC  chan struct{}
	quitC  chan struct{}
	sleep  bool

	quit bool
}

func NewSensorMan(refresh time.Duration,
	handler func(event *android.SensorEvent)) *SensorMan {
	s := &SensorMan{
		handler:     handler,
		refreshRate: refresh,

		startC: make(chan struct{}),
		stopC:  make(chan struct{}),
		quitC:  make(chan struct{}),
	}
	go s.loop()
	return s
}

const queueIdent = 15

func (s *SensorMan) loop() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	s.looper = android.LooperPrepare(android.LooperPrepareAllowNonCallbacks)
	s.sensorManager = android.SensorManagerGetInstance()
	s.accelerometer = android.SensorManagerGetDefaultSensor(s.sensorManager, android.SensorTypeAccelerometer)
	s.sensorEventQueue = android.SensorManagerCreateEventQueue(s.sensorManager,
		s.looper, queueIdent, nil, nil)

	for {
		select {
		case <-s.quitC:
			s.quit = true
			android.SensorEventQueueDisableSensor(s.sensorEventQueue, s.accelerometer)
			android.SensorManagerDestroyEventQueue(s.sensorManager, s.sensorEventQueue)
			return
		case <-s.startC:
			s.sleep = false
			android.SensorEventQueueEnableSensor(s.sensorEventQueue, s.accelerometer)
			android.SensorEventQueueSetEventRate(s.sensorEventQueue, s.accelerometer,
				int32(s.refreshRate/time.Microsecond))
		case <-s.stopC:
			s.sleep = true
			android.SensorEventQueueDisableSensor(s.sensorEventQueue, s.accelerometer)
		default:
			var delay int32 = 10
			if s.sleep {
				delay = 250
			}
			if ident := android.LooperPollAll(delay, nil, nil, nil); ident != queueIdent {
				switch ident {
				case android.LooperPollTimeout:
				default:
					log.Println("illegal poll ident:", ident)
				}
				continue
			}
			var ev android.SensorEvent
			for android.SensorEventQueueGetEvents(s.sensorEventQueue, &ev, 1) > 0 {
				s.handler(&ev)
			}
		}
	}
}

func (s *SensorMan) Start() {
	if s.quit {
		return
	}
	s.startC <- struct{}{}
}

func (s *SensorMan) Stop() {
	if s.quit {
		return
	}
	s.stopC <- struct{}{}
}

func (s *SensorMan) Destroy() {
	if s.quit {
		return
	}
	s.quitC <- struct{}{}
}
