// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"

	sqlite "github.com/glebarez/go-sqlite"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
)

// Enable SQLite foreign key support.
const dsnParams string = "?_pragma=foreign_keys(1)"

type (
	config struct {
		client *ent.Client
		ctx    context.Context
		debug  bool
	}

	// Backend implements the protobom.pkg.storage.Backend interface.
	Backend[T storage.ProtobomType] struct {
		config
		storage.Storer[T]
		storage.Retriever[T]
		Options BackendOptions[T]
	}

	// BackendOptions contains options specific to the protobom ent backend.
	BackendOptions[T storage.ProtobomType] struct {
		DatabaseFile string
	}

	// Option represents a single configuration option for the ent backend.
	Option[T storage.ProtobomType] func(*Backend[T])
)

var errInvalidEntOptions = errors.New("invalid ent backend options")

func NewBackend[T storage.ProtobomType](opts ...Option[T]) *Backend[T] {
	backend := &Backend[T]{
		Options: NewBackendOptions[T](),
	}

	for _, opt := range opts {
		opt(backend)
	}

	if err := backend.ClientSetup(); err != nil {
		panic(err)
	}

	return backend
}

func NewBackendOptions[T storage.ProtobomType]() BackendOptions[T] {
	return BackendOptions[T]{
		DatabaseFile: ":memory:",
	}
}

func WithDatabaseFile[T storage.ProtobomType](file string) Option[T] {
	return func(backend *Backend[T]) {
		backend.WithDatabaseFile(file)
	}
}
func WithBackendOptions[T storage.ProtobomType](opt BackendOptions[T]) Option[T] {
	return func(backend *Backend[T]) {
		backend.WithBackendOptions(opt)
	}
}

func (backend *Backend[any]) WithDatabaseFile(file string) *Backend[any] {
	backend.Options.DatabaseFile = file

	return backend
}

func (backend *Backend[any]) WithBackendOptions(opt BackendOptions[any]) *Backend[any] {
	backend.Options = opt
	return backend
}

func Debug[T storage.ProtobomType]() Option[T] {
	return func(backend *Backend[T]) {
		backend.Debug()
	}
}

func (backend *Backend[any]) Debug() *Backend[any] {
	backend.debug = true
	backend.client = backend.client.Debug()

	return backend
}

func (backend *Backend[any]) ClientSetup() error {
	// Register the SQLite driver as "sqlite3".
	if !slices.Contains(sql.Drivers(), "sqlite3") {
		sqlite.RegisterAsSQLITE3()
	}

	clientOpts := []ent.Option{}
	if backend.debug {
		clientOpts = append(clientOpts, ent.Debug())
	}

	client, err := ent.Open("sqlite3", fmt.Sprintf("%s%s", backend.Options.DatabaseFile, dsnParams), clientOpts...)
	if err != nil {
		return fmt.Errorf("failed opening connection to sqlite: %w", err)
	}

	backend.client = client

	backend.ctx = ent.NewContext(context.Background(), client)

	// Run the auto migration tool.
	if err := backend.client.Schema.Create(backend.ctx); err != nil {
		return fmt.Errorf("failed creating schema resources: %w", err)
	}

	return nil
}
