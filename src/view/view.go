package view

import (
  "fmt"
  "board"
  "artifact"
  "math/rand"
)

func init() {
}

type View struct{
  x0 int
  x1 int
  y0 int
  y1 int
}

func (view View) RandomPos() (float32, float32) {
  x := float32(view.x0) + float32(view.x1 - view.x0)*rand.Float32()
  y := float32(view.y0) + float32(view.y1 - view.y0)*rand.Float32()

  return x, y
}

func (view View) ISeeYou(artifact *artifact.Artifact) bool {

  return artifact.X() >= float32(view.x0) && artifact.X() <= float32(view.x1) &&
     artifact.Y() >= float32(view.y0) && artifact.Y() <= float32(view.y1)

}

func NewView(numSlots int, slot int, board *board.Board) *View {
     slotSize := int(board.Width()/numSlots)
     view := new(View)
     view.x0 = slotSize * slot
     view.x1 = view.x0 + slotSize
     view.y0 = 0
     view.y1 = board.Height()

     fmt.Println("Slot:", slot,
       "x0:", view.x0,
       "x1:", view.x1,
       "y0:", view.y0,
       "y1:", view.y1)

     return view
}

func (this *View) Width() int {
  return this.x1 - this.x0
}

func (this *View) Height() int {
  return this.y1 - this.y0
}
