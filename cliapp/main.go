package main

import (
    "os"
    "github.com/urfave/cli"
    )

func main() {
    app := cli.NewApp()
    app.Name = "Sample Cli app"
    app.Usage = "A sample cli App"
    app.Action = func(c *cli.Context) {
        println("Hello World from Cli app")
    }
    app.Run(os.Args)
}
