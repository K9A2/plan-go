package command

import (
  "fmt"
  "github.com/stormlin/plan-go/util"
)

type addCommand struct {
  args []string
}

func NewAddCommand(args []string) *addCommand {
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
  if len(command.args) == 1 {
    // add a new plan with specified plan title
    title := command.args[0]
    newPlan := util.NewPlanItem(title)

    planList, err := util.ReadFromJsonFile(util.DefaultFilePath)
    if err != nil {
      return err
    }
    planList.MajorPlan = append(planList.MajorPlan, newPlan)
    err = util.SaveAsJsonFile(util.DefaultFilePath, planList)
    if err != nil {
      return err
    }
  }

  if len(command.args) == 3 && command.args[1] == "-p" {
    // add a new plan as a child of given parent plan
    planList, err := util.ReadFromJsonFile(util.DefaultFilePath)
    if err != nil {
      return err
    }
    title := command.args[0]
    targetPlanId := command.args[2]
    var parentPlan *util.PlanItem
    for _, root := range planList.MajorPlan {
      result := util.LocateParentPlan(root, targetPlanId)
      if result != nil {
        parentPlan = result
        break
      }
    }
    if parentPlan == nil {
      return &util.PlanNotExistsError{}
    }
    newPlan := util.NewPlanItem(title)
    parentPlan.ChildrenPlan = append(parentPlan.ChildrenPlan, newPlan)
    err = util.SaveAsJsonFile(util.DefaultFilePath, planList)
    if err != nil {
      return err
    }
    return nil 
  }

  return &util.WrongArgumentNumberError{}
}
