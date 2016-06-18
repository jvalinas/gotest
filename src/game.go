package main

import (
  "board"
  "core"
  "math/rand"
  _ "player"
  "artifact"
  "github.com/nsf/termbox-go"
)

func init() {
}

func main(){
  /*
  board := board.NewBoard(20, 20)
  board.Fill(50)
  board.Print()

  //player1 :=
  } player.NewPlayer(1, "cnn", 0)
  //player2 := player.NewPlayer(2, "foo", 1)
  //map of players

  // create a map of players
  players := make(map[int]*player.Player)

  players[3] = player.NewPlayer(3, "three", 2)
  players[5] = player.NewPlayer(5, "five", 3)

  for key, value := range players {
    fmt.Println("key:", key, "id:", player.Player(*value).Id())
    fmt.Println("key:", key, "name:", player.Player(*value).Name())
    fmt.Println("key:", key, "slot:", player.Player(*value).Slot())
  }

  ball := artifact.NewBall(
    1,       // id
    "ball",  // name
    0.0,     // x
    0.0,     // y
    2.1,     // dX
    1.1,     // dY
    "green")

  for i:=0; i<1000; i++ {
    ball.Pulse()
    fmt.Println("Ball: ", *ball.Artifact())
  }
  */

  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  termbox.HideCursor()
  defer termbox.Close()

  width, height := termbox.Size()
  board := board.NewBoard(width, height)
  core := core.NewCore(board)
  go core.Run()

  for i:=0; i<100; i++ {
    artifact := artifact.NewArtifact(i, "cosa", float32(rand.Intn(width)),
                                                float32(rand.Intn(height)),
                                                -6.0 + float32(rand.Intn(6)),
                                                -4.0 + float32(rand.Intn(4)))
    board.AddArtifact(artifact)
  }

  //var b []byte = make([]byte, 1)
  //os.Stdin.Read(b)
  for {
    event := termbox.PollEvent()
    if event.Type == termbox.EventResize {
      width, height := termbox.Size()
      board.SetSize(width, height)
    }
    if event.Type == termbox.EventKey {
      break
    }
  }
}
