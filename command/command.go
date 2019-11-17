package command

type I interface {
  Execute() error
  usage()
}
