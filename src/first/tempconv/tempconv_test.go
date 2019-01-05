// Tests for the package tempconv
package tempconv

import (
	"math"
	"testing"
)

func TestThreeWayFor5C(t *testing.T) {
	var c Celsius = 5 // 5°C
	var expf Fahrenheit = 41 // 41°F
	var expk Kelvin = 278.15 // 278.15K
	actf := CToF(c)
	if actf != expf {
		t.Errorf("expected: %v, actual: %v", expf, actf)
	}
	actk := CToK(c)
	if actk != expk {
		t.Errorf("expected: %v, actual: %v", expk, actk)
	}
	// back
	cBack := KToC(actk)
	if cBack != c {
		t.Errorf("expected: %v, actual: %v", c, cBack)
	}
	cBack = FToC(actf) // assignment
	if cBack != c {
		t.Errorf("expected: %v, actual: %v", c, cBack)
	}
}

func TestThreeWayForSameCF(t *testing.T) {
	var f Fahrenheit = -40 // 0°F
	var expc Celsius = -40 // -40°C
	var expk Kelvin = 233.15 // 233.15K
	actc := FToC(f)
	if actc != expc {
		t.Errorf("expected: %v, actual: %v", expc, actc)
	}
	actk := FToK(f)
	if math.Abs(float64(actk - expk)) > 1e-4 {
		t.Errorf("expected: %v, actual: %v", expk, actk)
	}
	// back
	fBack := KToF(actk)
	if fBack != f {
		t.Errorf("expected: %v, actual: %v", f, fBack)
	}
	fBack = CToF(actc) // assignment
	if fBack != f {
		t.Errorf("expected: %v, actual: %v", f, fBack)
	}
}
