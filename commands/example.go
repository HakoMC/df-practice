package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

// ExampleCommand contains our command parameters.
type ExampleCommand struct {
	// Each parameter requires the field to be exported in order for the command to work properly.
	// The `cmd:""` struct tag may be specified to change the name of the parameter and its suffix.
	// Supported type: int*, uint*, float*, string, bool, mgl64.Vec3, []cmd.Target, cmd.Enum ...
	Number int `cmd:"number"`
	// Wrapping the parameter type in `cmd.Optional` makes the parameter optional, meaning a cmd.Source
	// does not need to supply it.
	OptionalMessage cmd.Optional[string]
}

// Run will be called when the player runs the command. In this case, we will print the number back to the player
func (c ExampleCommand) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	msg, ok := c.OptionalMessage.Load()
	output.Printf("%d %s (optional arg set? %v)", c.Number, msg, ok)
}
