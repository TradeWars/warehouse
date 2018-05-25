// Package util provides some miscellaneous utility functions that do not depend
// on any of functionality or structures of the application.
package util

// ErrSeq takes a sequence of function calls that all return errors and returns
// on the first non-nil error, if none are nil it returns nil
func ErrSeq(calls ...error) (err error) {
	for _, call := range calls {
		if call != nil {
			return call
		}
	}
	return nil
}
