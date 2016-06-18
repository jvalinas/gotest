package core

import (
  "board"
  "time"
  "github.com/nsf/termbox-go"
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

func (core *Core) MoveArtifacts() {
  board := core.board
  for _, artifact := range board.Artifacts() {
    artifact.Pulse(board.Width(), board.Height())
    //fmt.Println("Key:", key, "Pos:",
    //   artifact.X(), artifact.Y())
    x := int(artifact.X())
    y := int(artifact.Y())
    termbox.SetCell(x, y, 0x2500, termbox.ColorRed, termbox.ColorRed)
    }
}

func (object *Core) Run() {
  i := 0

  for{
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    object.MoveArtifacts()
    termbox.Flush()
    i += 1
    time.Sleep(50*time.Millisecond)
  }
}

func main(){
}
