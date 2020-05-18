package factorymethod

import (
	"github.com/losev-al/GoLessons/pkg/command"
)

// Creator is interface for realization factory method pattern
type Creator interface {
	Create(from string, to string) command.Command
}
