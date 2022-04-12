package testcontainers

import (
	"context"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type Testcontainer struct {
	Database     string
	RootPassword string
}

func SetupMysqlContainer(config *Testcontainer) ContainerResult {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "mysql:latest",
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": config.RootPassword,
			"MYSQL_DATABASE":      config.Database,
		},
		ExposedPorts: []string{"3306:3306"},
		WaitingFor:   wait.ForListeningPort("3306"),
	}

	port, _ := nat.NewPort("", "3306")

	return setupContainer(ctx, req, port, true)
}
