package main

import (
  "fmt"
  "board"
  "core"
  "time"
  _ "player"
  "artifact"
)

func init() {
  fmt.Println("init game")
}

func main(){
  /*
  board := board.NewBoard(20, 20)
  board.Fill(50)
  board.Print()

  //player1 := player.NewPlayer(1, "cnn", 0)
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

  board := board.NewBoard(1000, 1000)
  core := core.NewCore(board)
  go core.Run()

  artifact1 := artifact.NewArtifact(1, "art1", 0.0, 0.0, 0.0, 1.0)
  artifact2 := artifact.NewArtifact(2, "art2", 0.0, 0.0, 1.0, 0.0)

  board.AddArtifact(artifact1)
  board.AddArtifact(artifact2)

  for{
    time.Sleep(10*time.Second)
  }

}
