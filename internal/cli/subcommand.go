package cli

import (
	"fmt"
)

// SubcommandSet is a set of subcommands
type SubcommandSet map[string]func()

// ValidCommands returns the set of valid subcommands for the given set
func (s SubcommandSet) ValidCommands() string {
	cmds := ""

	for key := range s {
		cmds += key + " "
	}

	return cmds
}

// Run runs the given subcommand
func (s SubcommandSet) Run(key string) error {
	cmd, ok := s[key]
	if !ok {
		return fmt.Errorf("Invalid subcommand '%v'. Valid subcommands: %v", key, s.ValidCommands())
	}

	cmd()
	return nil
}
