package utils

import (
     "fmt"
     "github.com/spf13/cobra"
)

func NewCLI(app string, desc string) {
    return &cobra.Command {
        Use: app,
        Short: desc,
        Run: func (cmd *cobra.Command, args []string) {
            fmt.Printf("Running %s\n", app)
        },
    }
}
