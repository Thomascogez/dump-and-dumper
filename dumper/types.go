package dumper

const (
	PG = "pg"
)

type DumpOptions struct {
	Enabled bool
	User    string
	Type    string
}
