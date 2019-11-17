package util

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

func Usage() {
  fmt.Println("plan-go is a simple and stupid plan management utility tool.")
  fmt.Println("")
  fmt.Println("usage: plan [--version] [--help | -h] [and other command (see below)]")
  fmt.Println("    add: add new [major | child] plan")
  fmt.Println("      [plan add 'plan_title'] or [plan add 'plan_title' -p <parent_plan_id>]")
  fmt.Println("    done: mark a plan (and its children) as done")
  fmt.Println("      [plan done <plan_id>]")
  fmt.Println("    status: check the status of plans")
  fmt.Println("      [plan status]")
  fmt.Println("    reopen: reopen a done plan (and its children)")
  fmt.Println("      [plan reopen <plan_id>")
  fmt.Println("    delete: delete a plan (and its children)")
  fmt.Println("      [plan delete <plan_id>]")
  fmt.Println("    retitle: retitle an existing plan")
  fmt.Println("      [plan retitle <plan_id> 'a_new_title']")
}

// Entry for each plan item
type PlanItem struct {
  PlanId       string     `json:"plan_id"`
  Date         string     `json:"date"`
  Title        string     `json:"title"`
  ChildrenPlan []PlanItem `json:"children_plan"`
}

// Root entry for plan list
type PlanList struct {
  MajorPlan []PlanItem `json:"major_plan"`
}

// Parse PlanList object from json file
func ReadFromJsonFile(filePath string) (*PlanList, error) {
  jsonFile, err := os.Open(filePath)
  defer jsonFile.Close()
  if err != nil {
    return nil, err
  }

  byteValue, err := ioutil.ReadAll(jsonFile)
  if err != nil {
    return nil, err
  }

  var planList PlanList
  err = json.Unmarshal(byteValue, &planList)
  if err != nil {
    return nil, err
  }

  return &planList, nil
}

func SaveAsJsonFile(filePath string, planList *PlanList) error {
  file, err := json.MarshalIndent(planList, "", "  ")
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(filePath, file, 0644)
  return err
}