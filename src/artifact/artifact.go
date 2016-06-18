package artifact

import (
  "fmt"
)

func init() {
}

type Artifact struct {
  id int
  name string
  x float32
  y float32
  dX float32
  dY float32
}

func NewArtifact(
   id int,
   name string,
   x float32,
   y float32,
   dX float32,
   dY float32) *Artifact {
     object := new(Artifact)
     object.id = id
     object.x = x
     object.y = y
     object.dX = dX
     object.dY = dY
     return object
}

func (object *Artifact) Pulse(width int, height int) {

  if object.x + object.dX > float32(width-1) || object.x + object.dX < 0.0 {
    object.dX = -object.dX
  }

  if object.y + object.dY > float32(height-1) || object.y + object.dY < 0 {
    object.dY = -object.dY
  }

  object.x += object.dX
  object.y += object.dY
}

func (object Artifact) Id() int {
  return object.id
}

func (object Artifact) Name() string {
  return object.name
}

func (object Artifact) X() float32 {
  return object.x
}

func (object Artifact) Y() float32 {
  return object.y
}

func (object Artifact) DX() float32 {
  return object.dX
}

func (object Artifact) DY() float32 {
  return object.dY
}

////////////////////////////////////////////////////////////////////////////

type Ball struct {
  artifact *Artifact
  color string
}

func NewBall(
   id int,
   name string,
   x float32,
   y float32,
   dX float32,
   dY float32,
   color string) *Ball {
     object := new(Ball)
     object.artifact = NewArtifact(id, name, x, y, dX, dY)
     object.color = color
     return object
}

func (ball Ball) Artifact() *Artifact {
  return ball.artifact
}

func (ball Ball) Color() string {
  return ball.color
}

////////////////////////////////////////////////////////////////////////////

func main(){
  fmt.Println("Artifact")
}
