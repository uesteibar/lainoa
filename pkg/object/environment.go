package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	if _, exists := e.Get(name); exists {
		return NewError("can't re-bind already bound identifier `%s`", name)
	}

	e.store[name] = val
	return val
}

func (e *Environment) Rebind(name string, val Object) Object {
	env := e
	_, exists := e.store[name]
	if !exists && e.outer != nil {
		_, exists = e.outer.Get(name)
		if exists {
			env = e.outer
		}
	}

	if !exists {
		return NewError(
			"can't assign identier `%s` because it doesn't exist, you need to do `let %s = %s` first",
			name, name, val.Inspect(),
		)
	}

	env.store[name] = val
	return val
}
