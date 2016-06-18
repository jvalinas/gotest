package core

import (
  "fmt"
  "time"
  "board"
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
  for key, artifact := range board.Artifacts() {
    artifact.Pulse()
    fmt.Println("Key:", key, "Pos:", artifact.X(), artifact.Y())
    }
}

func (object *Core) Run() {
  i := 0
  for{
    object.MoveArtifacts()
    i += 1
    time.Sleep(1*time.Millisecond)
  }
}

func main(){
}
