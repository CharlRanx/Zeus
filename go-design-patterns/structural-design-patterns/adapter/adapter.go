package adapter

type Target interface {
	Request() string
}

type Adaptee struct {
}

func (a *Adaptee) SpecificRequest() string {
	return "Specific Request"
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
