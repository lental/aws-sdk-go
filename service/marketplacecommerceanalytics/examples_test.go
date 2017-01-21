// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package marketplacecommerceanalytics_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/session"
	"github.com/lental/aws-sdk-go/service/marketplacecommerceanalytics"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleMarketplaceCommerceAnalytics_GenerateDataSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := marketplacecommerceanalytics.New(sess)

	params := &marketplacecommerceanalytics.GenerateDataSetInput{
		DataSetPublicationDate:  aws.Time(time.Now()),                  // Required
		DataSetType:             aws.String("DataSetType"),             // Required
		DestinationS3BucketName: aws.String("DestinationS3BucketName"), // Required
		RoleNameArn:             aws.String("RoleNameArn"),             // Required
		SnsTopicArn:             aws.String("SnsTopicArn"),             // Required
		CustomerDefinedValues: map[string]*string{
			"Key": aws.String("OptionalValue"), // Required
			// More values...
		},
		DestinationS3Prefix: aws.String("DestinationS3Prefix"),
	}
	resp, err := svc.GenerateDataSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleMarketplaceCommerceAnalytics_StartSupportDataExport() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := marketplacecommerceanalytics.New(sess)

	params := &marketplacecommerceanalytics.StartSupportDataExportInput{
		DataSetType:             aws.String("SupportDataSetType"),      // Required
		DestinationS3BucketName: aws.String("DestinationS3BucketName"), // Required
		FromDate:                aws.Time(time.Now()),                  // Required
		RoleNameArn:             aws.String("RoleNameArn"),             // Required
		SnsTopicArn:             aws.String("SnsTopicArn"),             // Required
		CustomerDefinedValues: map[string]*string{
			"Key": aws.String("OptionalValue"), // Required
			// More values...
		},
		DestinationS3Prefix: aws.String("DestinationS3Prefix"),
	}
	resp, err := svc.StartSupportDataExport(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
