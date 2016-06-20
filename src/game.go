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

  numSlots := 2
  slot := 1
  artifacts := 1
  width := 4000
  height := 1000

  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  termbox.HideCursor()
  defer termbox.Close()

  board := board.NewBoard(width, height)
  core := core.NewCore(numSlots, slot, board)
  go core.Run()

  for i:=0; i<artifacts; i++ {
    x, y := core.View().RandomPos()
    artifact := artifact.NewArtifact(i, "cosa", x, y,
                                                -50.0 + 100.0*rand.Float32(),
                                                -50.0 + 100.0*rand.Float32(),
                                                0.0, 0.0 )

    board.AddArtifact(artifact)
  }

  for {
    event := termbox.PollEvent()
    if event.Type == termbox.EventResize {
    }
    if event.Type == termbox.EventKey {
      break
    }
  }
}
