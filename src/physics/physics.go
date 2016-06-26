package physics

import (
  "math"
)

func init() {
}

type Vector struct {
  x float64
  y float64
  Z int
}

func NewVector(x float64, y float64) *Vector {
  v := new(Vector)
  v.x = x
  v.y = y
  return v
}

func (vector Vector) X() float64 {
  return vector.x
}

func (vector *Vector) SetX(x float64) {
  vector.x = x
}

func (vector Vector) Y() float64 {
  return vector.y
}

func (vector *Vector) SetY(y float64) {
  vector.y = y
}

func (vector *Vector) Scale(factor float64) Vector {
  newVector := *NewVector(vector.x * factor, vector.y * factor)
  return newVector
}

func (vector *Vector) Module() float64 {
  return (math.Sqrt(vector.x*vector.x+vector.y*vector.y))
}

func (vector *Vector) Sum(v Vector) Vector {
  return *NewVector(vector.x + v.x, vector.y + v.y)
}

func (vector *Vector) Normalize() Vector {
  newVector := *NewVector(0.0, 0.0)
  module := vector.Module()
  if (module != 0) {
      newVector = vector.Scale(1.0 / module)
  } else
  {
      newVector = *NewVector(0.0, 0.0)
  }

  return newVector
}
