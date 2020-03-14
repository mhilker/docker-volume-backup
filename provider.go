package dockervolumebackup

import (
	"context"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

// Provider contains methods to find volumes to backup
type Provider struct {
	client ProviderClient
}

// ProviderClient provides a list of Volumes to backup
type ProviderClient interface {
	VolumeList(ctx context.Context, filter filters.Args) (volume.VolumesListOKBody, error)
}

// Volume represents a docker volume to backup
type Volume struct {
	Name string
	Path string
}

// NewDockerProviderClient creates a docker api client
func NewDockerProviderClient() (ProviderClient, error) {
	return client.NewEnvClient()
}

// NewProvider creates a new provider
func NewProvider(cli ProviderClient) (*Provider, error) {
	return &Provider{
		client: cli,
	}, nil
}

// GetVolumesWithLabel returns a list of all docker volumes with the given label
func (p *Provider) GetVolumesWithLabel(label string) ([]Volume, error) {
	args := filters.NewArgs()
	args.Add("label", label)

	ctx := context.Background()
	response, err := p.client.VolumeList(ctx, args)
	if err != nil {
		return nil, err
	}

	dirs := make([]Volume, 0)
	for _, v := range response.Volumes {
		dirs = append(dirs, Volume{
			Name: v.Name,
			Path: v.Mountpoint,
		})
	}

	return dirs, nil
}
