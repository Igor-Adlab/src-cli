package main

import (
	
)

var usersCommands commander

func init() {
	// flagSet := flag.NewFlagSet("users", flag.ExitOnError)
	// handler := func(args []string) error {
	// 	usersCommands.run(flagSet, "src users", usage, args)
	// 	return nil
	// }

	// Register the command.
	// commands = append(commands, &command{
	// 	flagSet: flagSet,
	// 	aliases: []string{"user"},
	// 	handler: handler,
	// 	usageFunc: func() {
	// 		fmt.Println(usage)
	// 	},
	// })
}

const userFragment = `
fragment UserFields on User {
    id
    username
    displayName
    siteAdmin
    organizations {
		nodes {
        	id
        	name
        	displayName
		}
    }
    emails {
        email
        verified
    }
    url
}
`

type User struct {
	ID            string
	Username      string
	DisplayName   string
	SiteAdmin     bool
	Organizations struct {
		Nodes []Org
	}
	Emails []UserEmail
	URL    string
}

type UserEmail struct {
	Email    string
	Verified bool
}
