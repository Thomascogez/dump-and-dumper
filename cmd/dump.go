package cmd

import (
	"context"
	"fmt"

	"github.com/Thomascogez/dump-and-dumper/dumper"
	"github.com/Thomascogez/dump-and-dumper/helpers"
	upload "github.com/Thomascogez/dump-and-dumper/uploader"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Execute dump against all container using required label",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := client.NewClientWithOpts(client.FromEnv)
		helpers.CheckError(err)

		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		helpers.CheckError(err)

		notifiers := helpers.GetNotifiersFromFlags(cmd.Flags())
		s3Uploader := upload.NewS3UploaderFromFlags(cmd.Flags())
		dockerDumper := dumper.DockerDumper{Uploader: s3Uploader, Notifiers: notifiers}

		containersToDump := dumper.FindContainersToDump(containers)

		fmt.Printf("[dump & dumper] - Start dumping %d container(s) ...\n", len(containersToDump))

		dockerDumper.Dump(containersToDump)

		fmt.Printf("[dump & dumper] - Done dumping %d container(s)\n", len(containersToDump))
	},
}

func init() {
	dumpCmd.PersistentFlags().String("s3-endpoint", "", "Define endpoint of s3 bucket")
	dumpCmd.PersistentFlags().String("s3-region", "", "Define region of s3 bucket")
	dumpCmd.PersistentFlags().String("s3-bucket", "", "Define name of the destination bucket")
	dumpCmd.PersistentFlags().String("s3-secretKeyId", "", "Id of the secret in order to access s3 bucket")
	dumpCmd.PersistentFlags().String("s3-secretKey", "", "Secret key in order to access s3 bucket")
	dumpCmd.PersistentFlags().String("ntfy-endpoint", "", "Url of the ntfy server including the topic")

	rootCmd.AddCommand(dumpCmd)
}
