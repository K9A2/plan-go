package util

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "time"
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
  PlanId       string      `json:"plan_id"`
  Date         string      `json:"date"`
  Title        string      `json:"title"`
  ChildrenPlan []*PlanItem `json:"children_plan"`
}

// Root entry for plan list
type PlanList struct {
  MajorPlan []*PlanItem `json:"major_plan"`
}

func NewPlanItem(title string) *PlanItem {
  planIdString := fmt.Sprintf("%x", time.Now().UnixNano())
  return &PlanItem{
    PlanId:       planIdString[len(planIdString)-8:],
    Date:         time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006"),
    Title:        title,
    ChildrenPlan: nil,
  }
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

func PrintPlanSlice(planSlice *[]*PlanItem, parentIndex string, indentCount int) {
  // prepare prefix
  prefix := ""
  for i := 0; i < indentCount; i++ {
    prefix = prefix + DefaultIndent
  }
  // print the plans recursively
  for i := 0; i < len(*planSlice); i++ {
    item := (*planSlice)[i]
    index := ""
    if parentIndex != "" {
      index = parentIndex + fmt.Sprintf(".%d", i+1)
    } else {
      index = parentIndex + fmt.Sprintf("%d", i+1)
    }
    fmt.Printf("%s%s plan id:<%s>, date: <%s>\n", prefix, index, item.PlanId, item.Date)
    fmt.Printf("%s%s\n", prefix+DefaultIndent, item.Title)
    PrintPlanSlice(&item.ChildrenPlan, index, indentCount+1)
  }
}

func LocateParentPlan(currentPlan *PlanItem, planId string) *PlanItem {
  if currentPlan == nil {
    return nil
  }
  if currentPlan.PlanId == planId {
    return currentPlan
  }
  for _, child := range currentPlan.ChildrenPlan {
    result := LocateParentPlan(child, planId)
    if result != nil {
      return result
    }
  }
  return nil
}
