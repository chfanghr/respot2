package core

import (
	"github.com/chfanghr/respot2/protocol/spotify"
	"sync"
)

type Credential struct {
	lock     sync.Locker
	username string
	authType spotify.AuthenticationType
	authData []byte
}

func NewCredentialWithPassword(username, password string) (cre *Credential) {
	return &Credential{
		lock:     new(sync.Mutex),
		username: username,
		authType: spotify.AuthenticationType_AUTHENTICATION_USER_PASS,
		authData: []byte(password),
	}
}
