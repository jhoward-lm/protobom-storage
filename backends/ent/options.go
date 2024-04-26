// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

type (
	// Backend implements the protobom model.v1.storage.Backend interface.
	Backend struct {
		Options BackendOptions
	}

	// BackendOptions contains options specific to the protobom ent backend.
	BackendOptions struct {
		DatabaseFile string
	}

	// Option represents a single configuration option for the ent backend.
	Option func(*Backend)
)

func NewBackend(opts ...Option) *Backend {
	backend := &Backend{
		Options: NewBackendOptions(),
	}

	for _, opt := range opts {
		opt(backend)
	}

	return backend
}

func NewBackendOptions() BackendOptions {
	return BackendOptions{
		DatabaseFile: ":memory:",
	}
}

func WithDatabaseFile(file string) Option {
	return func(backend *Backend) {
		backend.WithDatabaseFile(file)
	}
}

func (backend *Backend) WithDatabaseFile(file string) *Backend {
	if backend == nil {
		backend = NewBackend()
	}

	backend.Options.DatabaseFile = file

	return backend
}
