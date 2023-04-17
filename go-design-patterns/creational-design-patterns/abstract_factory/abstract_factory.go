package abstractfactory

// AbstractFactory provides an interface for creating families of related objects.
type AbstractFactory interface {
	CreateWater(volume float64) AbstractWater
	CreateBottle(volume float64) AbstractBottle
}

// AbstractWater provides a water interface.
type AbstractWater interface {
	GetVolume() float64
}

// AbstractBottle provides a bottle interface.
type AbstractBottle interface {
	PourWater(water AbstractWater)
	GetBottleVolume() float64
	GetWaterVolume() float64
}

// ColaFactory implements AbstractFactory interface.
type ColaFactory struct{}

// NewColaFactory is the ColaFactory constructor.
func NewColaFactory() AbstractFactory {
	return &ColaFactory{}
}

// CreateWater implementation.
func (cf *ColaFactory) CreateWater(volume float64) AbstractWater {
	return &ColaWater{volume: volume}

}

// CreateBottle implementation.
func (cf *ColaFactory) CreateBottle(volume float64) AbstractBottle {
	return &ColaBottle{volume: volume}
}

// ColaWater implements AbstractWater interface.
type ColaWater struct {
	volume float64
}

func (w *ColaWater) GetVolume() float64 {
	return w.volume
}

// ColaBottle implements AbstractBottle interface.
type ColaBottle struct {
	water  AbstractWater
	volume float64
}

func (b *ColaBottle) PourWater(water AbstractWater) {
	b.water = water
}

func (b *ColaBottle) GetBottleVolume() float64 {
	return b.volume
}

func (b *ColaBottle) GetWaterVolume() float64 {
	return b.water.GetVolume()
}
