package command

import (
  "fmt"
  "github.com/stormlin/plan-go/util"
)

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
  if len(command.args) != 2 {
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

  newTitle := command.args[1]
  // mark it and its children as done
  targetPlan.Title = newTitle
  err = util.SaveAsJsonFile(util.DefaultFilePath, planList)
  if err != nil {
    return err
  }
  return nil
}
