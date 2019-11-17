package command

import (
  "fmt"
  "github.com/stormlin/plan-go/util"
)

type deleteCommand struct {
  args []string
}

func NewDeleteCommand(args []string) *deleteCommand {
  return &deleteCommand{args: args}
}

func (command deleteCommand) usage() {
  fmt.Println("Use <delete> command to delete an existing plan.")
  fmt.Println("")
  fmt.Println("usage: plan delete <plan_id")
}

func (command deleteCommand) Execute() error {
  if len(command.args) != 1 {
    return &util.WrongArgumentNumberError{}
  }

  planList, err := util.ReadFromJsonFile(util.DefaultFilePath)
  if err != nil {
    return err
  }

  targetPlanId := command.args[0]
  var parentPlan *util.PlanItem
  var majorPlanToRemove *util.PlanItem
  var majorPlanToRemoveIndex int
  var targetIndex = -1

  for index, major := range planList.MajorPlan {
    if major.PlanId == targetPlanId {
      // remove this major plan
      majorPlanToRemove = major
      majorPlanToRemoveIndex = index
      break
    }
    result := util.FindParent(major, targetPlanId)
    if result != nil {
      parentPlan = result
    }
  }
  if majorPlanToRemove != nil {
    // remove a major plan
    err = util.RemovePlan(&planList.MajorPlan, majorPlanToRemoveIndex)
    if err != nil {
      return err
    }
    goto saveFile
  }

  // remove a child plan
  if parentPlan == nil {
    return &util.ParentPlanNotFoundError{}
  }
  for index, child := range parentPlan.ChildrenPlan {
    if child.PlanId == targetPlanId {
      targetIndex = index
      break
    }
  }
  err = util.RemovePlan(&parentPlan.ChildrenPlan, targetIndex)
  if err != nil {
    return err
  }

saveFile:
  err = util.SaveAsJsonFile(util.DefaultFilePath, planList)
  if err != nil {
    return err
  }
  return nil
}
