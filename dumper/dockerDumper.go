package dumper

import (
	"os"
	"os/exec"
	"path"
	"sync"

	"github.com/Thomascogez/dump-and-dumper/helpers"
	"github.com/Thomascogez/dump-and-dumper/notifier"
	upload "github.com/Thomascogez/dump-and-dumper/uploader"
	"github.com/docker/docker/api/types"
)

type DockerDumper struct {
	Uploader  upload.Uploader
	Notifiers []notifier.Notifier
}

func (dockerDumper DockerDumper) Dump(containers []types.Container) {

	var wg sync.WaitGroup

	for _, container := range containers {
		wg.Add(1)

		go func(container types.Container) {
			defer wg.Done()
			dumpConfig := ExtractDumpOptionsFromLabels(container.Labels)
			dumpCommandArgs := BuildContainerDumpCommandArgs(container.ID, dumpConfig)

			tempDumpFile, tempDumpFolderPath, tempDumpFileName := helpers.CreateTempDumpFile()
			defer os.RemoveAll(tempDumpFolderPath)
			defer tempDumpFile.Close()

			dumpCommand := exec.Command("docker", dumpCommandArgs...)
			dumpCommand.Stdout = tempDumpFile
			dumpCommand.Run()

			dockerDumper.Uploader.Upload(
				path.Join(tempDumpFolderPath, tempDumpFileName),
				container.Names[0]+"-"+tempDumpFileName,
			)

			notifierMessage := "Successfully Dumped " + container.Names[0]

			for _, notifier := range dockerDumper.Notifiers {
				notifier.Notify(notifierMessage)
			}

		}(container)
	}
	wg.Wait()
}
