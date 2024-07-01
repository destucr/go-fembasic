package event

import "fmt"

type Ticket struct {
  ID int
  Event string
}

func (t Ticket) PrintTicket() {
  fmt.Println(t.Event)
}
