package quicclients

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Client returns the path to the QUIC client
func Client() string {
	return execPath("client")
}

// Server returns the path to the QUIC server
func Server() string {
	return execPath("server")
}

func execPath(role string) string {
	_, thisfile, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current path")
	}
	return filepath.Join(thisfile, fmt.Sprintf("../%s-%s-debug", role, runtime.GOOS))
}
