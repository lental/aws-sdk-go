// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package snowball_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/session"
	"github.com/lental/aws-sdk-go/service/snowball"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSnowball_CancelCluster() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.CancelClusterInput{
		ClusterId: aws.String("ClusterId"), // Required
	}
	resp, err := svc.CancelCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_CancelJob() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.CancelJobInput{
		JobId: aws.String("JobId"), // Required
	}
	resp, err := svc.CancelJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_CreateAddress() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.CreateAddressInput{
		Address: &snowball.Address{ // Required
			AddressId:            aws.String("AddressId"),
			City:                 aws.String("String"),
			Company:              aws.String("String"),
			Country:              aws.String("String"),
			Landmark:             aws.String("String"),
			Name:                 aws.String("String"),
			PhoneNumber:          aws.String("String"),
			PostalCode:           aws.String("String"),
			PrefectureOrDistrict: aws.String("String"),
			StateOrProvince:      aws.String("String"),
			Street1:              aws.String("String"),
			Street2:              aws.String("String"),
			Street3:              aws.String("String"),
		},
	}
	resp, err := svc.CreateAddress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_CreateCluster() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.CreateClusterInput{
		AddressId: aws.String("AddressId"), // Required
		JobType:   aws.String("JobType"),   // Required
		Resources: &snowball.JobResource{ // Required
			LambdaResources: []*snowball.LambdaResource{
				{ // Required
					EventTriggers: []*snowball.EventTriggerDefinition{
						{ // Required
							EventResourceARN: aws.String("ResourceARN"),
						},
						// More values...
					},
					LambdaArn: aws.String("ResourceARN"),
				},
				// More values...
			},
			S3Resources: []*snowball.S3Resource{
				{ // Required
					BucketArn: aws.String("ResourceARN"),
					KeyRange: &snowball.KeyRange{
						BeginMarker: aws.String("String"),
						EndMarker:   aws.String("String"),
					},
				},
				// More values...
			},
		},
		RoleARN:        aws.String("RoleARN"),        // Required
		ShippingOption: aws.String("ShippingOption"), // Required
		Description:    aws.String("String"),
		KmsKeyARN:      aws.String("KmsKeyARN"),
		Notification: &snowball.Notification{
			JobStatesToNotify: []*string{
				aws.String("JobState"), // Required
				// More values...
			},
			NotifyAll:   aws.Bool(true),
			SnsTopicARN: aws.String("SnsTopicARN"),
		},
		SnowballType: aws.String("Type"),
	}
	resp, err := svc.CreateCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_CreateJob() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.CreateJobInput{
		AddressId:   aws.String("AddressId"),
		ClusterId:   aws.String("ClusterId"),
		Description: aws.String("String"),
		JobType:     aws.String("JobType"),
		KmsKeyARN:   aws.String("KmsKeyARN"),
		Notification: &snowball.Notification{
			JobStatesToNotify: []*string{
				aws.String("JobState"), // Required
				// More values...
			},
			NotifyAll:   aws.Bool(true),
			SnsTopicARN: aws.String("SnsTopicARN"),
		},
		Resources: &snowball.JobResource{
			LambdaResources: []*snowball.LambdaResource{
				{ // Required
					EventTriggers: []*snowball.EventTriggerDefinition{
						{ // Required
							EventResourceARN: aws.String("ResourceARN"),
						},
						// More values...
					},
					LambdaArn: aws.String("ResourceARN"),
				},
				// More values...
			},
			S3Resources: []*snowball.S3Resource{
				{ // Required
					BucketArn: aws.String("ResourceARN"),
					KeyRange: &snowball.KeyRange{
						BeginMarker: aws.String("String"),
						EndMarker:   aws.String("String"),
					},
				},
				// More values...
			},
		},
		RoleARN:                    aws.String("RoleARN"),
		ShippingOption:             aws.String("ShippingOption"),
		SnowballCapacityPreference: aws.String("Capacity"),
		SnowballType:               aws.String("Type"),
	}
	resp, err := svc.CreateJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_DescribeAddress() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.DescribeAddressInput{
		AddressId: aws.String("AddressId"), // Required
	}
	resp, err := svc.DescribeAddress(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_DescribeAddresses() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.DescribeAddressesInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.DescribeAddresses(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_DescribeCluster() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.DescribeClusterInput{
		ClusterId: aws.String("ClusterId"), // Required
	}
	resp, err := svc.DescribeCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_DescribeJob() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.DescribeJobInput{
		JobId: aws.String("JobId"), // Required
	}
	resp, err := svc.DescribeJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_GetJobManifest() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.GetJobManifestInput{
		JobId: aws.String("JobId"), // Required
	}
	resp, err := svc.GetJobManifest(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_GetJobUnlockCode() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.GetJobUnlockCodeInput{
		JobId: aws.String("JobId"), // Required
	}
	resp, err := svc.GetJobUnlockCode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_GetSnowballUsage() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	var params *snowball.GetSnowballUsageInput
	resp, err := svc.GetSnowballUsage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_ListClusterJobs() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.ListClusterJobsInput{
		ClusterId:  aws.String("ClusterId"), // Required
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListClusterJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_ListClusters() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.ListClustersInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListClusters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_ListJobs() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.ListJobsInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_UpdateCluster() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.UpdateClusterInput{
		ClusterId:   aws.String("ClusterId"), // Required
		AddressId:   aws.String("AddressId"),
		Description: aws.String("String"),
		Notification: &snowball.Notification{
			JobStatesToNotify: []*string{
				aws.String("JobState"), // Required
				// More values...
			},
			NotifyAll:   aws.Bool(true),
			SnsTopicARN: aws.String("SnsTopicARN"),
		},
		Resources: &snowball.JobResource{
			LambdaResources: []*snowball.LambdaResource{
				{ // Required
					EventTriggers: []*snowball.EventTriggerDefinition{
						{ // Required
							EventResourceARN: aws.String("ResourceARN"),
						},
						// More values...
					},
					LambdaArn: aws.String("ResourceARN"),
				},
				// More values...
			},
			S3Resources: []*snowball.S3Resource{
				{ // Required
					BucketArn: aws.String("ResourceARN"),
					KeyRange: &snowball.KeyRange{
						BeginMarker: aws.String("String"),
						EndMarker:   aws.String("String"),
					},
				},
				// More values...
			},
		},
		RoleARN:        aws.String("RoleARN"),
		ShippingOption: aws.String("ShippingOption"),
	}
	resp, err := svc.UpdateCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSnowball_UpdateJob() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := snowball.New(sess)

	params := &snowball.UpdateJobInput{
		JobId:       aws.String("JobId"), // Required
		AddressId:   aws.String("AddressId"),
		Description: aws.String("String"),
		Notification: &snowball.Notification{
			JobStatesToNotify: []*string{
				aws.String("JobState"), // Required
				// More values...
			},
			NotifyAll:   aws.Bool(true),
			SnsTopicARN: aws.String("SnsTopicARN"),
		},
		Resources: &snowball.JobResource{
			LambdaResources: []*snowball.LambdaResource{
				{ // Required
					EventTriggers: []*snowball.EventTriggerDefinition{
						{ // Required
							EventResourceARN: aws.String("ResourceARN"),
						},
						// More values...
					},
					LambdaArn: aws.String("ResourceARN"),
				},
				// More values...
			},
			S3Resources: []*snowball.S3Resource{
				{ // Required
					BucketArn: aws.String("ResourceARN"),
					KeyRange: &snowball.KeyRange{
						BeginMarker: aws.String("String"),
						EndMarker:   aws.String("String"),
					},
				},
				// More values...
			},
		},
		RoleARN:                    aws.String("RoleARN"),
		ShippingOption:             aws.String("ShippingOption"),
		SnowballCapacityPreference: aws.String("Capacity"),
	}
	resp, err := svc.UpdateJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
