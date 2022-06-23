package dumper_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Thomascogez/dump-and-dumper/dumper"
)

func TestExtractDumpConfigFromLabels(t *testing.T) {

	labels := make(map[string]string)

	labels["dump-and-dumper.enabled"] = "true"
	labels["dump-and-dumper.user"] = "user"
	labels["dump-and-dumper.type"] = "pg"

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
