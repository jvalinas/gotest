package main

import (
  "fmt"
  "math/rand"
  "github.com/nsf/termbox-go"
  "artifact"
  "board"
  "core"
  "physics"
  "udpproto"
  "os"
  "logging"
)

const (
  default_port="10001"
  loggingFile="my_test.log"
)
func init() {
}

func main(){

  if len(os.Args) < 2 {
      fmt.Println("Error: Insufficient arguments (ServerName needed)")
      return
  }
  numSlots := 2
  slot := 1
  serverName := os.Args[1]
  artifacts := 1
  width := 500
  height := 400
  var serverIp string
  switch serverName {
    case "Server1":
      slot = 0
      serverIp = "172.17.0.3:10001"
    case "Server2":
      slot = 1
      serverIp = "172.17.0.2:10001"
  }
  prefix := "[" + serverName + "] "
  logging.Init(loggingFile, prefix)

  //localAddr := ":10001"
  logging.Println("Server " + serverName + " started. Connecting to: ", serverIp)
  //fmt.Println("ServerIP: ", serverIp)

  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  termbox.HideCursor()
  defer termbox.Close()

  //width, height = termbox.Size()

  board := board.NewBoard(width, height)
  core := core.NewCore(numSlots, slot, board, serverName)
  queue1 := make(chan udpproto.GamePkg, 1000)
  queue2 := make(chan udpproto.GamePkg, 1000)
  queue3 := make(chan udpproto.EventPkg, 1000)
  go core.Run(queue1)
  udpproto.Start(serverIp, queue1, queue2, queue3)
  //go network.ServerStart(localAddr, queue2)
  //go network.ClientStart(serverIp, queue1)
  go core.UpdateBoard(queue2)

  for i:=0; i<artifacts; i++ {
    x, y := core.View().RandomPos()
    artifact := artifact.NewArtifact(i + slot * 1000, "cosa",
                                     *physics.NewVector(x, y),
                                     *physics.NewVector(-1 + 2*rand.Float64(),
                                                        -1 + 2*rand.Float64()),
                                      20 * rand.Float64(),
                                      5.0)
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
