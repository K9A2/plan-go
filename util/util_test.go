package util

import (
  "fmt"
  "testing"
)

func TestUsage(t *testing.T) {
  Usage()
}

func TestSaveAsJsonFile(t *testing.T) {
  planList := &PlanList{
    MajorPlan: []PlanItem{
      {
        PlanId:       "plan-1",
        Date:         "2018-01-01",
        Title:        "title-1",
        ChildrenPlan: nil,
      },
      {
        PlanId:       "plan-2",
        Date:         "2019-01-01",
        Title:        "title-2",
        ChildrenPlan: []PlanItem{
          {
            PlanId: "child-1",
            Date:   "child-1-date",
            Title:  "child-1-title",
            ChildrenPlan: nil,
          },
          {
            PlanId: "child-2",
            Date:   "child-2-date",
            Title:  "child-2-title",
            ChildrenPlan: nil,
          },
        },
      },
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
