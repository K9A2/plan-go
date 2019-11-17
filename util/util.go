package util

import (
  "fmt"
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
