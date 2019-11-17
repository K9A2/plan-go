package main

import (
  "fmt"
  "github.com/stormlin/plan-go/command"
  "github.com/stormlin/plan-go/util"
  "os"
)

func main() {
  // remove the program name
  args := os.Args[1:]
  if len(args) <= 0 {
    fmt.Println("error: not enough arguments provided.")
    return
  }

  // switch by command type constants
  var err error
  switch args[0] {
  case util.CommandAdd:
    commandAdd := command.NewAddCommand(args[1:])
    err = commandAdd.Execute()
    break
  case util.CommandDone:
    doneCommand := command.NewDoneCommand(args[1:])
    err = doneCommand.Execute()
    break
  case util.CommandStatus:
    statusCommand := command.NewStatusCommand()
    err = statusCommand.Execute()
    break
  case util.CommandReopen:
    reopenCommand := command.NewReopenCommand(args)
    err = reopenCommand.Execute()
    break
  case util.CommandDelete:
    deleteCommand := command.NewDeleteCommand(args[1:])
    err = deleteCommand.Execute()
    break
  case util.CommandRetitle:
    retitleCommand := command.NewRetitleCommand(args[1:])
    err = retitleCommand.Execute()
    break
  default:
    err = &util.UnrecognizedCommandError{}
  }

  if err != nil {
    fmt.Printf("Runtime error, err: \"%s\"", err.Error())
    return
  }
  return
}
