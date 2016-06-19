package core

import (
  "board"
  "time"
  "github.com/nsf/termbox-go"
  "artifact"
)

func init() {
}

type Core struct {
  board *board.Board
  objects []*artifact.Artifact
}

func NewCore( board *board.Board ) *Core {
     core := new(Core)
     core.board = board
     core.objects = make([]*artifact.Artifact, core.board.Width()*core.board.Height())
     for _, artifact := range core.board.Artifacts() {
       i := int(artifact.Y()) * core.board.Width() + int(artifact.X())
       core.objects[i] = artifact
     }

     return core
}

func (core *Core) MoveArtifacts() {
  board := core.board

  for _, artifact := range board.Artifacts() {
    artifact.Pulse(board.Width(), board.Height(), core.objects)
    x := int(artifact.X())
    y := int(artifact.Y())
    termbox.SetCell(x, y, 0x2500, artifact.Color(), artifact.Color())
    core.objects[y*core.board.Width() + x] = artifact
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
