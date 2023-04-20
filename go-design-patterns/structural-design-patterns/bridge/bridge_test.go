package bridge

import "testing"

func TestBridge(t *testing.T) {
	expect := "engine B"
	car := NewCar(&EngineB{})

	if car.Ignite() != expect {
		t.Errorf("Expect sound to %s, but %s", expect, car.Ignite())
	}
}
