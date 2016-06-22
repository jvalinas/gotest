package canvas

import (
  "fmt"
  "artifact"
  "view"
  "github.com/nsf/termbox-go"
)

func init() {
}

type Canvas struct{
  width int
  height int
}

func NewCanvas(width int, height int) *Canvas {
     canvas := new(Canvas)
     canvas.width = width
     canvas.height = height
     return canvas
}

func (this *Canvas) Log() {
  fmt.Println(this)
}

func (this *Canvas) Width() int {
  return this.width
}

func (this *Canvas) Height() int {
  return this.height
}

func (this *Canvas) Resize(width int, height int) {
  this.width = width
  this.height = height
}

func (this Canvas) Draw(view *view.View, artifact *artifact.Artifact) {

  xFactor := 1.0 / (float32(view.Width()) / float32(this.Width()))
  yFactor := 1.0 / (float32(view.Height()) / float32(this.Height()))
  relX := artifact.X() - float32(view.X0())
  relY := artifact.Y() - float32(view.Y0())

  x := int(relX * xFactor)
  y := int(relY * yFactor)

  /*
  fmt.Printf("%f %f [%f, %f] -> [%f, %f] -> [%d, %d]\n",
    xFactor, yFactor,
    artifact.X(), artifact.Y(),
    relX, relY,
    x, y)
  */

  color := artifact.Color()
  if artifact.Countdown() > 0 {
    color = termbox.ColorWhite
  }

  termbox.SetCell(x, y, 0x004F, color, termbox.ColorDefault)
}
