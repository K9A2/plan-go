package command

import "fmt"

type deleteCommand struct {
  args []string
}

func NewDeleteCommand(args []string) *deleteCommand {
  fmt.Println(args)
  return &deleteCommand{args: args}
}

func (command deleteCommand) usage() {
  fmt.Println("Use <delete> command to delete an existing plan.")
  fmt.Println("")
  fmt.Println("usage: plan delete <plan_id")
}

func (command deleteCommand) Execute() error {
  return nil
}
