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
  "udpproto"
  "strconv"
  "player"
  "logging"
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
  players []*player.Player
}

func NewCore( numSlots int, slot int, board *board.Board, serverName string) *Core {
     core := new(Core)
     core.ServerName = serverName
     core.board = board
     core.numSlots = numSlots
     core.slot = slot
     core.view = view.NewView(numSlots, slot, board)
     core.canvas = canvas.NewCanvas(termbox.Size())
     players := make([]*player.Player, 0)
     players = append(players, player.NewPlayer(1, serverName, slot))
     core.players = players
     return core
}

func (core *Core) SetSlot(slot int) {
  core.slot = slot
  core.view = view.NewView(core.numSlots, slot, core.board)
}

func (core *Core) Collitions(current *artifact.Artifact) {
  for _, artifact := range core.board.Artifacts() {
    if current == artifact {
      continue
    }

    // Calc distance between centers
    b := float64(current.X() - artifact.X())
    a := float64(current.Y() - artifact.Y())
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
      //current.SetdX(0.0)
      //current.SetdY(0.0)
      //artifact.SetdX(0.0)
      //artifact.SetdY(0.0)
      current.SetCountdown(10)
      artifact.SetCountdown(10)

      if current.X() * artifact.X() < 0 {
        current.SetdX(-current.DX())
        artifact.SetdX(-artifact.DX())
      }
      if current.Y() * artifact.Y() < 0 {
        current.SetdY(-current.DY())
        artifact.SetdY(-artifact.DY())
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
      //if artifact.Color() != termbox.ColorRed {
      core.Collitions(artifact)
      //}
    }
  }
  return artifacts
}

func (core *Core) Run(queue chan udpproto.NetPacker) {
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

func (core *Core) getEmptySlots() []string {
    slots := make([]string, 0)
    for i:=0; i < core.numSlots; i++ {
      slotUsed := false
      for _, player := range core.players {
        if player.Slot() == i {
          slotUsed = true
          break
        }
       if ! slotUsed {
         slots = append(slots, strconv.Itoa(i))
       }
      }
    }
    return slots
}

func (core *Core) ProccessEvents(queue chan udpproto.EventPkg, out chan udpproto.NetPacker) {

  for {
    event := <- queue
    switch event.Etype {
    case "conn":
      switch event.Subtype {
      case "start":
          logging.Println("Event start ", event.ServerName)
          slots := core.getEmptySlots()
          pkg := udpproto.NewEventPkg(core.ServerName, "conn", "replay", slots)
          out <- pkg
      case "replay":
         logging.Println("Event replay ", event.ServerName)
          slots := event.Data
          logging.Println("Slots available ", slots)
          if len(slots) > 0 {
            slot, _ := strconv.Atoi(slots[0])
            core.SetSlot(slot)
            logging.Println("Starting game in slot ", slot)
            go core.Run(out)

          } else {
            logging.Println("No Slots available in ", event.ServerName)
          }

        default:
          logging.Println("Unknown Event subtype: #", event.Etype, "#-#", event.Subtype)
        }
      default:
        logging.Println("Unknown Event type: #", event.Etype, "#-#", event.Subtype)
      }
  }
}

func main(){
}
