package view

import (
  "board"
  "artifact"
  "math/rand"
)

func init() {
}

type View struct{
  slot int
  x0 int
  x1 int
  y0 int
  y1 int
}

func (view View) RandomPos() (float64, float64) {
  x := float64(view.x0) + float64(view.x1 - view.x0)*rand.Float64()
  y := float64(view.y0) + float64(view.y1 - view.y0)*rand.Float64()

  return x, y
}

func (view View) ISeeYou(artifact *artifact.Artifact) bool {

  pos := artifact.Pos()
  return pos.X() >= float64(view.x0) && pos.X() <= float64(view.x1) &&
     pos.Y() >= float64(view.y0) && pos.Y() <= float64(view.y1)

}

func NewView(numSlots int, slot int, board *board.Board) *View {
     slotSize := int(board.Width()/numSlots)
     view := new(View)
     view.x0 = slotSize * slot
     view.x1 = view.x0 + slotSize
     view.y0 = 0
     view.y1 = board.Height()
     view.slot = slot
     return view
}

func (this View) Slot() int {
  return this.slot
}

func (this View) Width() int {
  return this.x1 - this.x0
}

func (this View) Height() int {
  return this.y1 - this.y0
}

func (this View) X0() int {
  return this.x0
}

func (this View) Y0() int {
  return this.y0
}
