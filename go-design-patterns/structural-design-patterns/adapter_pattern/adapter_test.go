package adapter_pattern

import (
	"testing"
)

func TestAdapter(t *testing.T) {

	adapter := NewAdapter(&Adaptee{})

	req := adapter.Request()

	if req != "Specific Request" {
		t.Errorf("Expect volume to %s, but %s", "Request", req)
	}
}
