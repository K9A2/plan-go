package util

type NotEnoughArgumentsError struct{}

func (e *NotEnoughArgumentsError) Error() string {
  return "Not enough arguments"
}

type UnrecognizedCommandError struct{}

func (e *UnrecognizedCommandError) Error() string {
  return "Unrecognized command"
}

type WrongArgumentNumberError struct{}

func (e *WrongArgumentNumberError) Error() string {
  return "Wrong argument number"
}

type PlanNotExistsError struct{}

func (e *PlanNotExistsError) Error() string {
  return "Plan not exits"
}

type IndexOutOfRangeError struct {}

func (e *IndexOutOfRangeError) Error() string {
  return "Index out of range"
}

type ParentPlanNotFoundError struct {}

func (e *ParentPlanNotFoundError) Error() string {
  return "Parent plan not found"
}
