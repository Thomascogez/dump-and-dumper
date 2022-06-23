package dumper

import "github.com/docker/docker/api/types"

func ExtractDumpOptionsFromLabels(labels map[string]string) DumpOptions {
	return DumpOptions{
		Enabled: labels["go-dumper.enabled"] == "true",
		User:    labels["go-dumper.user"],
		Type:    labels["go-dumper.type"],
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

func FindContainersByTypes(containers []types.Container, containerType string) []types.Container {
	containerToDump := make([]types.Container, 0)

	for _, container := range containers {
		containerDumpConfig := ExtractDumpOptionsFromLabels(container.Labels)

		if containerDumpConfig.Enabled && containerDumpConfig.Type == containerType {
			containerToDump = append(containerToDump, container)
		}
	}

	return containerToDump
}
