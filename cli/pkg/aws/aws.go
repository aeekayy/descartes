package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func ListBuckets(ctx context.Context) ([]s3Types.Bucket, error) {
	var bkts []s3Types.Bucket

	ac, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		return bkts, err
	}

	// create s3 service
	s3c := s3.NewFromConfig(ac)

	o, err := s3c.ListBuckets(ctx, &s3.ListBucketsInput{})

	if err != nil {
		return bkts, err
	}

	bkts = o.Buckets
	return bkts, err
}
