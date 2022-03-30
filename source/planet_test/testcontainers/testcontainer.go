package testcontainers

import (
	"context"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupMysqlContainer(t *testing.T, ctx context.Context, printContainerLogs bool) ContainerResult {
	req := testcontainers.ContainerRequest{
		Image: "mysql:latest",
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "mysql",
		},
		ExposedPorts: []string{"3306:3306"},
		WaitingFor:   wait.ForListeningPort("3306"),
	}

	port, _ := nat.NewPort("", "3306")

	return setupContainer(t, ctx, req, port, printContainerLogs)
}
