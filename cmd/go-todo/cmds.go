package main

import (
	"github.com/rikusen0335/go-todo/internal/server/application/api"

	"github.com/spf13/cobra"
)

func Cmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "MAINCMD of http server cmds",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "SUBCMD to run http server",
		Run: func(_ *cobra.Command, _ []string) {
			api.Run()
		},
	})

	return cmd
}