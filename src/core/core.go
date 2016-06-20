package core

import (
  "time"
  "os/exec"
  "board"
  "view"
  "canvas"
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
  canvas *canvas.Canvas
}

func NewCore( numSlots int, slot int, board *board.Board ) *Core {
     core := new(Core)
     core.board = board
     core.numSlots = numSlots
     core.slot = slot
     core.view = view.NewView(numSlots, slot, board)
     core.canvas = canvas.NewCanvas(100, 70)
     return core
}

func (core *Core) Collitions(current *artifact.Artifact) {
  for _, artifact := range core.board.Artifacts() {
    if current == artifact {
      continue
    }
    if current.X() == artifact.X() && current.Y() == artifact.Y() {
      current.SetColor(termbox.ColorRed)
      artifact.SetColor(termbox.ColorRed)
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

  for _, artifact := range board.Artifacts() {
    if core.View().ISeeYou(artifact) {
      artifact.Pulse(board.Width(), board.Height())
      if artifact.Color() != termbox.ColorRed {
          core.Collitions(artifact)
      }
      core.canvas.Draw(core.view, artifact)
    }

}
}

func (core *Core) Run() {
  for{
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    core.MoveArtifacts()
    termbox.Flush()
    time.Sleep(50*time.Millisecond)
  }
}

func main(){
}
