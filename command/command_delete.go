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

  // going to remove a major plan
  var majorPlanToRemove *util.PlanItem
  var majorPlanToRemoveIndex int

  // going to remove a child plan
  var parentPlan *util.PlanItem
  var childIndex = util.ElementNotFoundIndex

  // check major plan first
  for index, major := range planList.MajorPlan {
    if major.PlanId == targetPlanId {
      // remove this major plan
      majorPlanToRemove = major
      majorPlanToRemoveIndex = index
      break
    }
    // target may be one of its child
    parentPlan, childIndex = util.FindParent(major, targetPlanId)
    if parentPlan != nil {
      // we found its parent!
      break
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
  err = util.RemovePlan(&parentPlan.ChildrenPlan, childIndex)
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
