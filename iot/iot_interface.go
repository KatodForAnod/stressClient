package iot

type IoT interface {
	GetId() int
	GetName() string
	StartObserve() error
	StopObserve() error
	IsObserve() bool
	GetCurrentStateFunc1() (string, error)
	GetCurrentStateFunc2() (string, error)
	GetCurrentStateFunc3() (string, error)
	GetCurrentStateFunc4() (string, error)
	GetCurrentStateFunc5() (string, error)
}
