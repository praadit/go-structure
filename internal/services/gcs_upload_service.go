package services

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"mime/multipart"
// 	"time"

// 	"cloud.google.com/go/storage"
// 	"gitlab.com/m8851/pmo-echo-api/config"
// )

// type GcsUploadServiceInterface interface {
// 	UploadFile(file multipart.File, object string) error
// 	UploadLocalFile(file io.Reader, object string) error
// 	DeleteFile(object string) error
// }

// type GcsUploadService struct {
// 	client     *storage.Client
// 	projectID  string
// 	bucketName string
// 	uploadPath string
// }

// func NewGcsUploadService(client *storage.Client) GcsUploadServiceInterface {
// 	return &GcsUploadService{
// 		client:     client,
// 		projectID:  config.AppConfig[config.GCSProjectId],
// 		bucketName: config.AppConfig[config.GCSBucketName],
// 		uploadPath: config.AppConfig[config.GCSUploadPath],
// 	}
// }

// func (uploader *GcsUploadService) UploadFile(file multipart.File, object string) error {
// 	ctx := context.Background()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
// 	defer cancel()

// 	// Upload an object with storage.Writer.
// 	wc := uploader.client.Bucket(uploader.bucketName).Object(uploader.uploadPath + object).NewWriter(ctx)
// 	wc.Metadata = map[string]string{
// 		"Cache-Control": "public, max-age=0, no-transform, no-cache",
// 	}
// 	if written, err := io.Copy(wc, file); err != nil {
// 		return errors.New("io.Copy" + err.Error())
// 	} else {
// 		fmt.Printf("%v bytes written to Cloud", written)
// 		fmt.Println()
// 	}

// 	if err := wc.Close(); err != nil {
// 		return errors.New("Writer.Close" + err.Error())
// 	}

// 	return nil
// }

// func (uploader *GcsUploadService) UploadLocalFile(file io.Reader, object string) error {
// 	ctx := context.Background()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
// 	defer cancel()

// 	// Upload an object with storage.Writer.
// 	wc := uploader.client.Bucket(uploader.bucketName).Object(uploader.uploadPath + object).NewWriter(ctx)
// 	wc.Metadata = map[string]string{
// 		"Cache-Control": "public, max-age=0, no-transform, no-cache",
// 	}
// 	if written, err := io.Copy(wc, file); err != nil {
// 		return errors.New("io.Copy" + err.Error())
// 	} else {
// 		fmt.Printf("%v bytes written to Cloud", written)
// 		fmt.Println()
// 	}

// 	if err := wc.Close(); err != nil {
// 		return errors.New("Writer.Close" + err.Error())
// 	}

// 	return nil
// }

// func (uploader *GcsUploadService) DeleteFile(object string) error {
// 	ctx := context.Background()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
// 	defer cancel()

// 	o := uploader.client.Bucket(uploader.bucketName).Object(uploader.uploadPath + object)
// 	_, err := o.Attrs(ctx)
// 	if err == storage.ErrObjectNotExist {
// 		return nil
// 	}
// 	if err != nil {
// 		return err
// 	}

// 	if err := o.Delete(ctx); err != nil {
// 		return fmt.Errorf("Object(%q).Delete: %v", object, err)
// 	}
// 	return nil
// }
