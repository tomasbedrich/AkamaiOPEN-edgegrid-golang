// Package storage provides access to the Akamai NetStorage V1 API
package storage

import (
	"errors"
	"net/http"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/session"
)

var (
	// ErrStructValidation is returned returned when given struct validation failed
	ErrStructValidation = errors.New("struct validation")
)

type (
	// NetStorage is the storage api interface
	NetStorage interface {
		StorageGroups
	}

	storage struct {
		session.Session
	}

	// Option defines a NetStorage option
	Option func(*storage)

	// ClientFunc is a storage client new method, this can used for mocking
	ClientFunc func(sess session.Session, opts ...Option) NetStorage
)

// Client returns a new dns Client instance with the specified controller
func Client(sess session.Session, opts ...Option) NetStorage {
	p := &storage{
		Session: sess,
	}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Exec overrides the session.Exec to add dns options
func (p *storage) Exec(r *http.Request, out interface{}, in ...interface{}) (*http.Response, error) {

	return p.Session.Exec(r, out, in...)
}
