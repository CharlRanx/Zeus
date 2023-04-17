package builder

// Builder provides a builder interface.
type Builder interface {
	MakeHead(str string)
	MakeBody(str string)
	MakeFoot(str string)
}

// Director implements a manager.
type Director struct {
	builder Builder
}

// Construct tells the builder what to do and in what order.
func (d *Director) Construct() {
	d.builder.MakeHead("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFoot("Footer")
}

// ConcreteBuilder implements a Builder interface.
type ConcreteBuilder struct {
	product *Product
}

func (cb *ConcreteBuilder) MakeHead(str string) {
	cb.product.Content += "<header>" + str + "</header>"
}

func (cb *ConcreteBuilder) MakeBody(str string) {
	cb.product.Content += "<article>" + str + "</article>"
}

func (cb *ConcreteBuilder) MakeFoot(str string) {
	cb.product.Content += "<footer>" + str + "</footer>"
}

type Product struct {
	Content string
}

func (p *Product) Show() string {
	return p.Content
}
