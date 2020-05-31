package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	object, exists := e.store[name]

	return object, exists
}

func (e *Environment) Set(name string, value Object) Object {
	e.store[name] = value

	return value
}
