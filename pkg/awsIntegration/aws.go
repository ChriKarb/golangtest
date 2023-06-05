package awsIntegration

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type BucketInfo struct {
	Name         string
	CreationDate time.Time
}

type S3API interface {
	ListBuckets(input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error)
}

type AwsClient struct {
	S3Client *s3.S3
}

func NewS3Client(s3Client *s3.S3) *AwsClient {
	return &AwsClient{S3Client: s3Client}
}

func (c *AwsClient) ListBuckets(input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return c.S3Client.ListBuckets(input)
}

func main() {
	buckets, err := ListS3Buckets("your-region")
	if err != nil {
		fmt.Println("Error listing S3 buckets:", err)
		return
	}

	fmt.Println("S3 Buckets:")
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}
}

func ListS3Buckets(region string) ([]BucketInfo, error) {
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
	}

	// Create a new S3 client
	svc := s3.New(sess)

	// Create an AwsClient implementing the S3API
	client := NewS3Client(svc)

	// List S3 buckets
	result, err := client.ListBuckets(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list S3 buckets: %v", err)
	}

	bucketInfos := make([]BucketInfo, 0, len(result.Buckets))
	for _, bucket := range result.Buckets {
		info := BucketInfo{
			Name:         aws.StringValue(bucket.Name),
			CreationDate: aws.TimeValue(bucket.CreationDate),
		}
		bucketInfos = append(bucketInfos, info)
	}

	return bucketInfos, nil
}
