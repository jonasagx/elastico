package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"elastico/src/lib/commands"

	"github.com/elastic/go-elasticsearch/v7"
)

/*
	Author: Jonas Galvão Xavier
	Elastico exists to make the interaction with Elasticsearch's API more pleasant. The word elastico means elastic in Portuguese, in case you were wondering.

	The very first thing to happen is a connection, otherwise the user must set the host address.

	elastico <HOST:optional> <COMMAND> <OPTIONS>

	The host is optional, the host is the last host used

	Commands:
		* Ping - looks for an instance of ES locally
		* Help <topic> - shows help
*/

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("please provide a command, try elastico help")
		return
	}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Printf("cannot create client with: %v\n", err)
		return
	}

	to, _ := context.WithTimeout(context.Background(), time.Minute)

	ctx := context.WithValue(to, "client", es)
	cmd, ok := commands.Get(args[0], args[1:]...)
	if !ok {
		fmt.Println("command not found, try help this time")
		return
	}
	err = cmd(ctx, args[1:]...)
	if err != nil {
		fmt.Printf("failed running command with: %v", err)
	}
}
