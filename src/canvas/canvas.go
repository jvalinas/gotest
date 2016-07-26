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

  pos := artifact.Pos()

  xFactor := 1.0 / (float64(view.Width()) / float64(this.Width()))
  yFactor := 1.0 / (float64(view.Height()) / float64(this.Height()))
  relX := pos.X - float64(view.X0())
  relY := pos.Y - float64(view.Y0())

  x := int(relX * xFactor)
  y := int(relY * yFactor)

  color := artifact.Color()
  if artifact.Countdown() > 0 {
    color = termbox.ColorWhite
  }

  termbox.SetCell(x, y, 0x004F, color, termbox.ColorDefault)
}
