package dumper

import (
	"os"
	"os/exec"
	"sync"

	"github.com/Thomascogez/dump-and-dumper/helpers"
	upload "github.com/Thomascogez/dump-and-dumper/uploader"
	"github.com/docker/docker/api/types"
)

type DockerDumper struct {
	Uploader upload.Uploader
}

func (dockerDumper DockerDumper) Dump(containers []types.Container) {

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