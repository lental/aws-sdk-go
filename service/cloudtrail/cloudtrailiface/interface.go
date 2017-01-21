// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudtrailiface provides an interface to enable mocking the AWS CloudTrail service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudtrailiface

import (
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/service/cloudtrail"
)

// CloudTrailAPI provides an interface to enable mocking the
// cloudtrail.CloudTrail service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CloudTrail.
//    func myFunc(svc cloudtrailiface.CloudTrailAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudtrail.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudTrailClient struct {
//        cloudtrailiface.CloudTrailAPI
//    }
//    func (m *mockCloudTrailClient) AddTags(input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudTrailClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudTrailAPI interface {
	AddTagsRequest(*cloudtrail.AddTagsInput) (*request.Request, *cloudtrail.AddTagsOutput)

	AddTags(*cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error)

	CreateTrailRequest(*cloudtrail.CreateTrailInput) (*request.Request, *cloudtrail.CreateTrailOutput)

	CreateTrail(*cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error)

	DeleteTrailRequest(*cloudtrail.DeleteTrailInput) (*request.Request, *cloudtrail.DeleteTrailOutput)

	DeleteTrail(*cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error)

	DescribeTrailsRequest(*cloudtrail.DescribeTrailsInput) (*request.Request, *cloudtrail.DescribeTrailsOutput)

	DescribeTrails(*cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error)

	GetEventSelectorsRequest(*cloudtrail.GetEventSelectorsInput) (*request.Request, *cloudtrail.GetEventSelectorsOutput)

	GetEventSelectors(*cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error)

	GetTrailStatusRequest(*cloudtrail.GetTrailStatusInput) (*request.Request, *cloudtrail.GetTrailStatusOutput)

	GetTrailStatus(*cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error)

	ListPublicKeysRequest(*cloudtrail.ListPublicKeysInput) (*request.Request, *cloudtrail.ListPublicKeysOutput)

	ListPublicKeys(*cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error)

	ListTagsRequest(*cloudtrail.ListTagsInput) (*request.Request, *cloudtrail.ListTagsOutput)

	ListTags(*cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error)

	LookupEventsRequest(*cloudtrail.LookupEventsInput) (*request.Request, *cloudtrail.LookupEventsOutput)

	LookupEvents(*cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error)

	PutEventSelectorsRequest(*cloudtrail.PutEventSelectorsInput) (*request.Request, *cloudtrail.PutEventSelectorsOutput)

	PutEventSelectors(*cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error)

	RemoveTagsRequest(*cloudtrail.RemoveTagsInput) (*request.Request, *cloudtrail.RemoveTagsOutput)

	RemoveTags(*cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error)

	StartLoggingRequest(*cloudtrail.StartLoggingInput) (*request.Request, *cloudtrail.StartLoggingOutput)

	StartLogging(*cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error)

	StopLoggingRequest(*cloudtrail.StopLoggingInput) (*request.Request, *cloudtrail.StopLoggingOutput)

	StopLogging(*cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error)

	UpdateTrailRequest(*cloudtrail.UpdateTrailInput) (*request.Request, *cloudtrail.UpdateTrailOutput)

	UpdateTrail(*cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error)
}

var _ CloudTrailAPI = (*cloudtrail.CloudTrail)(nil)
