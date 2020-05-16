package command

const chanceOfSuccess int = 60

// Command is interface for realization command pattern
type Command interface {
	Execute() error
	Undo() error
}
