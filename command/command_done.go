package command

import "fmt"

type doneCommand struct {
  args []string
}

func NewDoneCommand(args []string) *doneCommand {
  fmt.Println(args)
  return &doneCommand{args: args}
}

func (command doneCommand) usage() {
  fmt.Println("Use <done> command to mark a plan as done.")
  fmt.Println("")
  fmt.Println("usage: plan done <plan_id>")
}

func (command doneCommand) Execute() error {
  return nil
}
