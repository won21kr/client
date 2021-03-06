// Auto-generated by avdl-compiler v1.3.19 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_team.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type TeamChangeSet struct {
	MembershipChanged bool `codec:"membershipChanged" json:"membershipChanged"`
	KeyRotated        bool `codec:"keyRotated" json:"keyRotated"`
	Renamed           bool `codec:"renamed" json:"renamed"`
}

func (o TeamChangeSet) DeepCopy() TeamChangeSet {
	return TeamChangeSet{
		MembershipChanged: o.MembershipChanged,
		KeyRotated:        o.KeyRotated,
		Renamed:           o.Renamed,
	}
}

type TeamChangedArg struct {
	TeamID      TeamID        `codec:"teamID" json:"teamID"`
	TeamName    string        `codec:"teamName" json:"teamName"`
	LatestSeqno Seqno         `codec:"latestSeqno" json:"latestSeqno"`
	Changes     TeamChangeSet `codec:"changes" json:"changes"`
}

func (o TeamChangedArg) DeepCopy() TeamChangedArg {
	return TeamChangedArg{
		TeamID:      o.TeamID.DeepCopy(),
		TeamName:    o.TeamName,
		LatestSeqno: o.LatestSeqno.DeepCopy(),
		Changes:     o.Changes.DeepCopy(),
	}
}

type TeamDeletedArg struct {
	TeamID TeamID `codec:"teamID" json:"teamID"`
}

func (o TeamDeletedArg) DeepCopy() TeamDeletedArg {
	return TeamDeletedArg{
		TeamID: o.TeamID.DeepCopy(),
	}
}

type NotifyTeamInterface interface {
	TeamChanged(context.Context, TeamChangedArg) error
	TeamDeleted(context.Context, TeamID) error
}

func NotifyTeamProtocol(i NotifyTeamInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyTeam",
		Methods: map[string]rpc.ServeHandlerDescription{
			"teamChanged": {
				MakeArg: func() interface{} {
					ret := make([]TeamChangedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamChangedArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamChangedArg)(nil), args)
						return
					}
					err = i.TeamChanged(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"teamDeleted": {
				MakeArg: func() interface{} {
					ret := make([]TeamDeletedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamDeletedArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamDeletedArg)(nil), args)
						return
					}
					err = i.TeamDeleted(ctx, (*typedArgs)[0].TeamID)
					return
				},
				MethodType: rpc.MethodNotify,
			},
		},
	}
}

type NotifyTeamClient struct {
	Cli rpc.GenericClient
}

func (c NotifyTeamClient) TeamChanged(ctx context.Context, __arg TeamChangedArg) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifyTeam.teamChanged", []interface{}{__arg})
	return
}

func (c NotifyTeamClient) TeamDeleted(ctx context.Context, teamID TeamID) (err error) {
	__arg := TeamDeletedArg{TeamID: teamID}
	err = c.Cli.Notify(ctx, "keybase.1.NotifyTeam.teamDeleted", []interface{}{__arg})
	return
}
