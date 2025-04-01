package minioClient

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
	"os"
)

var minioClient *minio.Client

func NewMinioClient() {
	endpoint := os.Getenv("ENDPOINT_MINIO")
	accessKeyID := os.Getenv("ACCESS_KEY_MINIO")
	secretAccessKey := os.Getenv("SECRET_ACCESS_MINIO")

	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Error().Err(err).Msg("Erro ao inicializar o cliente do MinIO")
		return
	}
}

func GetMinioClient() *minio.Client {
	return minioClient
}
