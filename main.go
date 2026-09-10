package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Syink/erlc-cli/internal/erlc"
)

func main() {
	key := os.Getenv("ERLC_SERVER_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "set ERLC_SERVER_KEY")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: erlc-cli <status|players|command \"...\">")
		os.Exit(1)
	}

	client := erlc.New(key)
	ctx := context.Background()

	switch os.Args[1] {
	case "status":
		info, err := client.GetServer(ctx, erlc.Include{})
		check(err)
		fmt.Printf("%s — %d/%d players\n", info.Name, info.CurrentPlayers, info.MaxPlayers)

	case "players":
		info, err := client.GetServer(ctx, erlc.Include{Players: true})
		check(err)
		for _, p := range info.Players {
			fmt.Printf("%-25s %-10s %s\n", p.Player, p.Team, p.Permission)
		}

	case "command":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: erlc-cli command \":h message\"")
			os.Exit(1)
		}
		msg, err := client.RunCommand(ctx, os.Args[2])
		check(err)
		fmt.Println(msg)

	default:
		fmt.Fprintln(os.Stderr, "unknown subcommand:", os.Args[1])
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}