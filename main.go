package main

import (
  "femBasics/imports"
)

func main() {
  myTicket := event.Ticket{
    ID: 23,
    Event: "Pestapora",
  }

  myTicket.PrintTicket()

}
