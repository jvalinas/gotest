package physics

import (
  "math"
)

func init() {
}

type Vector struct {
  X float64
  Y float64
  Z int
}

func NewVector(x float64, y float64) *Vector {
  v := new(Vector)
  v.X = x
  v.Y = y
  return v
}


func (vector *Vector) Scale(factor float64) Vector {
  newVector := *NewVector(vector.X * factor, vector.Y * factor)
  return newVector
}

func (vector *Vector) Module() float64 {
  return (math.Sqrt(vector.X*vector.X+vector.Y*vector.Y))
}

func (vector *Vector) Sum(v Vector) Vector {
  return *NewVector(vector.X + v.X, vector.Y + v.Y)
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
