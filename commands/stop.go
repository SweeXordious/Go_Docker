package commands

import (
	"context"
	"github.com/docker/docker/client"
)

func StopContainer(containerID string) error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}

	cli.NegotiateAPIVersion(context.Background())

	err = cli.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		panic(err)
	}
	return err
}
