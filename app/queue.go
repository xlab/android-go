package app

import (
	"fmt"
	"log"

	"github.com/xlab/android-go/android"
)

var SkipInputEvents = func(ev *android.InputEvent) {}

var LogInputEvents = func(ev *android.InputEvent) {
	switch android.InputEventGetType(ev) {
	case android.InputEventTypeKey:
		key := android.KeyEventGetKeyCode(ev)
		log.Printf("key event [%d]", key)
	case android.InputEventTypeMotion:
		str := "motion event "
		fingers := android.MotionEventGetPointerCount(ev)
		for i := uint32(0); i < fingers; i++ {
			x := android.MotionEventGetX(ev, i)
			y := android.MotionEventGetY(ev, i)
			pressure := android.MotionEventGetPressure(ev, i)
			str += fmt.Sprintf("[%.0f; %.0f; %.2f]", x, y, pressure)
		}
		log.Println(str)
	}
}

func HandleInputQueues(queueChan <-chan *android.InputQueue, onProcessed func(),
	evHandler func(ev *android.InputEvent)) {

	looper := android.LooperPrepare(android.LooperPrepareAllowNonCallbacks)
	pending := make(chan *android.InputQueue, 1)
	go func() {
		for queue := range queueChan {
			pending <- queue
			android.LooperWake(looper)
		}
	}()

	var current *android.InputQueue
	for {
		if android.LooperPollAll(-1, nil, nil, nil) == android.LooperPollWake {
			select {
			default:
			case p := <-pending:
				if current != nil {
					handleEvents(current, evHandler)
					android.InputQueueDetachLooper(current)
				}
				current = p
				if current != nil {
					android.InputQueueAttachLooper(current, looper, 0, nil, nil)
				}
				onProcessed()
			}
		}
		if current != nil {
			handleEvents(current, evHandler)
		}
	}
}

func handleEvents(queue *android.InputQueue, evHandler func(ev *android.InputEvent)) {
	var ev *android.InputEvent
	for android.InputQueueGetEvent(queue, &ev) >= 0 {
		if android.InputQueuePreDispatchEvent(queue, ev) != 0 {
			continue
		}
		evHandler(ev)
		android.InputQueueFinishEvent(queue, ev, 0)
	}
}
