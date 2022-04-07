package testcontainers

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
)

type ContainerResult struct {
	Container testcontainers.Container
	Host      string
	Port      uint
}

func (c ContainerResult) ConnectionURI() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type TestLogConsumer struct {
}

func (g *TestLogConsumer) Accept(l testcontainers.Log) {
	log.Print(string(l.Content))
}

func setupContainer(
	ctx context.Context,
	containerRequest testcontainers.ContainerRequest,
	nPort nat.Port,
	printContainerLogs bool) ContainerResult {

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerRequest,
		Started:          true,
	})

	if err != nil {
		log.Panicf("Failed to start container %+v, with error: %v", containerRequest, err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		log.Panicf("Failed to retrive host %+v, with error: %v", containerRequest, err)
	}

	port, err := container.MappedPort(ctx, nPort)
	if err != nil {
		log.Panicf("Failed to retrive port %+v, with error: %v", containerRequest, err)
	}

	if printContainerLogs {
		logConsumer := TestLogConsumer{}

		err = container.StartLogProducer(ctx)
		if err != nil {
			log.Panicf("%s", err)
		}
		container.FollowOutput(&logConsumer)
	}

	return ContainerResult{
		Container: container,
		Host:      host,
		Port:      uint(port.Int()),
	}
}
