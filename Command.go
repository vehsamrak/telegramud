package main

type GameCommand interface {
	Name() string
	Execute() *CommandResult
}
