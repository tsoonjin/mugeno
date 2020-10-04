package cli

import (
     "github.com/thatisuday/commando"
)

type CLIConfig struct {
    arguments [][]string
}

func Create(name string, version string, desc string) *commando.CommandRegistry {
    var app = commando.
        SetExecutableName(name).
        SetVersion(version).
        SetDescription(desc)
    return app
}
