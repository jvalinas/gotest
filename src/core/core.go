package core

import (
  "time"
  "os/exec"
  "math"
  "github.com/nsf/termbox-go"
  "board"
  "view"
  "canvas"
  "artifact"
  "physics"
  "udpproto"
)

func init() {
}

type Core struct {
  ServerName string
  board *board.Board
  numSlots int
  slot int
  view *view.View
  canvas *canvas.Canvas
}

func NewCore( numSlots int, slot int, board *board.Board, serverName string) *Core {
     core := new(Core)
     core.ServerName = serverName
     core.board = board
     core.numSlots = numSlots
     core.slot = slot
     core.view = view.NewView(numSlots, slot, board)
     core.canvas = canvas.NewCanvas(termbox.Size())
     return core
}

func (core *Core) Collitions(current *artifact.Artifact) {

  currentPos := current.Pos()

  for _, artifact := range core.board.Artifacts() {
    if current == artifact {
      continue
    }

    artifactPos := artifact.Pos()

    // Calc distance between centers
    b := float64(currentPos.X - artifactPos.X)
    a := float64(currentPos.Y - artifactPos.Y)
    distance := math.Sqrt(math.Pow(b, 2) + math.Pow(a, 2))

    // Probe of close but not collitioning objects
    if float32(distance) < current.R() + artifact.R() + 2.0 {

      if current.Color() == termbox.ColorYellow {
        current.SetColor(termbox.ColorCyan)
      }

      if artifact.Color() == termbox.ColorYellow {
        artifact.SetColor(termbox.ColorCyan)
      }

    }

    if float32(distance) < current.R() + artifact.R() {
      current.SetColor(termbox.ColorRed)
      artifact.SetColor(termbox.ColorRed)

      current.SetCountdown(20)
      artifact.SetCountdown(20)

      if current.Dir().X * artifact.Dir().X < 0 {
        current.SetDir(*physics.NewVector(-current.Dir().X, current.Dir().Y))
        artifact.SetDir(*physics.NewVector(-artifact.Dir().X, artifact.Dir().Y))
      }

      if current.Dir().Y * artifact.Dir().Y < 0 {
        current.SetDir(*physics.NewVector(current.Dir().X, -current.Dir().Y))
        artifact.SetDir(*physics.NewVector(artifact.Dir().X, -artifact.Dir().Y))
      }

      cmd := exec.Command("/usr/bin/beep")
      cmd.Start()
    }
  }
}

func (core *Core) View() *view.View {
  return core.view
}

func (core *Core) Canvas() *canvas.Canvas {
  return core.canvas
}

func (core *Core) MoveArtifacts() map[int]*artifact.Artifact {
  board := core.board
  artifacts := make(map[int]*artifact.Artifact)

  for _, artifact := range board.Artifacts() {
    if core.View().ISeeYou(artifact) {
      artifact.Pulse(board.Width(), board.Height())
      artifacts[artifact.Id()] = artifact
      core.canvas.Draw(core.view, artifact)
      core.Collitions(artifact)
    }
  }
  return artifacts
}

func (core *Core) Run(queue chan udpproto.GamePkg) {
  for{
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    artifactsMoved := core.MoveArtifacts()
    if len(artifactsMoved) != 0 {
      info := udpproto.NewGamePkg(core.ServerName, artifactsMoved)
      queue <- info
    }

    termbox.Flush()
    time.Sleep(100*time.Millisecond)
  }
}

func (core *Core) UpdateBoard(queue chan udpproto.GamePkg) {
  //var netPkg netinfo.NetPackage
  for {
    netPkg := <-queue
    core.board.MergeArtifacts(netPkg.Artifacts)
  }
}

func main(){
}
