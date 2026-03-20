package github

import (
	"context"

	"github.com/google/go-github/github"
)

type Client interface {
	InstallationService
}

type InstallationService interface {
	ListInstallations(ctx context.Context) ([]*github.Installation, error)
}
