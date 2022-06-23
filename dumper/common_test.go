package dumper_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Thomascogez/dump-and-dumper/dumper"
)

func TestExtractDumpConfigFromLabels(t *testing.T) {

	labels := make(map[string]string)

	labels["go-dumper.enabled"] = "true"
	labels["go-dumper.user"] = "user"
	labels["go-dumper.type"] = "pg"

	dumpConfig := dumper.ExtractDumpOptionsFromLabels(labels)

	if !dumpConfig.Enabled || dumpConfig.Type != "pg" || dumpConfig.User != "user" {
		t.Fail()
	}
}

func TestBuildContainerDumpCommandArgs(t *testing.T) {
	containerDumpConfig := dumper.DumpOptions{
		Enabled: true,
		User:    "postgres",
		Type:    "pg",
	}

	testContainerId := "containerId"

	pgContainerDumpArgs := dumper.BuildContainerDumpCommandArgs(testContainerId, containerDumpConfig)

	argsString := strings.Join(pgContainerDumpArgs[:], " ")

	if argsString != fmt.Sprintf("exec %s pg_dumpall -U %s --if-exists -c", testContainerId, containerDumpConfig.User) {
		t.Fail()
	}

}
