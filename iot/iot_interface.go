package iot

type IoT interface {
	GetId() int
	GetName() string

	GetCurrentStateFunc1() (string, error)
	GetCurrentStateFunc2() (string, error)
	GetCurrentStateFunc3() (string, error)
	GetCurrentStateFunc4() (string, error)
	GetCurrentStateFunc5() (string, error)
}
