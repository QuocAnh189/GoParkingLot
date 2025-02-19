package minio

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client  *minio.Client
	Bucket  string
	BaseURL string
}

// NewMinioClient khởi tạo Minio client
func NewMinioClient(endpoint, accessKey, secretKey, bucket, baseURL string, useSSL bool) (*MinioClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	// Kiểm tra bucket, nếu chưa có thì tạo
	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return &MinioClient{
		Client:  client,
		Bucket:  bucket,
		BaseURL: baseURL,
	}, nil
}

func (m *MinioClient) UploadFile(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if folder == "" {
		folder = "avatars"
	}

	objectName := fmt.Sprintf("%s/%d-%s", folder, time.Now().UnixNano(), file.Filename)

	_, err = m.Client.PutObject(ctx, m.Bucket, objectName, src, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	// Trả về URL ảnh đã upload
	return fmt.Sprintf("%s/%s/%s", m.BaseURL, m.Bucket, objectName), nil
}
