package artifact

import (
  "testing"
  "physics"
  "fmt"
)

func defaultArtifact() *Artifact {
  artifact := NewArtifact( 1,
                           "foo",
                            *physics.NewVector(1.0, 1.0),
                            *physics.NewVector(1.0, 1.0),
                            2.0, // speed
                            1.0) // radio
  return artifact
}

func floatEquals(a, b float64) bool {
  EPSILON := 0.00000001
	if ((a - b) < EPSILON && (b - a) < EPSILON) {
		return true
	}
	return false
}

func TestArtifact(t *testing.T) {
  artifact := defaultArtifact()
  if artifact == nil {
    t.Error("Failed")
  }
}

func TestMoveArtifact(t *testing.T) {
  artifact := defaultArtifact()
  artifact1 := defaultArtifact()
  artifact1.Pulse(100,100)

  if artifact1.Pos().X() < artifact.Pos().X() || artifact1.Pos().Y() < artifact.Pos().Y() {
    fmt.Println(artifact)
    fmt.Println(artifact1)
    t.Error("Wrong movement")
  }

  // Test change direction
  pos := *physics.NewVector(99.0,99.0)
  artifact.SetPos(pos)
  artifact.Pulse(100,100)
  if artifact.Dir().X() > 1 || artifact.Dir().Y() > 1 {
    t.Error("Failed changing direction")
  }

}
