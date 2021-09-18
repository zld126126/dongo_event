# dongo event
    A Simple Golang EventDispatcher
- [dongo event](#dongo-event)
  - [1.how to use?](#1how-to-use)
  - [2.todo](#2todo)

## 1.how to use?
```
// Main
func main() {
	dongo_event.EventManager.RegisterEvent("SayHello", EventTest_SayHello)
	dongo_event.EventManagerEventManager.DispatchEvent("SayHello", []int{1, 2, 3})
	dongo_event.EventManagerEventManager.RegisterEvent("SayHello", EventTest_SayHello)
	dongo_event.EventManagerEventManager.RemoveEvent("SayHello", EventTest_SayHello)
	dongo_event.EventManagerEventManager.DispatchEvent("SayHello", []int{1, 3})
	dongo_event.EventManagerEventManager.RegisterEvent("SayHello", EventTest_SayHello)
	dongo_event.EventManagerEventManager.DispatchEvent("SayHello", []int{1, 3})
}

// Custom Event
func EventTest_SayHello(in interface{}) error {
	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(in)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	case reflect.String:
		s := reflect.ValueOf(in)
		fmt.Println(s.String(), "I am a string type variable.")
	case reflect.Int:
		s := reflect.ValueOf(in)
		t := s.Int()
		fmt.Println(t, " I am a int type variable.")
	default:
		return errors.New("unknown support type")
	}
	return nil
}
```

## 2.todo
    SingleTon && Mute