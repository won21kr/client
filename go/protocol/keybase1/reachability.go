// Auto-generated by avdl-compiler v1.3.16 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/reachability.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type Reachable int

const (
	Reachable_UNKNOWN Reachable = 0
	Reachable_YES     Reachable = 1
	Reachable_NO      Reachable = 2
)

func (o Reachable) DeepCopy() Reachable { return o }

var ReachableMap = map[string]Reachable{
	"UNKNOWN": 0,
	"YES":     1,
	"NO":      2,
}

var ReachableRevMap = map[Reachable]string{
	0: "UNKNOWN",
	1: "YES",
	2: "NO",
}

func (e Reachable) String() string {
	if v, ok := ReachableRevMap[e]; ok {
		return v
	}
	return ""
}

type Reachability struct {
	Reachable Reachable `codec:"reachable" json:"reachable"`
}

func (o Reachability) DeepCopy() Reachability {
	return Reachability{
		Reachable: o.Reachable.DeepCopy(),
	}
}

type ReachabilityChangedArg struct {
	Reachability Reachability `codec:"reachability" json:"reachability"`
}

func (o ReachabilityChangedArg) DeepCopy() ReachabilityChangedArg {
	return ReachabilityChangedArg{
		Reachability: o.Reachability.DeepCopy(),
	}
}

type StartReachabilityArg struct {
}

func (o StartReachabilityArg) DeepCopy() StartReachabilityArg {
	return StartReachabilityArg{}
}

type CheckReachabilityArg struct {
}

func (o CheckReachabilityArg) DeepCopy() CheckReachabilityArg {
	return CheckReachabilityArg{}
}

type ReachabilityInterface interface {
	ReachabilityChanged(context.Context, Reachability) error
	// Start reachability checks and return current status, which
	// may be cached.
	StartReachability(context.Context) (Reachability, error)
	// Performs a reachability check. This is not a cached response.
	CheckReachability(context.Context) (Reachability, error)
}

func ReachabilityProtocol(i ReachabilityInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.reachability",
		Methods: map[string]rpc.ServeHandlerDescription{
			"reachabilityChanged": {
				MakeArg: func() interface{} {
					ret := make([]ReachabilityChangedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ReachabilityChangedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ReachabilityChangedArg)(nil), args)
						return
					}
					err = i.ReachabilityChanged(ctx, (*typedArgs)[0].Reachability)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"startReachability": {
				MakeArg: func() interface{} {
					ret := make([]StartReachabilityArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.StartReachability(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"checkReachability": {
				MakeArg: func() interface{} {
					ret := make([]CheckReachabilityArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.CheckReachability(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ReachabilityClient struct {
	Cli rpc.GenericClient
}

func (c ReachabilityClient) ReachabilityChanged(ctx context.Context, reachability Reachability) (err error) {
	__arg := ReachabilityChangedArg{Reachability: reachability}
	err = c.Cli.Notify(ctx, "keybase.1.reachability.reachabilityChanged", []interface{}{__arg})
	return
}

// Start reachability checks and return current status, which
// may be cached.
func (c ReachabilityClient) StartReachability(ctx context.Context) (res Reachability, err error) {
	err = c.Cli.Call(ctx, "keybase.1.reachability.startReachability", []interface{}{StartReachabilityArg{}}, &res)
	return
}

// Performs a reachability check. This is not a cached response.
func (c ReachabilityClient) CheckReachability(ctx context.Context) (res Reachability, err error) {
	err = c.Cli.Call(ctx, "keybase.1.reachability.checkReachability", []interface{}{CheckReachabilityArg{}}, &res)
	return
}
