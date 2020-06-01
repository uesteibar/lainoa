package object

type Nil struct{}

func (n *Nil) Inspect() string  { return "nil" }
func (n *Nil) Type() ObjectType { return NIL_OBJECT }
