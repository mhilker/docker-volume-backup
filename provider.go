package dockervolumebackup

import (
	"context"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type DockerProvider struct {
	client *client.Client
}

type BackupDirectory struct {
	Name string
	Path string
}

func NewProvider() (*DockerProvider, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &DockerProvider{
		client: cli,
	}, nil
}

func (p *DockerProvider) GetDirectories(label string) ([]BackupDirectory, error) {
	args := filters.NewArgs()
	args.Add("label", label)

	ctx := context.Background()
	response, err := p.client.VolumeList(ctx, args)
	if err != nil {
		return nil, err
	}

	dirs := make([]BackupDirectory, 0)
	for _, v := range response.Volumes {
		dirs = append(dirs, BackupDirectory{
			Name: v.Name,
			Path: v.Mountpoint,
		})
	}

	return dirs, nil
}
