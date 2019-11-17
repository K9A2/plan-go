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
    result := FindPlan(root, targetPlanId)
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

func TestMarkAsDone(t *testing.T) {
  planList, err := ReadFromJsonFile(DefaultFilePath)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  var targetPlan *PlanItem
  for _, plan := range planList.MajorPlan {
    result := FindPlan(plan, "9ae5f9f2")
    if result != nil {
      targetPlan = result
      break
    }
  }
  if targetPlan == nil {
    fmt.Println("plan not exists")
    return
  }

  // mark it and its children as done
  MarkAsDone(targetPlan)
  err = SaveAsJsonFile(DefaultFilePath, planList)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  return
}

func TestFindParent(t *testing.T) {
  planList, err := ReadFromJsonFile(DefaultFilePath)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  targetPlanId := "9ae5f9f2"
  var parentPlan *PlanItem
  var majorPlanToRemove *PlanItem
  var majorPlanToRemoveIndex int
  var targetIndex = -1
  for index, major := range planList.MajorPlan {
    if major.PlanId == targetPlanId {
      // remove this major plan
      majorPlanToRemove = major
      majorPlanToRemoveIndex = index
      break
    }
    result := FindParent(major, targetPlanId)
    if result != nil {
      parentPlan = result
    }
  }
  if majorPlanToRemove != nil {
    // remove a major plan
    err = RemovePlan(&planList.MajorPlan, majorPlanToRemoveIndex)
    if err != nil {
      fmt.Printf("err in removing a major plan, err: <%s>\n", err.Error())
      return
    }
    goto saveFile
  }

  // remove a child plan
  if parentPlan == nil {
    fmt.Print("parent plan not found")
    return
  }
  for index, child := range parentPlan.ChildrenPlan {
    if child.PlanId == targetPlanId {
      targetIndex = index
      break
    }
  }
  if targetIndex == -1 {
    fmt.Println("error in locating child plan")
    return
  }
  err = RemovePlan(&parentPlan.ChildrenPlan, targetIndex)
  if err != nil {
    fmt.Printf("error in removing plan")
    return
  }

saveFile:
  err = SaveAsJsonFile(DefaultFilePath, planList)
  if err != nil {
    fmt.Print("error in saving json file")
    return
  }
}
