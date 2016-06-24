package artifact

import (
  "fmt"
  "github.com/nsf/termbox-go"
  "encoding/gob"
  "bytes"
)

func init() {
}

type artifact struct {
  Id int
  Name string
  X float32
  Y float32
  R float32
  DX float32
  DY float32
  AX float32
  AY float32
  Countdown int
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
   r float32,
   dX float32,
   dY float32,
   aX float32,
   aY float32) *Artifact {
     object := new(Artifact)
     object.artifact.Id = id
     object.artifact.X = x
     object.artifact.Y = y
     object.artifact.R = r
     object.artifact.DX = dX
     object.artifact.DY = dY
     object.artifact.AX = aX
     object.artifact.AY = aY
     object.artifact.Countdown = 0
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

  if object.artifact.Countdown > 0 {
    object.artifact.Countdown -= 1
  }

}

func (object Artifact) Color() termbox.Attribute {
  return object.artifact.Color
}

func (object *Artifact) SetColor(color termbox.Attribute) {
  object.artifact.Color = color
}

func (object Artifact) Countdown() int {
  return object.artifact.Countdown
}

func (object *Artifact) SetCountdown(counter int) {
  object.artifact.Countdown = counter
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

func (object Artifact) R() float32 {
  return object.artifact.R
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

func (object *Artifact) SetPos(x float32, y float32) {
  object.artifact.X = x
  object.artifact.Y = y
}

func (object Artifact) Pos() (float32, float32) {
  return object.artifact.X, object.artifact.Y
}

func NewActifactFromBytes(data []byte) Artifact {
  d := gob.NewDecoder(bytes.NewBuffer(data))

  // Decoding the serialized data
  var artifact artifact
  err := d.Decode(&artifact)
  if err != nil {
      panic(err)
  }

  //fmt.Printf("%#v\n", artifact)
  return Artifact{artifact: artifact}
}

func (a Artifact) Bytes() []byte{
  data := new(bytes.Buffer)

  e := gob.NewEncoder(data)

  // Encoding
  err := e.Encode(a.artifact)
  if err != nil {
      panic(err)
  }
  return data.Bytes()
}


////////////////////////////////////////////////////////////////////////////

func main(){
  fmt.Println("Artifact")
}
