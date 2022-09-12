package switchsample

import "testing"

func TestSwitchOrder(t *testing.T) {
	SwtichExecOrder(3)
}

func TestTypeSwitch(t *testing.T) {
	SwitchType(10, &Dog{Name: "bucky"})
}
