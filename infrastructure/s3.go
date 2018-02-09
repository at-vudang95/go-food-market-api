package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	// StorageS3 is s3.
	StorageS3 = "s3"
	// StorageMinio is minio.
	StorageMinio = "minio"
)

// S3 struct.
type S3 struct {
	Client *s3.S3
}

// NewS3 returns new S3 Client instance.
func NewS3() *S3 {
	endpoint := GetConfigString("objectstorage.endpoint")
	region := GetConfigString("objectstorage.region")
	accessKeyID := GetConfigString("objectstorage.accesskey")
	secretAccessKey := GetConfigString("objectstorage.secretkey")
	secure := GetConfigBool("objectstorage.secure")
	storage := GetConfigString("objectstorage.storage")

	var s3Config *aws.Config
	if storage == StorageS3 {
		// Configure to use Server
		s3Config = &aws.Config{
			Region:           aws.String(region),
			DisableSSL:       aws.Bool(!secure),
			S3ForcePathStyle: aws.Bool(true),
		}
	} else if storage == StorageMinio {
		// Configure to use Server
		s3Config = &aws.Config{
			Endpoint:         aws.String(endpoint),
			Region:           aws.String(region),
			DisableSSL:       aws.Bool(!secure),
			S3ForcePathStyle: aws.Bool(true),
		}
	} else {
		panic("storage must select " + StorageS3 + " or " + StorageMinio)
	}

	s3Config.Credentials = credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	sess := session.Must(session.NewSession(s3Config))
	s3Client := s3.New(sess)

	return &S3{Client: s3Client}
}
