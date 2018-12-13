package main

import (
	
)

var orgsCommands commander

func init() {
	// flagSet := flag.NewFlagSet("orgs", flag.ExitOnError)
	// handler := func(args []string) error {
	// 	orgsCommands.run(flagSet, "src orgs", usage, args)
	// 	return nil
	// }

	// Register the command.
	// commands = append(commands, &command{
	// 	flagSet: flagSet,
	// 	aliases: []string{"org"},
	// 	handler: handler,
	// 	usageFunc: func() {
	// 		fmt.Println(usage)
	// 	},
	// })
}

const orgFragment = `
fragment OrgFields on Org {
    id
    name
    displayName
    members {
        nodes {
			id
			username
		}
    }
}
`

type Org struct {
	ID          string
	Name        string
	DisplayName string
	Members     struct {
		Nodes []User
	}
}
