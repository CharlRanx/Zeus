package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	colaFactory := NewColaFactory()
	colaWater := colaFactory.CreateWater(2.5)
	colaBottle := colaFactory.CreateBottle(2.5)

	colaBottle.PourWater(colaWater)

	if colaBottle.GetWaterVolume() != colaBottle.GetBottleVolume() {
		t.Errorf("Expect volume to %.1fL, but %.1fL", colaBottle.GetWaterVolume(), colaBottle.GetBottleVolume())
	}
}
