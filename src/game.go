package main

import (
  "math/rand"
  "github.com/nsf/termbox-go"
  "artifact"
  "board"
  "core"
)

func init() {
}

func main(){

  numSlots := 1
  slot := 0
  artifacts := 100
  width := 500
  height := 500

  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  termbox.HideCursor()
  defer termbox.Close()

  // remove this.
  width, height = termbox.Size()

  board := board.NewBoard(width, height)
  core := core.NewCore(numSlots, slot, board)
  go core.Run()

  for i:=0; i<artifacts; i++ {
    x, y := core.View().RandomPos()
    artifact := artifact.NewArtifact(i, "cosa", x, y,
                                                0.5,
                                                -1.0 + 2.0*rand.Float32(),
                                                -1.0 + 2.0*rand.Float32(),
                                                0.0, 0.0 )

    board.AddArtifact(artifact)
  }

  for {
    event := termbox.PollEvent()
    if event.Type == termbox.EventResize {
      core.Canvas().Resize(termbox.Size())
    }
    if event.Type == termbox.EventKey {
      break
    }
  }
}
