package types

import (
	"time"
)

// Eventer declares an interface for storing and retreiving events
type Eventer interface {
	Emit(timestamp time.Time, event interface{}) (err error)
}
