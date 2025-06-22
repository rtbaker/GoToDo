package http

import (
	"net/http"
	"time"
)

// Variables from the main config we want to pass to the
// http session manager, a subset of what scs session manager accepts

type SessionConfig struct {
	// IdleTimeout controls the maximum length of time a session can be inactive
	// before it expires. For example, some applications may wish to set this so
	// there is a timeout after 20 minutes of inactivity. By default IdleTimeout
	// is not set and there is no inactivity timeout.
	IdleTimeout time.Duration

	// Lifetime controls the maximum length of time that a session is valid for
	// before it expires. The lifetime is an 'absolute expiry' which is set when
	// the session is first created and does not change. The default value is 24
	// hours.
	Lifetime time.Duration

	// Name of the session cookie
	Name string

	SameSite http.SameSite
	Secure   bool
}
