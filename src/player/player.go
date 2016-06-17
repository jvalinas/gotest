package player

import (
  "fmt"
)

func init() {
}

type Player struct {
  id int
  name string
  slot int
}

func NewPlayer(id int, name string, slot int) *Player {
    p := new(Player)
    p.id = id
    p.name = name
    p.slot = slot
    return p
}

func (p Player) Id() int {
  return p.id
}

func (p Player) Name() string {
  return p.name
}

func (p Player) Slot() int {
  return p.slot
}

func (p Player) Print() {
  fmt.Printf("Player %d is %s on slot %d\n", p.id, p.name, p.slot )
}

func main(){
  // player := NewPlayer(1, "Bitelxux", 1)
}
