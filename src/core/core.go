package core

import (
  "board"
  "time"
  "os/exec"
  "github.com/nsf/termbox-go"
  "artifact"
)

func init() {
}

type Core struct {
  board *board.Board
}

func NewCore( board *board.Board ) *Core {
     core := new(Core)
     core.board = board
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

func (core *Core) MoveArtifacts() {
  board := core.board
  width, height := termbox.Size()
  board.SetSize(width, height)

  for _, art := range board.Artifacts() {
    art.Pulse(width, height)
    x := int(art.X())
    y := int(art.Y())
    termbox.SetCell(x, y, 0x2500, art.Color(), art.Color())
    if art.Color() != termbox.ColorRed {
      core.Collitions(art)
    }
}
}

func (object *Core) Run() {
  for{
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    object.MoveArtifacts()
    termbox.Flush()
    time.Sleep(50*time.Millisecond)
  }
}

func main(){
}
