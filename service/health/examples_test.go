// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package health_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/session"
	"github.com/lental/aws-sdk-go/service/health"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleHealth_DescribeAffectedEntities() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := health.New(sess)

	params := &health.DescribeAffectedEntitiesInput{
		Filter: &health.EntityFilter{ // Required
			EventArns: []*string{ // Required
				aws.String("eventArn"), // Required
				// More values...
			},
			EntityArns: []*string{
				aws.String("entityArn"), // Required
				// More values...
			},
			EntityValues: []*string{
				aws.String("entityValue"), // Required
				// More values...
			},
			LastUpdatedTimes: []*health.DateTimeRange{
				{ // Required
					From: aws.Time(time.Now()),
					To:   aws.Time(time.Now()),
				},
				// More values...
			},
			StatusCodes: []*string{
				aws.String("entityStatusCode"), // Required
				// More values...
			},
			Tags: []map[string]*string{
				{ // Required
					"Key": aws.String("tagValue"), // Required
					// More values...
				},
				// More values...
			},
		},
		Locale:     aws.String("locale"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("nextToken"),
	}
	resp, err := svc.DescribeAffectedEntities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleHealth_DescribeEntityAggregates() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := health.New(sess)

	params := &health.DescribeEntityAggregatesInput{
		EventArns: []*string{
			aws.String("eventArn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeEntityAggregates(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleHealth_DescribeEventAggregates() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := health.New(sess)

	params := &health.DescribeEventAggregatesInput{
		AggregateField: aws.String("eventAggregateField"), // Required
		Filter: &health.EventFilter{
			AvailabilityZones: []*string{
				aws.String("availabilityZone"), // Required
				// More values...
			},
			EndTimes: []*health.DateTimeRange{
				{ // Required
					From: aws.Time(time.Now()),
					To:   aws.Time(time.Now()),
				},
				// More values...
			},
			EntityArns: []*string{
				aws.String("entityArn"), // Required
				// More values...
			},
			EntityValues: []*string{
				aws.String("entityValue"), // Required
				// More values...
			},
			EventArns: []*string{
				aws.String("eventArn"), // Required
				// More values...
			},
			EventStatusCodes: []*string{
				aws.String("eventStatusCode"), // Required
				// More values...
			},
			EventTypeCategories: []*string{
				aws.String("eventTypeCategory"), // Required
				// More values...
			},
			EventTypeCodes: []*string{
				aws.String("eventType"), // Required
				// More values...
			},
			LastUpdatedTimes: []*health.DateTimeRange{
				{ // Required
					From: aws.Time(time.Now()),
					To:   aws.Time(time.Now()),
				},
				// More values...
			},
			Regions: []*string{
				aws.String("region"), // Required
				// More values...
			},
			Services: []*string{
				aws.String("service"), // Required
				// More values...
			},
			StartTimes: []*health.DateTimeRange{
				{ // Required
					From: aws.Time(time.Now()),
					To:   aws.Time(time.Now()),
				},
				// More values...
			},
			Tags: []map[string]*string{
				{ // Required
					"Key": aws.String("tagValue"), // Required
					// More values...
				},
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("nextToken"),
	}
	resp, err := svc.DescribeEventAggregates(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleHealth_DescribeEventDetails() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := health.New(sess)

	params := &health.DescribeEventDetailsInput{
		EventArns: []*string{ // Required
			aws.String("eventArn"), // Required
			// More values...
		},
		Locale: aws.String("locale"),
	}
	resp, err := svc.DescribeEventDetails(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleHealth_DescribeEventTypes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := health.New(sess)

	params := &health.DescribeEventTypesInput{
		Filter: &health.EventTypeFilter{
			EventTypeCategories: []*string{
				aws.String("eventTypeCategory"), // Required
				// More values...
			},
			EventTypeCodes: []*string{
				aws.String("eventTypeCode"), // Required
				// More values...
			},
			Services: []*string{
				aws.String("service"), // Required
				// More values...
			},
		},
		Locale:     aws.String("locale"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("nextToken"),
	}
	resp, err := svc.DescribeEventTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleHealth_DescribeEvents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := health.New(sess)

	params := &health.DescribeEventsInput{
		Filter: &health.EventFilter{
			AvailabilityZones: []*string{
				aws.String("availabilityZone"), // Required
				// More values...
			},
			EndTimes: []*health.DateTimeRange{
				{ // Required
					From: aws.Time(time.Now()),
					To:   aws.Time(time.Now()),
				},
				// More values...
			},
			EntityArns: []*string{
				aws.String("entityArn"), // Required
				// More values...
			},
			EntityValues: []*string{
				aws.String("entityValue"), // Required
				// More values...
			},
			EventArns: []*string{
				aws.String("eventArn"), // Required
				// More values...
			},
			EventStatusCodes: []*string{
				aws.String("eventStatusCode"), // Required
				// More values...
			},
			EventTypeCategories: []*string{
				aws.String("eventTypeCategory"), // Required
				// More values...
			},
			EventTypeCodes: []*string{
				aws.String("eventType"), // Required
				// More values...
			},
			LastUpdatedTimes: []*health.DateTimeRange{
				{ // Required
					From: aws.Time(time.Now()),
					To:   aws.Time(time.Now()),
				},
				// More values...
			},
			Regions: []*string{
				aws.String("region"), // Required
				// More values...
			},
			Services: []*string{
				aws.String("service"), // Required
				// More values...
			},
			StartTimes: []*health.DateTimeRange{
				{ // Required
					From: aws.Time(time.Now()),
					To:   aws.Time(time.Now()),
				},
				// More values...
			},
			Tags: []map[string]*string{
				{ // Required
					"Key": aws.String("tagValue"), // Required
					// More values...
				},
				// More values...
			},
		},
		Locale:     aws.String("locale"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("nextToken"),
	}
	resp, err := svc.DescribeEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
