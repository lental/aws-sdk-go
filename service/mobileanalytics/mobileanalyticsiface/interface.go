// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package mobileanalyticsiface provides an interface to enable mocking the Amazon Mobile Analytics service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mobileanalyticsiface

import (
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/service/mobileanalytics"
)

// MobileAnalyticsAPI provides an interface to enable mocking the
// mobileanalytics.MobileAnalytics service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Mobile Analytics.
//    func myFunc(svc mobileanalyticsiface.MobileAnalyticsAPI) bool {
//        // Make svc.PutEvents request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mobileanalytics.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMobileAnalyticsClient struct {
//        mobileanalyticsiface.MobileAnalyticsAPI
//    }
//    func (m *mockMobileAnalyticsClient) PutEvents(input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMobileAnalyticsClient{}
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
type MobileAnalyticsAPI interface {
	PutEventsRequest(*mobileanalytics.PutEventsInput) (*request.Request, *mobileanalytics.PutEventsOutput)

	PutEvents(*mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error)
}

var _ MobileAnalyticsAPI = (*mobileanalytics.MobileAnalytics)(nil)
