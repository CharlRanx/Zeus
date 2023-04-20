package bridge

type Carer interface {
	Ignite() string
}

type Engine interface {
	GetSound() string
}

type Car struct {
	engine Engine
}

func NewCar(engine Engine) Carer {
	return &Car{
		engine,
	}
}

func (c *Car) Ignite() string {
	return c.engine.GetSound()
}

type EngineA struct {
}

func (e *EngineA) GetSound() string {
	return "engine A"
}

type EngineB struct {
}

func (e *EngineB) GetSound() string {
	return "engine B"
}
