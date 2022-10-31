package tests

import (
	"context"
	"fmt"

	containers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	DockerContext = "../"
	Dockerfile    = "Dockerfile"
	ContainerName = "testcontainers_app"
)

type applicationContainer struct {
	containers.Container

	URI   string
	close func(ctx context.Context)
}

func setupTestEnvironment(ctx context.Context) (*applicationContainer, error) {
	container, err := containers.GenericContainer(
		ctx,
		containers.GenericContainerRequest{
			ContainerRequest: containers.ContainerRequest{
				FromDockerfile: containers.FromDockerfile{
					Context:       DockerContext,
					Dockerfile:    Dockerfile,
					PrintBuildLog: true,
				},
				ExposedPorts: []string{"8080"},
				WaitingFor:   wait.ForLog("starting endpoints on :8080"),
				Name:         ContainerName,
			},
			Started: true,
			Reuse:   true,
		},
	)
	if err != nil {
		return nil, err
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := container.MappedPort(ctx, "8080")
	if err != nil {
		return nil, err
	}

	return &applicationContainer{
		Container: container,
		URI:       fmt.Sprintf("http://%s:%s", host, port.Port()),
		close: func(ctx context.Context) {
			container.Terminate(ctx) //nolint:errcheck
		},
	}, nil
}
