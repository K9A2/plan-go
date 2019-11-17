package command

import "fmt"

type addCommand struct {
  args []string
}

func NewAddCommand(args []string) *addCommand {
  fmt.Println(args)
  return &addCommand{args: args}
}

func (command addCommand) usage() {
  fmt.Println("Use <add> command to add a new major plan, or add as a child")
  fmt.Println("plan of existing plan.")
  fmt.Println("")
  fmt.Println("usage:")
  fmt.Println("  1. plan add 'plan_title'")
  fmt.Println("  2. plan add 'plan_title' -p <parent_plan_id>")
}

func (command *addCommand) Execute() error {
  return nil
}
