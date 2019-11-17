package util

import (
  "fmt"
  "testing"
  "time"
)

func TestUsage(t *testing.T) {
  Usage()
}

func TestSaveAsJsonFile(t *testing.T) {
  majorPlan1 := NewPlanItem("major-plan-1")
  time.Sleep(1 * time.Millisecond)
  majorPlan2 := NewPlanItem("major-plan-2")

  majorPlan2.ChildrenPlan = append(
    majorPlan2.ChildrenPlan, NewPlanItem("child-plan-1"))
  majorPlan2.ChildrenPlan = append(
    majorPlan2.ChildrenPlan, NewPlanItem("child-plan-2"))

  planList := &PlanList{
    MajorPlan: []*PlanItem{
      majorPlan1,
      majorPlan2,
    },
  }

  err := SaveAsJsonFile(DefaultFilePath, planList)
  if err != nil {
    fmt.Println(err.Error())
  }
}

func TestPrintPlanSlice(t *testing.T) {
  planList, err := ReadFromJsonFile(DefaultFilePath)
  if err != nil {
    fmt.Print(err.Error())
    return
  }
  PrintPlanSlice(&planList.MajorPlan, "", 0)
}

func TestLocateParentPlan(t *testing.T) {
  planList, err := ReadFromJsonFile(DefaultFilePath)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  title := "a new plan"
  targetPlanId := "f69fb1aa"
  var parentPlan *PlanItem
  for _, root := range planList.MajorPlan {
    result := LocateParentPlan(root, targetPlanId)
    if result != nil {
      parentPlan = result
      break
    }
  }
  if parentPlan == nil {
    fmt.Println("Plan not exists")
    return
  }
  newPlan := NewPlanItem(title)
  parentPlan.ChildrenPlan = append(parentPlan.ChildrenPlan, newPlan)
  err = SaveAsJsonFile(DefaultFilePath, planList)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
}
