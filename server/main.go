package main

import (
	"github.com/jawher/mow.cli"
	"os"
)

func main() {
	httpserv := cli.App("server", "server app")

	httpserv.Command("run", "raise server in 127.0.0.1:8000", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			httpserver()
		}
	})

	httpserv.Run(os.Args)
}
