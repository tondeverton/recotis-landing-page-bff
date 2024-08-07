package internal

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

const errCreateClient = "failed to create client"

func generateSignedURL(bucketName, objectName string, expiresIn time.Time) (string, error) {
	context := context.Background()

	client, err := storage.NewClient(context)
	if err != nil {
		return "", fmt.Errorf("%v: %v", errCreateClient, err)
	}
	defer client.Close()

	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: expiresIn,
	}

	url, err := client.Bucket(bucketName).SignedURL(objectName, opts)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %v", err)
	}

	return url, nil
}

func GetUrlsByAssets(assets []string) (map[string]string, time.Time, error) {
	bucketName := GetBucketName()
	bucketBasePath := GetBucketBasePath()
	signedUrlExpiration := GetSignedUrlExpiration()

	expiresIn := time.Now().Add(time.Duration(signedUrlExpiration) * time.Second)
	urls := make(map[string]string)

	for i := 0; i < len(assets); i++ {
		asset := assets[i]
		signedUrl, err := generateSignedURL(bucketName, bucketBasePath+"/"+asset, expiresIn)
		if err != nil {
			return nil, expiresIn, err
		}
		urls[asset] = signedUrl
	}

	return urls, expiresIn, nil
}
