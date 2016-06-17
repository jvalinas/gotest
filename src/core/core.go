package core

import (
  "fmt"
  "time"
  "board"
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

func (core *Core) MoveArtifacts() {
  board := core.board
  for key, value := range board.Artifacts() {
    art := artifact.Artifact(value)
    fmt.Println("Key", key, "Artifact", art.Id())
    }
}

func (object *Core) Run() {
  i := 0
  for{
    object.MoveArtifacts()
    i += 1
    time.Sleep(100*time.Millisecond)
  }
}

func main(){
}
