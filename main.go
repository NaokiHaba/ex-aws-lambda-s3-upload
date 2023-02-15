package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"path/filepath"
)

func main() {
	// Initialize the session that the SDK
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	if err != nil {
		fmt.Errorf("error creating session: %v", err)
		return
	}

	// Create a new service client
	svc := s3.New(sess)

	// Create a new request for the ListBuckets operation
	dumpFilePath := "./hoge.sql"

	// パスの最後の要素（ファイル名を取得）
	dumpFileName := filepath.Base(dumpFilePath)

	// S3バケット
	bucket := "ex-bucket-lambda"

	// S3オブジェクトキー
	key := "dump/" + dumpFileName

	// ファイルを開く
	file, err := os.Open(dumpFilePath)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
		return
	}

	defer file.Close()

	// S3にアップロード
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		fmt.Errorf("error uploading file: %v", err)
		return
	}

	fmt.Println("success")
}
