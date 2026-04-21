// Package container provides support for starting and stoping containers
// for running tests.
//
// The container engine depends on the host OS.
package container

// Container tracks information about the container started for tests.
type Container struct {
	ID   string
	Host string // IP:Port
}
