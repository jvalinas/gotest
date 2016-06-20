package core

import (
  "time"
  "os/exec"
  "board"
  "view"
  "github.com/nsf/termbox-go"
  "artifact"
)

func init() {
}

type Core struct {
  board *board.Board
  numSlots int
  slot int
  view *view.View
}

func NewCore( numSlots int, slot int, board *board.Board ) *Core {
     core := new(Core)
     core.board = board
     core.numSlots = numSlots
     core.slot = slot
     core.view = view.NewView(numSlots, slot, board)
     return core
}

func (core *Core) Collitions(current *artifact.Artifact) {
  for _, artifact := range core.board.Artifacts() {
    if current == artifact {
      continue
    }
    if current.X() == artifact.X() && current.Y() == artifact.Y() {
      //current.SetColor(termbox.ColorRed)
      //artifact.SetColor(termbox.ColorRed)
      current.SetdX(0.0)
      current.SetdY(0.0)
      artifact.SetdX(0.0)
      artifact.SetdY(0.0)
      cmd := exec.Command("/usr/bin/beep")
      cmd.Start()
    }
  }
}

func (core *Core) View() *view.View {
  return core.view
}

func (core *Core) MoveArtifacts() {
  board := core.board
  board.SetSize(termbox.Size())

  for _, artifact := range board.Artifacts() {
    artifact.Pulse(board.Width(), board.Height())
    //x := int(artifact.X())
    //y := int(artifact.Y())
    //termbox.SetCell(x, y, 0x2500, artifact.Color(), artifact.Color())
    //if artifact.Color() != termbox.ColorRed {
    //    core.Collitions(artifact)
    //}
}
}

func (core *Core) Run() {
  for{
    //termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    core.MoveArtifacts()
    //termbox.Flush()
    time.Sleep(50*time.Millisecond)
  }
}

func main(){
}
