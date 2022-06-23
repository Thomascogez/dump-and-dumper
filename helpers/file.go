package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func FileNameFromCurrentTimestamp(fileExtension string) string {
	return fmt.Sprintf("%d.%s", time.Now().UnixMilli(), fileExtension)
}

func CreateTempDumpFile() (*os.File, string, string) {
	cwd, err := os.Getwd()
	CheckError(err)

	tempDumpFolderPath, err := ioutil.TempDir(cwd, "dump-")

	tempDumpFileName := FileNameFromCurrentTimestamp("sql")
	tempDumpFilePath := path.Join(tempDumpFolderPath, tempDumpFileName)
	tempDumpFile, err := os.Create(tempDumpFilePath)

	if err != nil {
		os.RemoveAll(tempDumpFolderPath)
		CheckError(err)
	}

	return tempDumpFile, tempDumpFolderPath, tempDumpFileName
}
