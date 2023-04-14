package prototype

type Prototype interface {
	Clone() Prototype
	GetName() string
}

type ConcreteProduct struct {
	name string
}

func NewConcreteProduct(name string) Prototype {
	return &ConcreteProduct{
		name: name,
	}
}

func (c *ConcreteProduct) Clone() Prototype {
	return &ConcreteProduct{name: c.name}
}

func (c *ConcreteProduct) GetName() string {
	return c.name
}
