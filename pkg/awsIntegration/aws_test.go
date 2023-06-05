package awsIntegration_test

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/ckone4you/golangtest/pkg/awsIntegration"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockS3API struct {
	mock.Mock
}

func (m *mockS3API) ListBuckets(input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*s3.ListBucketsOutput), args.Error(1)
}

func TestListS3Buckets(t *testing.T) {
	// Create a new mock S3API
	ctrl := NewController(t)
	defer ctrl.Finish()

	mockS3 := NewMockS3Interface(ctrl)

	mockS3 := new(mockS3API)

	// Set up the expected response
	mockOutput := &s3.ListBucketsOutput{
		Buckets: []*s3.Bucket{
			{Name: aws.String("bucket1"), CreationDate: aws.Time(time.Now())},
			{Name: aws.String("bucket2"), CreationDate: aws.Time(time.Now())},
		},
	}
	mockS3.On("ListBuckets", mock.AnythingOfType("*s3.ListBucketsInput")).Return(mockOutput, nil)

	// Create a new AWS session with the mock S3API
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("your-region"),
		Credentials: credentials.NewStaticCredentials("mock", "mock", "mock"),
		Endpoint:    aws.String("http://localhost:4566"), // for localstack
	}))

	// Create an S3 client with the mock S3API
	svc := s3.New(sess, &aws.Config{
		Endpoint: aws.String("http://localhost:4566"), // for localstack
	})
	client := awsIntegration.NewS3Client(svc)
	client.S3Client = mockS3

	// List S3 buckets
	buckets, err := awsIntegration.ListS3Buckets("your-region", client)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(buckets))
	assert.Equal(t, "bucket1", buckets[0].Name)
	assert.Equal(t, "bucket2", buckets[1].Name)

	// Verify that the mock S3API was called with the expected input
	mockS3.AssertCalled(t, "ListBuckets", mock.AnythingOfType("*s3.ListBucketsInput"))
}
