package strategy

import "testing"

func TestNewStrategy(t *testing.T) {
	s, err := NewStrategy("3")
	if err != nil {
		t.Logf("err: %+v", err)
		return
	}
	s.fun()
}
