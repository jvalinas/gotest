package main

import (
  "board"
  "core"
  "time"
  "math/rand"
  "artifact"
  _"github.com/nsf/termbox-go"
)

func init() {
}

func main(){

  numSlots := 2
  slot := 0
  artifacts := 2000
  width := 4000
  height := 1000

  /*
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  termbox.HideCursor()
  defer termbox.Close()
  */

  //width, height := termbox.Size()

  board := board.NewBoard(width, height)
  core := core.NewCore(numSlots, slot, board)
  go core.Run()

  for i:=0; i<artifacts; i++ {
    artifact := artifact.NewArtifact(i, "cosa", float32(rand.Intn(width)),
                                                float32(rand.Intn(height)),
                                                -3.0 + float32(rand.Intn(6)),
                                                -2.0 + float32(rand.Intn(4)),
                                                0.0, 0.0 )
                                                //-0.1 + rand.Float32()/10.0,
                                                //-0.1 + rand.Float32()/10.0)
    board.AddArtifact(artifact)
  }

  for {
    time.Sleep(10*time.Second)
  }

}
