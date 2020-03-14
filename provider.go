package dockervolumebackup

import (
	"context"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// DockerProvider contains methods to find volumes to backup
type DockerProvider struct {
	client *client.Client
}

// Volume represents a docker volume to backup
type Volume struct {
	Name string
	Path string
}

// NewProvider creates a new provider
func NewProvider() (*DockerProvider, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &DockerProvider{
		client: cli,
	}, nil
}

// GetVolumesWithLabel returns a list of all docker volumes with the given label
func (p *DockerProvider) GetVolumesWithLabel(label string) ([]Volume, error) {
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
