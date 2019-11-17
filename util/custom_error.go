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
