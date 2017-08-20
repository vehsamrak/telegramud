package commands

type Command interface {
	Execute() string
	GetNames() []string
}
