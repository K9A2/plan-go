package command

import (
  "fmt"
)

type reopenCommand struct {
  args []string
}

func NewReopenCommand(args []string) *reopenCommand {
  fmt.Println(args)
  return &reopenCommand{args: args}
}

func (command reopenCommand) usage() {
  fmt.Println("Use <reopen> command to reopen a done plan.")
  fmt.Println("")
  fmt.Println("usage: plan reopen <plan_id>")
}

func (command reopenCommand) Execute() error {
  return nil
}
