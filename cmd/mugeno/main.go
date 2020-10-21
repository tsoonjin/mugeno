package main

import (
    "fmt"
    "github.com/thatisuday/commando"
    "github.com/tsoonjin/mugeno/pkg/load"
)

func setupCLI(appName string, version string, desc string) *commando.CommandRegistry {

    // Set main app
    var cmd = commando.
                SetExecutableName(appName).
                SetVersion(version).
                SetDescription(desc)
    return cmd
}

func setupLoadTest(cli *commando.CommandRegistry) {
    test := load.LoadTest{
        Target: "https://de-awani-web-portal-dev.eco.astro.com.my/foto-terkini",
        Concurrency: 100,
        Duration: 18000,
    }

    cli.
        Register("load").
        SetDescription("Runs a load test on designated url").
        AddArgument("url", "target url", "").
        SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
            fmt.Printf("Arguments: % v", args)
            fmt.Printf("Flags: % v", flags)
            test.Start()

        })
}

func main() {
    var app = "mugeno"
    var version = "0.0.1"
    var desc = "Developer swiss army knife"

    cli := setupCLI(app, version, desc);
    setupLoadTest(cli);
    cli.Parse(nil)
}
