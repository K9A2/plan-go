package command

import (
  "fmt"
  "testing"
)

func TestNewAddCommand(t *testing.T) {
 args := []string{
   "test_plan_title",
   "-p",
   "parent_plan_id",
 }
  add := NewAddCommand(args)
  add.usage()
  fmt.Println(add.args)
}

func TestNewStatusCommand(t *testing.T) {
  status := NewStatusCommand()
  status.usage()
}

func TestNewDoneCommand(t *testing.T) {
  args := []string{
    "plan_id",
  }
  done := NewDoneCommand(args)
  done.usage()
}

func TestNewReopenCommand(t *testing.T) {
  args := []string{
    "plan_id",
  }
  reopen := NewReopenCommand(args)
  reopen.usage()
}

func TestNewDeleteCommand(t *testing.T) {
  args := []string{
    "plan_id",
  }
  deleteCommand := NewDeleteCommand(args)
  deleteCommand.usage()
}

func TestNewRetitleCommand(t *testing.T) {
  args := []string{
    "plan_id",
  }
  retitle := NewRetitleCommand(args)
  retitle.usage()
}
