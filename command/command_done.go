package command

import (
  "fmt"
  "github.com/stormlin/plan-go/util"
)

type doneCommand struct {
  args []string
}

func NewDoneCommand(args []string) *doneCommand {
  return &doneCommand{args: args}
}

func (command doneCommand) usage() {
  fmt.Println("Use <done> command to mark a plan as done.")
  fmt.Println("")
  fmt.Println("usage: plan done <plan_id>")
}

func (command doneCommand) Execute() error {
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

  // mark it and its children as done
  util.MarkAsDone(targetPlan)
  err = util.SaveAsJsonFile(util.DefaultFilePath, planList)
  if err != nil {
    return err
  }
  return nil
}
