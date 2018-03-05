//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package service

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"golang.org/x/sys/unix"
	"k8s.io/client-go/kubernetes"

	"github.com/arangodb/k8s-operator/pkg/storage/provisioner"
)

// Config for the storage provisioner
type Config struct {
	Address        string // Server address to listen on
	NodeName       string
	Namespace      string
	ServiceAccount string
	LocalPath      []string
}

// Dependencies for the storage provisioner
type Dependencies struct {
	Log     zerolog.Logger
	KubeCli kubernetes.Interface
}

// Provisioner implements a Local storage provisioner
type Provisioner struct {
	Config
	Dependencies
}

// New creates a new local storage provisioner
func New(config Config, deps Dependencies) (*Provisioner, error) {
	return &Provisioner{
		Config:       config,
		Dependencies: deps,
	}, nil
}

// Run the provisioner until the given context is canceled.
func (p *Provisioner) Run(ctx context.Context) {
	runServer(ctx, p.Log, p.Address, p)
}

// GetFSInfo fetches information from the filesystem containing
// the given local path.
func (p *Provisioner) GetFSInfo(ctx context.Context, localPath string) (provisioner.FSInfo, error) {
	statfs := &unix.Statfs_t{}
	if err := unix.Statfs(localPath, statfs); err != nil {
		return provisioner.FSInfo{}, maskAny(err)
	}

	// Available is blocks available * fragment size
	available := int64(statfs.Bavail) * int64(statfs.Bsize)

	// Capacity is total block count * fragment size
	capacity := int64(statfs.Blocks) * int64(statfs.Bsize)

	return provisioner.FSInfo{
		Available: available,
		Capacity:  capacity,
	}, nil
}

// Prepare a volume at the given local path
func (p *Provisioner) Prepare(ctx context.Context, localPath string) error {
	// Make sure directory is empty
	if err := os.RemoveAll(localPath); err != nil && !os.IsNotExist(err) {
		return maskAny(err)
	}
	// Make sure directory exists
	if err := os.MkdirAll(localPath, 0755); err != nil {
		return maskAny(err)
	}
	// Set access rights
	if err := os.Chmod(localPath, 0777); err != nil {
		return maskAny(err)
	}
	return nil
}

// Remove a volume with the given local path
func (p *Provisioner) Remove(ctx context.Context, localPath string) error {
	// Make sure directory is empty
	if err := os.RemoveAll(localPath); err != nil && !os.IsNotExist(err) {
		return maskAny(err)
	}
	return nil
}
