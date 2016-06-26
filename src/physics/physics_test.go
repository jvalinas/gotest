package physics

import (
  "testing"
)

func floatEquals(a, b float64) bool {
  EPSILON := 0.00000001
	if ((a - b) < EPSILON && (b - a) < EPSILON) {
		return true
	}
	return false
}

func TestVector(t *testing.T) {
  v := NewVector(4.0, 2.0)
  if v.x != 4.0 || v.y != 2.0 {
    t.Error("Expected {4.0, 2.0}, got ", v)
  }
}

func TestNormalize(t *testing.T) {
  v := NewVector(4.0, 2.0)
  normalized := v.Normalize()
  module := normalized.Module()
  if !floatEquals(module, 1.0) {
    t.Error("Expected 1.0, got ", module)
  }
}

func TestSum(t *testing.T) {
  v := NewVector(4.0, 2.0)
  v2 := *NewVector(1.0, 2.0)
  v3 := v.Sum(v2)
  if v3.x != 5.0 || v3.y != 4.0 {
    t.Error("Expected {5.0, 4.0}, got ", v3)
  }
}

func TestScale(t *testing.T) {
  v := NewVector(4.0, 2.0)
  v3 := v.Scale(2.0)
  if v3.x != 8.0 || v3.y != 4.0 {
    t.Error("Expected {8.0, 4.0}, got ", v3)
  }
}
