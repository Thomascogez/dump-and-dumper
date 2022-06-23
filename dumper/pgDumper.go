package dumper

import (
	"os"
	"os/exec"
	"sync"

	"github.com/Thomascogez/dump-and-dumper/helpers"
	"github.com/docker/docker/api/types"
)

type PgDumper struct{}

func (pgDumper PgDumper) Dump(containers []types.Container) {

	var wg sync.WaitGroup

	for _, container := range containers {
		wg.Add(1)

		go func(container types.Container) {
			dumpConfig := ExtractDumpOptionsFromLabels(container.Labels)
			dumpCommandArgs := BuildContainerDumpCommandArgs(container.ID, dumpConfig)

			tempDumpFile, tempDumpFolderPath, tempDumpFileName := helpers.CreateTempDumpFile()

			dumpCommand := exec.Command("docker", dumpCommandArgs...)
			dumpCommand.Stdout = tempDumpFile
			dumpCommand.Run()

			println(tempDumpFileName)
			tempDumpFile.Close()
			os.RemoveAll(tempDumpFolderPath)
			wg.Done()
		}(container)
	}

	wg.Wait()
}
