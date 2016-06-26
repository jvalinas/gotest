package artifact

import (
  "fmt"
  "github.com/nsf/termbox-go"
  "encoding/gob"
  "bytes"
  "physics"
)

func init() {
}

type artifact struct {
  Id int
  Name string
  Pos physics.Vector
  Dir physics.Vector
  Speed float64
  R float32
  Countdown int
  Color termbox.Attribute
}

type Artifact struct {
  artifact artifact
}

func NewArtifact(
   id int,
   name string,
   pos physics.Vector,
   dir physics.Vector,
   speed float64,
   r float32) *Artifact {
     object := new(Artifact)
     object.artifact.Id = id
     object.artifact.Pos = pos
     object.artifact.Dir = dir
     object.artifact.Speed = speed
     object.artifact.R = r
     object.artifact.Countdown = 0
     object.artifact.Color = termbox.ColorYellow
     return object
}

func (object *Artifact) Pulse(width int, height int) {

  normDir := object.artifact.Dir.Normalize()
  incPos := normDir.Scale(object.artifact.Speed)
  newPos := object.artifact.Pos.Sum(incPos)

  recalculate := false
  if newPos.X()  > float64(width-1) || newPos.X() < 0.0 {
    object.artifact.Dir.SetX(-object.artifact.Dir.X())
    recalculate = true
  }

  if newPos.Y()  > float64(height-1) || newPos.Y() < 0.0 {
    object.artifact.Dir.SetY(-object.artifact.Dir.Y())
    recalculate = true
  }

  if recalculate {
    normDir = object.artifact.Dir.Normalize()
    incPos = normDir.Scale(object.artifact.Speed)
    newPos = object.artifact.Pos.Sum(incPos)
  }

  object.artifact.Pos = newPos

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

func (object Artifact) Pos() physics.Vector {
  return object.artifact.Pos
}

func (object *Artifact) SetPos(pos physics.Vector) {
  object.artifact.Pos = pos
}

func (object Artifact) Dir() physics.Vector {
  return object.artifact.Dir
}

func (object *Artifact) SetDir(dir physics.Vector) {
  object.artifact.Dir = dir
}

func (object Artifact) Speed() float64 {
  return object.artifact.Speed
}

func (object *Artifact) SetSpeed(speed float64) {
  object.artifact.Speed = speed
}

func (object Artifact) R() float32 {
  return object.artifact.R
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
