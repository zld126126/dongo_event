package dongo_event

import (
	"fmt"
	"reflect"
)

// TFunction
type _Event struct {
	functions []*TFunction
}

type TFunction func(args interface{}) error

// Call
func (p TFunction) Call(args interface{}) error {
	return p(args)
}

// Add
func (p *_Event) Add(c *TFunction) {
	t := reflect.TypeOf(c)
	isExist := false
	for _, h := range p.functions {
		temp := reflect.TypeOf(h)
		if reflect.DeepEqual(t, temp) {
			isExist = true
		}
	}
	if !isExist {
		p.functions = append(p.functions, c)
	}
}

// Remove
func (p *_Event) Remove(c *TFunction) {
	t := reflect.TypeOf(c)
	result := []*TFunction{}
	for _, h := range p.functions {
		temp := reflect.TypeOf(h)
		if reflect.DeepEqual(t, temp) {
			continue
		} else {
			result = append(result, h)
		}
	}
	p.functions = result
}

// Exec
func (p *_Event) Exec(args ...interface{}) error {
	for _, h := range p.functions {
		err := h.Call(args)
		if err != nil {
			return err
		}
	}
	return nil
}

// EventManager
type _EventManager struct {
	//eventMap
	eventMap map[string]*_Event
}

// getKeys
func (p *_EventManager) getKeys() []string {
	keys := []string{}
	for k := range p.eventMap {
		keys = append(keys, k)
	}
	return keys
}

// DispatchEvent
func (p *_EventManager) DispatchEvent(tp string, args interface{}) {
	if f, ok := p.eventMap[tp]; ok {
		err := f.Exec(args)
		if err != nil {
			fmt.Println(fmt.Sprintf("%+v", err.Error()))
		}
	}
}

// RegisterEvent
func (p *_EventManager) RegisterEvent(tp string, f TFunction) {
	e, ok := p.eventMap[tp]
	if ok {
		e.Add(&f)
	} else {
		tEvent := &_Event{
			functions: []*TFunction{&f},
		}
		p.eventMap[tp] = tEvent
	}
}

// RemoveEvent
func (p *_EventManager) RemoveEvent(tp string, f TFunction) {
	if e, ok := p.eventMap[tp]; ok {
		e.Remove(&f)
		return
	}
}

var EventManager = &_EventManager{eventMap: map[string]*_Event{}}
