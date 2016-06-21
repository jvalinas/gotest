package artifact

import (
  "fmt"
  "github.com/nsf/termbox-go"
)

func init() {
}

type artifact struct {
  Id int
  Name string
  X float32
  Y float32
  DX float32
  DY float32
  AX float32
  AY float32
  Color termbox.Attribute
}

type Artifact struct {
  artifact artifact
}

func NewArtifact(
   id int,
   name string,
   x float32,
   y float32,
   dX float32,
   dY float32,
   aX float32,
   aY float32) *Artifact {
     object := new(Artifact)
     object.artifact.Id = id
     object.artifact.X = x
     object.artifact.Y = y
     object.artifact.DX = dX
     object.artifact.DY = dY
     object.artifact.AX = aX
     object.artifact.AY = aY
     object.artifact.Color = termbox.ColorYellow
     return object
}

func (object *Artifact) Pulse(width int, height int) {

  if object.artifact.X + object.artifact.DX > float32(width-1) || object.artifact.X + object.artifact.DX < 0.0 {
    object.artifact.DX = -object.artifact.DX
  }

  if object.artifact.Y + object.artifact.DY > float32(height-1) || object.artifact.Y + object.artifact.DY < 0 {
    object.artifact.DY = -object.artifact.DY
  }

  object.artifact.X += object.artifact.DX
  object.artifact.Y += object.artifact.DY

  object.artifact.DX += object.artifact.AX
  object.artifact.DY += object.artifact.AY
}

func (object Artifact) Color() termbox.Attribute {
  return object.artifact.Color
}

func (object *Artifact) SetColor(color termbox.Attribute) {
  //object.artifact.color = color
}

func (object Artifact) Id() int {
  return object.artifact.Id
}

func (object Artifact) Name() string {
  return object.artifact.Name
}

func (object Artifact) X() float32 {
  return object.artifact.X
}

func (object Artifact) Y() float32 {
  return object.artifact.Y
}

func (object Artifact) DX() float32 {
  return object.artifact.DX
}

func (object *Artifact) SetdX(dX float32) {
  object.artifact.DX = dX
}

func (object Artifact) DY() float32 {
  return object.artifact.DY
}

func (object *Artifact) SetdY(dY float32) {
  object.artifact.DY = dY
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
     object.artifact = NewArtifact(id, name, x, y, dX, dY, 0.0, 0.0)
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
