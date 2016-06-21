package main

import (
  "math/rand"
  "github.com/nsf/termbox-go"
  "artifact"
  "board"
  "core"
  "netinfo"
  "network"
)

func init() {
}

func main(){

  numSlots := 1
  slot := 0
  artifacts := 50
  width := 1000
  height := 1000
  serverIp := "172.17.0.3:10001"
  localAddr := ":10001"
  //fmt.Println("ServerIP: ", serverIp)

  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  termbox.HideCursor()
  defer termbox.Close()

  board := board.NewBoard(width, height)
  core := core.NewCore(numSlots, slot, board)
  queue1 := make(chan netinfo.NetPackage, 100)
  queue2 := make(chan netinfo.NetPackage, 100)
  go core.Run(queue1)
  go network.ServerStart(localAddr, queue2)
  go network.ClientStart(serverIp, queue1)

  for i:=0; i<artifacts; i++ {
    x, y := core.View().RandomPos()
    artifact := artifact.NewArtifact(i + 1000 * slot, "cosa", x, y,
                                                0.5,
                                                -2.0 + 4.0*rand.Float32(),
                                                -2.0 + 4.0*rand.Float32(),
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
