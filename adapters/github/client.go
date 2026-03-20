package github

import "context"

// Installation represents a GitHub App installation.
type Installation struct {
	ID      int64
	Account string
}

// Client is the interface for interacting with the GitHub API.
type Client interface {
	InstallationService
}

// InstallationService provides methods for managing GitHub App installations.
type InstallationService interface {
	ListInstallations(ctx context.Context) ([]*Installation, error)
}
