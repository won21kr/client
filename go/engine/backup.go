// Backup creates backup keys for a user and pushes them to the server.
// It checks for existing backup devices and offers to revoke the
// keys.
//

package engine

import (
	"fmt"
	"strings"

	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/protocol/go"
)

// BackupKeypush is an engine.
type Backup struct {
	passphrase string
	libkb.Contextified
}

// NewBackupKeypush creates a BackupKeypush engine.
func NewBackup(g *libkb.GlobalContext) *Backup {
	return &Backup{
		Contextified: libkb.NewContextified(g),
	}
}

// Name is the unique engine name.
func (e *Backup) Name() string {
	return "BackupKeypush"
}

// GetPrereqs returns the engine prereqs.
func (e *Backup) Prereqs() Prereqs {
	return Prereqs{
		Session: true,
	}
}

// RequiredUIs returns the required UIs.
func (e *Backup) RequiredUIs() []libkb.UIKind {
	return []libkb.UIKind{
		libkb.LoginUIKind,
	}
}

// SubConsumers returns the other UI consumers for this engine.
func (e *Backup) SubConsumers() []libkb.UIConsumer {
	return []libkb.UIConsumer{
		&DetKeyEngine{},
		&RevokeEngine{},
	}
}

// Run starts the engine.
func (e *Backup) Run(ctx *Context) error {
	me, err := libkb.LoadMe(libkb.LoadUserArg{})
	if err != nil {
		return err
	}

	// check for existing backup keys
	cki := me.GetComputedKeyInfos()
	if cki == nil {
		return fmt.Errorf("no computed key infos")
	}
	var needReload bool
	for _, bdev := range cki.BackupDevices() {
		revoke, err := ctx.LoginUI.PromptRevokeBackupDeviceKeys(
			keybase1.PromptRevokeBackupDeviceKeysArg{
				Device: *bdev.ProtExport(),
			})
		if err != nil {
			e.G().Log.Warning("prompt error: %s", err)
			continue
		}
		if !revoke {
			continue
		}
		reng := NewRevokeDeviceEngine(RevokeDeviceEngineArgs{ID: bdev.ID}, e.G())
		if err := RunEngine(reng, ctx); err != nil {
			// probably not a good idea to continue...
			return err
		}
		needReload = true
	}

	if needReload {
		me, err = libkb.LoadMe(libkb.LoadUserArg{})
		if err != nil {
			return err
		}
	}

	signingKey, _, err := e.G().Keyrings.GetSecretKeyWithPrompt(ctx.LoginContext, libkb.SecretKeyArg{
		Me:      me,
		KeyType: libkb.DeviceSigningKeyType,
	}, ctx.SecretUI, "backup key signature")
	if err != nil {
		return err
	}

	words, err := libkb.SecWordList(libkb.BackupKeyPhraseEntropy)
	if err != nil {
		return err
	}
	e.passphrase = strings.Join(words, " ")

	kgarg := &BackupKeygenArg{
		Passphrase: e.passphrase,
		Me:         me,
		SigningKey: signingKey,
	}
	kgeng := NewBackupKeygen(kgarg, e.G())
	if err := RunEngine(kgeng, ctx); err != nil {
		return err
	}

	return ctx.LoginUI.DisplayBackupPhrase(keybase1.DisplayBackupPhraseArg{Phrase: e.passphrase})

}

func (e *Backup) Passphrase() string {
	return e.passphrase
}
