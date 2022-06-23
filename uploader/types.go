package upload

type Uploader interface {
	Upload(filePath string, fileName string)
}
