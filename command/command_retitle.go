package command

import "fmt"

type retitleCommand struct {
  args []string
}

func NewRetitleCommand(args []string) *retitleCommand {
  return &retitleCommand{args: args}
}

func (command retitleCommand) usage() {
  fmt.Println("Use <retitle> to retitle an existing plan.")
  fmt.Println("")
  fmt.Println("usage: plan retitle <plan_id> 'a_new_title'")
}

func (command *retitleCommand) Execute() error {
  return nil
}
