package command

import (
  "fmt"
  "github.com/stormlin/plan-go/util"
)

type statusCommand struct {
}

func NewStatusCommand() *statusCommand {
  return &statusCommand{}
}

func (command statusCommand) usage() {
  fmt.Println("Use <status> command to get status of all stored plans.")
  fmt.Println("")
  fmt.Println("usage: plan status")
}

func (command statusCommand) Execute() error {
  planList, err := util.ReadFromJsonFile(util.DefaultFilePath)
  if err != nil {
    return err
  }
  if len(planList.MajorPlan) == 0 {
    fmt.Println("There are no plans yet.")
    return nil
  }

  util.PrintPlanSlice(&planList.MajorPlan, "", 0)
  return nil
}
