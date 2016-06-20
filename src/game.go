package main

import (
  "fmt"
  "time"
  "math/rand"
  _"github.com/nsf/termbox-go"
  "artifact"
  "board"
  "core"
)

func init() {
}

func main(){

  numSlots := 2
  slot := 1
  artifacts := 1
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
    x, y := core.View().RandomPos()
    artifact := artifact.NewArtifact(i, "cosa", x, y,
                                                -3.0 + float32(rand.Intn(6)),
                                                -2.0 + float32(rand.Intn(4)),
                                                0.0, 0.0 )

    board.AddArtifact(artifact)
    ISeeYou := core.View().ISeeYou(artifact)
    fmt.Printf("[%f, %f] -> %t\n", artifact.X(), artifact.Y(), ISeeYou)

  }

  for {
    time.Sleep(10*time.Second)
  }

}
