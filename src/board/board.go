package board

import (
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
  artifacts map[int]*artifact.Artifact
}

func NewBoard(width int, height int) *Board {
    b := new(Board)
    b.width = width
    b.height = height
    b.artifacts = make(map[int]*artifact.Artifact)
    return b
}

func (board *Board) SetSize(width int, height int) {
  board.width = width
  board.height = height
}

func (board Board) Width() int {
  return board.width
}

func (board Board) Height() int {
  return board.height
}

func (board Board) Artifacts() map[int]*artifact.Artifact {
  return board.artifacts
}

func (board Board) AddArtifact(value *artifact.Artifact) {
  board.artifacts[value.Id()] = value
}

func (board Board) MergeArtifacts(artifacts map[int]*artifact.Artifact) {
  for _, value := range artifacts {
    board.artifacts[value.Id()] = value
  }
}
