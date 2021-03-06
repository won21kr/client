// Copyright 2017 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"context"
	"fmt"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

type CmdTeamRemoveMember struct {
	libkb.Contextified
	Team     string
	Username string
	Force    bool
}

func NewCmdTeamRemoveMemberRunner(g *libkb.GlobalContext) *CmdTeamRemoveMember {
	return &CmdTeamRemoveMember{Contextified: libkb.NewContextified(g)}
}

func newCmdTeamRemoveMember(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name:         "remove-member",
		ArgumentHelp: "<team name> --user=<username>",
		Usage:        "Remove a user from a team.",
		Action: func(c *cli.Context) {
			cmd := &CmdTeamRemoveMember{Contextified: libkb.NewContextified(g)}
			cl.ChooseCommand(cmd, "remove-member", c)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "u, user",
				Usage: "username",
			},
			cli.BoolFlag{
				Name:  "f, force",
				Usage: "bypass warnings",
			},
		},
	}
}

func (c *CmdTeamRemoveMember) ParseArgv(ctx *cli.Context) error {
	var err error
	c.Team, err = ParseOneTeamName(ctx)
	if err != nil {
		return err
	}

	c.Username, err = ParseUser(ctx)
	if err != nil {
		return err
	}

	c.Force = ctx.Bool("force")

	return nil
}

func (c *CmdTeamRemoveMember) Run() error {
	ui := c.G().UI.GetTerminalUI()

	if !c.Force {
		configCLI, err := GetConfigClient(c.G())
		if err != nil {
			return err
		}
		curStatus, err := configCLI.GetCurrentStatus(context.TODO(), 0)
		if err != nil {
			return err
		}

		var prompt string
		if curStatus.User != nil && libkb.NewNormalizedUsername(c.Username).Eq(libkb.NewNormalizedUsername(curStatus.User.Username)) {
			prompt = fmt.Sprintf("Are you sure you want to remove yourself from team %s?", c.Team)
		} else {
			prompt = fmt.Sprint("Are you sure?")
		}
		proceed, err := ui.PromptYesNo(PromptDescriptorRemoveMember, prompt, libkb.PromptDefaultNo)
		if err != nil {
			return err
		}
		if !proceed {
			return nil
		}
	}

	cli, err := GetTeamsClient(c.G())
	if err != nil {
		return err
	}

	arg := keybase1.TeamRemoveMemberArg{
		Name:     c.Team,
		Username: c.Username,
	}

	if err = cli.TeamRemoveMember(context.Background(), arg); err != nil {
		return err
	}

	ui.Printf("Success! %s removed from team %s.\n", c.Username, c.Team)

	return nil
}

func (c *CmdTeamRemoveMember) GetUsage() libkb.Usage {
	return libkb.Usage{
		Config:    true,
		API:       true,
		KbKeyring: true,
	}
}
