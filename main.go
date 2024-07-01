package main

import (
  "femBasics/imports"
)

func main() {
  myTicket := imports.Ticket{
    ID: 23,
    Event: "Pestapora",
  }

  myTicket.PrintTicket()

}
