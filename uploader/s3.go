package upload

import (
	"fmt"
	"os"

	"github.com/Thomascogez/dump-and-dumper/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/pflag"
)

type S3Uploader struct {
	SecretKeyID string
	SecretKey   string
	Endpoint    string
	Region      string
	Bucket      string
}

func (s3Uploader S3Uploader) Upload(filePath string, fileName string) {
	s3ClientConfig := &aws.Config{
		Endpoint: aws.String(s3Uploader.Endpoint),
		Region:   aws.String(s3Uploader.Region),
		Credentials: credentials.NewStaticCredentials(
			s3Uploader.SecretKeyID,
			s3Uploader.SecretKey,
			"",
		),
	}

	sess, err := session.NewSession(s3ClientConfig)
	helpers.CheckError(err)

	uploader := s3manager.NewUploader(sess)
	dumpFile, _ := os.Open(filePath)
	defer dumpFile.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3Uploader.Bucket),
		Key:    aws.String(fileName),
		Body:   dumpFile,
	})

	helpers.CheckError(err)

	fmt.Printf("[dump & dumper] - file uploaded to, %s\n", result.Location)
}

func NewS3UploaderFromFlags(flags *pflag.FlagSet) S3Uploader {
	endPoint, _ := flags.GetString("s3-endpoint")
	region, _ := flags.GetString("s3-region")
	bucket, _ := flags.GetString("s3-bucket")
	secretKeyId, _ := flags.GetString("s3-secretKeyId")
	secretKey, _ := flags.GetString("s3-secretKey")

	return S3Uploader{
		SecretKeyID: secretKeyId,
		SecretKey:   secretKey,
		Bucket:      bucket,
		Endpoint:    endPoint,
		Region:      region,
	}
}
