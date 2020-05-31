package object

import (
	"errors"
	"sync"
)

func NewEnvironment() *Environment {
	e := &Environment{scopes: newStack()}
	e.AddScope()

	return e
}

type Environment struct {
	scopes *stack
}

func (e *Environment) Get(name string) (Object, bool) {
	object, exists := e.scopes.GetIdentifier(name)

	return object, exists
}

func (e *Environment) Set(name string, value Object) Object {
	if _, exists := e.Get(name); exists {
		return NewError("can't re-bind already bound identifier `%s`", name)
	}

	scope, err := e.scopes.Current()
	if err != nil {
		panic(err)
	}

	scope.store[name] = value

	return value
}

func (e *Environment) Rebind(name string, value Object) Object {
	scope, exists := e.scopes.GetScopeForIdentifier(name)

	if !exists {
		return NewError(
			"can't assign identier `%s` because it doesn't exist, you need to do `let %s = %s` first",
			name, name, value.Inspect(),
		)
	}

	scope.store[name] = value

	return value
}

func (e *Environment) AddScope() {
	e.scopes.Push(newScope())
}

func (e *Environment) ReleaseScope() {
	e.scopes.Pop()
}

func newScope() *scope {
	return &scope{store: make(map[string]Object)}
}

type scope struct {
	store map[string]Object
}

type stack struct {
	lock   sync.Mutex
	scopes []*scope
}

func newStack() *stack {
	return &stack{scopes: []*scope{}}
}

func (s *stack) GetScopeForIdentifier(name string) (*scope, bool) {
	for _, sc := range s.scopes {
		_, exists := sc.store[name]

		if exists {
			return sc, true
		}
	}

	return nil, false
}

func (s *stack) GetIdentifier(name string) (Object, bool) {
	for _, sc := range s.scopes {
		obj, exists := sc.store[name]

		if exists {
			return obj, true
		}
	}

	return nil, false
}

func (s *stack) Push(sc *scope) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.scopes = append(s.scopes, sc)
}

func (s *stack) Pop() (*scope, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.scopes)
	if l == 0 {
		return nil, errors.New("Empty Stack")
	}

	res := s.scopes[l-1]
	s.scopes = s.scopes[:l-1]
	return res, nil
}

func (s *stack) Current() (*scope, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.scopes)
	if l == 0 {
		return nil, errors.New("Empty Stack")
	}

	return s.scopes[l-1], nil
}
