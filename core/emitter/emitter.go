package emitter

import (
	"fmt"
	"sync"
)

type Emitter struct {
	listeners map[string][]func(interface{})
	mutex     sync.RWMutex
}

func New() *Emitter {
	return &Emitter{
		listeners: make(map[string][]func(interface{})),
	}
}

func (e *Emitter) On(event string, listener func(interface{})) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.listeners[event] = append(e.listeners[event], listener)
}

func (e *Emitter) Emit(event string, data interface{}) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	// Use a WaitGroup to wait for all listeners to finish
	var wg sync.WaitGroup
	for _, listener := range e.listeners[event] {
		wg.Add(1)
		go func(listener func(interface{})) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovered from panic in listener for event %s: %v\n", event, r)
				}
			}()
			listener(data)
		}(listener)
	}
	wg.Wait() // Block until all listeners complete
}

func (e *Emitter) Clear() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.listeners = make(map[string][]func(interface{}))
}
