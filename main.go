package main

import (
	"context"

	"github.com/Thomascogez/dump-and-dumper/dumper"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	pgContainersToDump := dumper.FindContainersByTypes(containers, dumper.PG)

	dumper.DockerDumper{}.Dump(pgContainersToDump)
}
