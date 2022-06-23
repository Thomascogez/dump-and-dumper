package dumper

import "github.com/docker/docker/api/types"

type Dumper interface {
	Dump(container *types.Container)
}

const (
	PG = "pg"
)

type DumpOptions struct {
	Enabled bool
	User    string
	Type    string
}
