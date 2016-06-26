package main

import (
  _"fmt"
  "math/rand"
  "github.com/nsf/termbox-go"
  "artifact"
  "board"
  "core"
  "physics"
  "netinfo"
  "network"
  "os"
  "strconv"
)

func init() {
}

func main(){

  numSlots := 1
  slot, _ := strconv.Atoi(os.Args[1])
  artifacts := 50
  width := 1000
  height := 500
  serverIp := "172.17.0.7:10001"
  if slot == 1 {
    serverIp = "172.17.0.5:10001"
  }

  localAddr := ":10001"
  //fmt.Println("ServerIP: ", serverIp)

  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  termbox.HideCursor()
  defer termbox.Close()

  //width, height = termbox.Size()

  board := board.NewBoard(width, height)
  core := core.NewCore(numSlots, slot, board)
  queue1 := make(chan netinfo.NetPackage, 1000)
  queue2 := make(chan netinfo.NetPackage, 1000)
  go core.Run(queue1)
  go network.ServerStart(localAddr, queue2)
  go network.ClientStart(serverIp, queue1)
  go core.UpdateBoard(queue2)

  for i:=0; i<artifacts; i++ {
    x, y := core.View().RandomPos()
    artifact := artifact.NewArtifact(i + slot * 1000, "cosa",
                                     *physics.NewVector(x, y),
                                     *physics.NewVector(-10.0 + 20.0*rand.Float64(),
                                                        -10.0 + 20.0*rand.Float64()),
                                      20 * rand.Float64(),
                                      1.0)

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
