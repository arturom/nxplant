package diagrams

type Writable interface {
}

type Field struct {
	Char string
	Type string
	Name string
}

type Separation struct {
	Val string
}

type Class struct {
	Type    string
	Name    string
	Content []Writable
}

type Relation struct {
	From   string
	To     string
	Label  string
	Symbol string
}

type Diagram struct {
	Name      string
	Classes   []Class
	Relations []Relation
}

func (c *Class) HasFields() bool {
	return len(c.Content) > 0
}
