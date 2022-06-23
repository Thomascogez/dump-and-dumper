package dumper

import "github.com/docker/docker/api/types"

func ExtractDumpOptionsFromLabels(labels map[string]string) DumpOptions {
	return DumpOptions{
		Enabled: labels["dump-and-dumper.enabled"] == "true",
		User:    labels["dump-and-dumper.user"],
		Type:    labels["dump-and-dumper.type"],
	}
}

func BuildContainerDumpCommandArgs(containerId string, dumpOptions DumpOptions) []string {
	args := []string{"exec", containerId}

	if dumpOptions.Type == PG {
		args = append(args,
			"pg_dumpall",
			"-U", dumpOptions.User,
			"--if-exists",
			"-c",
		)
	}

	return args
}

func FindContainersToDump(containers []types.Container) []types.Container {
	containerToDump := make([]types.Container, 0)

	for _, container := range containers {
		containerDumpConfig := ExtractDumpOptionsFromLabels(container.Labels)

		if containerDumpConfig.Enabled {
			containerToDump = append(containerToDump, container)
		}
	}

	return containerToDump
}
