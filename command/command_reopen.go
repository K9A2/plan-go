package command

import (
  "fmt"
  "github.com/stormlin/plan-go/util"
)

type reopenCommand struct {
  args []string
}

func NewReopenCommand(args []string) *reopenCommand {
  return &reopenCommand{args: args}
}

func (command reopenCommand) usage() {
  fmt.Println("Use <reopen> command to reopen a done plan.")
  fmt.Println("")
  fmt.Println("usage: plan reopen <plan_id>")
}

func (command reopenCommand) Execute() error {
  if len(command.args) != 1 {
    return &util.WrongArgumentNumberError{}
  }

  planList, err := util.ReadFromJsonFile(util.DefaultFilePath)
  if err != nil {
    return err
  }

  var targetPlan *util.PlanItem
  for _, plan := range planList.MajorPlan {
    result := util.FindPlan(plan, command.args[0])
    if result != nil {
      targetPlan = result
      break
    }
  }
  if targetPlan == nil {
    return &util.PlanNotExistsError{}
  }

  util.MarkAsOpen(targetPlan)
  err = util.SaveAsJsonFile(util.DefaultFilePath, planList)
  if err != nil {
    return err
  }

  return nil
}
