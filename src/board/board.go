package board

import (
  "fmt"
  "math/rand"
  "time"
  "artifact"
)

func init() {
  rand.Seed(int64(time.Now().Unix()))
}

type Board struct {
  width int
  height int
  artifacts map[int]artifact.Artifact
}

func NewBoard(width int, height int) *Board {
    b := new(Board)
    b.width = width
    b.height = height
    b.artifacts = make(map[int]artifact.Artifact)

    // b.buffer = make([]byte, b.width*b.height)
    return b
}

func (board Board) Artifacts() map[int]artifact.Artifact {
  return board.artifacts
}

func (board Board) AddArtifact(value *artifact.Artifact) {
  artf := artifact.Artifact(*value)
  board.artifacts[artf.Id()] = artf
  fmt.Println("Artifacts:", board.artifacts)
}
