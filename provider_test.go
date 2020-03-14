package dockervolumebackup

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/stretchr/testify/require"
)

type testClient struct {
	body volume.VolumesListOKBody
}

func (c *testClient) VolumeList(ctx context.Context, filter filters.Args) (volume.VolumesListOKBody, error) {
	return c.body, nil
}

func TestGetVolumesWithLabel(t *testing.T) {
	mock := &testClient{
		body: volume.VolumesListOKBody{
			Volumes: []*types.Volume{
				&types.Volume{Name: "the-first-volume", Mountpoint: "/path/to/first/mountpoint"},
				&types.Volume{Name: "the-second-volume", Mountpoint: "/path/to/second/mountpoint"},
				&types.Volume{Name: "the-third-volume", Mountpoint: "/path/to/third/mountpoint"},
			},
		},
	}

	provider, err := NewProvider(mock)
	require.Nil(t, err)

	volumes, err := provider.GetVolumesWithLabel("com.example.label")
	require.Nil(t, err)
	require.Len(t, volumes, 3)
	require.Equal(t, Volume{Name: "the-first-volume", Path: "/path/to/first/mountpoint"}, volumes[0])
	require.Equal(t, Volume{Name: "the-second-volume", Path: "/path/to/second/mountpoint"}, volumes[1])
	require.Equal(t, Volume{Name: "the-third-volume", Path: "/path/to/third/mountpoint"}, volumes[2])
}
