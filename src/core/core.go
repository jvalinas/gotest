package core

import (
  "time"
  "os/exec"
  "math"
  "github.com/nsf/termbox-go"
  "board"
  "view"
  "canvas"
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
     core.canvas = canvas.NewCanvas(termbox.Size())
     return core
}

func (core *Core) Collitions(current *artifact.Artifact) {
  for _, artifact := range core.board.Artifacts() {
    if current == artifact {
      continue
    }

    // Calc distance between centers
    b := float64(current.X() - artifact.X())
    a := float64(current.Y() - artifact.Y())
    distance := math.Sqrt(math.Pow(b, 2) + math.Pow(a, 2))

    if float32(distance) < current.R() + artifact.R() {
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
    time.Sleep(10*time.Millisecond)
  }
}

func main(){
}
