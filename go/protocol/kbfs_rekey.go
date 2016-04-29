// Auto-generated by avdl-compiler v1.3.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/kbfs_rekey.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type NeedsRekeyForCurrentDeviceArg struct {
	Folder Folder `codec:"folder" json:"folder"`
}

// KbfsRekeyInterface is a protocol that lets the service ask KBFS about the
// rekey status of a folder.
type KbfsRekeyInterface interface {
	// NeedsRekeyForCurrentDevice returns true if the current device should be
	// able to read a folder, but cannot because it cannot decrypt the secret key.
	// For public folders, the return value is always false.  Returns an error if
	// there is no current session.  Note that this call may block on other calls
	// back to the service, and so no locks should be held while calling it.
	NeedsRekeyForCurrentDevice(context.Context, Folder) (bool, error)
}

func KbfsRekeyProtocol(i KbfsRekeyInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.KbfsRekey",
		Methods: map[string]rpc.ServeHandlerDescription{
			"NeedsRekeyForCurrentDevice": {
				MakeArg: func() interface{} {
					ret := make([]NeedsRekeyForCurrentDeviceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NeedsRekeyForCurrentDeviceArg)
					if !ok {
						err = rpc.NewTypeError((*[]NeedsRekeyForCurrentDeviceArg)(nil), args)
						return
					}
					ret, err = i.NeedsRekeyForCurrentDevice(ctx, (*typedArgs)[0].Folder)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type KbfsRekeyClient struct {
	Cli rpc.GenericClient
}

// NeedsRekeyForCurrentDevice returns true if the current device should be
// able to read a folder, but cannot because it cannot decrypt the secret key.
// For public folders, the return value is always false.  Returns an error if
// there is no current session.  Note that this call may block on other calls
// back to the service, and so no locks should be held while calling it.
func (c KbfsRekeyClient) NeedsRekeyForCurrentDevice(ctx context.Context, folder Folder) (res bool, err error) {
	__arg := NeedsRekeyForCurrentDeviceArg{Folder: folder}
	err = c.Cli.Call(ctx, "keybase.1.KbfsRekey.NeedsRekeyForCurrentDevice", []interface{}{__arg}, &res)
	return
}