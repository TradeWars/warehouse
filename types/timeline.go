package types

import (
	"time"
)

// Timeline declares an interface for an event manager that is responsible for
// storing and retreiving events.
type Timeline interface {
	Emit(timestamp time.Time, data []byte) (err error)
}
