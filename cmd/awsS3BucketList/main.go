package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	result, err := listBuckets(ctx, s3Client)
	if err != nil {
		log.Fatalf("unable to list buckets: %v", err)
	}

	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Printf("* %s\n", aws.ToString(bucket.Name))
	}
}

func listBuckets(ctx context.Context, s3Client *s3.Client) (*s3.ListBucketsOutput, error) {
	input := &s3.ListBucketsInput{}

	result, err := s3Client.ListBuckets(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to list buckets: %w", err)
	}
	return result, nil
}
