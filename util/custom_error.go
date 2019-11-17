package util

type NotEnoughArgumentsError struct{}

func (e *NotEnoughArgumentsError) Error() string {
  return "Not enough arguments"
}

type UnrecognizedCommandError struct{}

func (e *UnrecognizedCommandError) Error() string {
  return "Unrecognized command"
}
