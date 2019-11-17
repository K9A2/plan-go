package command

import "fmt"

type statusCommand struct {
}

func NewStatusCommand() *statusCommand {
  fmt.Println("status")
  return &statusCommand{}
}

func (command statusCommand) usage() {
  fmt.Println("Use <status> command to get status of all stored plans.")
  fmt.Println("")
  fmt.Println("usage: plan status")
}

func (command statusCommand) Execute() error {
  return nil
}
