package eventmethod

import (
	"sync"
)

type EventFunc func(map[string]string, map[string]interface{}) error

var EventMethod = map[string]EventFunc{}
var eventMethodLock = new(sync.RWMutex)

func RegisterEventMethod(eventType string, f EventFunc) {
	eventMethodLock.Lock()
	EventMethod[eventType] = f
	eventMethodLock.Unlock()
}

func GetEventMethod(eventType string) EventFunc {
	eventMethodLock.RLock()
	defer eventMethodLock.RUnlock()
	if method, ok := EventMethod[eventType]; ok {
		return method
	}
	return nil
}
