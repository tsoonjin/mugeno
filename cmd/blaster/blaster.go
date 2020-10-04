package main

import (
    cli "github.com/tsoonjin/mugeno/pkg/utils"
)

func main() {
    var app = "blaster"
    var version = "0.0.1"
    var desc = "A simple DDOS tool"

    // Parse user arguments
    var cmd = cli.Create(app, version, desc)
    cmd.Parse(nil)
}
