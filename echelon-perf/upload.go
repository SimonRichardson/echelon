package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/SimonRichardson/echelon/internal/typex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	fileType = "text/plain; charset=utf-8"
	filePath = "/perf/echelon/perf-%s.csv"

	awsId     = ""
	awsSecret = ""
	awsToken  = ""
	awsRegion = ""
	awsBucket = "ci-metrics"
)

func upload(content []byte) {
	creds := credentials.NewStaticCredentials(awsId, awsSecret, awsToken)
	if _, err := creds.Get(); err != nil {
		typex.Fatal("Invalid credentials", err)
	}

	var (
		cfg = aws.NewConfig().WithRegion(awsRegion).WithCredentials(creds)
		svc = s3.New(session.New(), cfg)

		now  = time.Now()
		path = fmt.Sprintf(filePath, now.Format(time.RFC3339))
	)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(awsBucket),
		Key:           aws.String(path),
		Body:          bytes.NewReader(content),
		ContentLength: aws.Int64(int64(len(content))),
		ContentType:   aws.String(fileType),
	}

	_, err := svc.PutObject(params)
	if err != nil {
		typex.Fatal(err)
	}
}
