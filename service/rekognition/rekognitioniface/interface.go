// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package rekognitioniface provides an interface to enable mocking the Amazon Rekognition service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rekognitioniface

import (
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/service/rekognition"
)

// RekognitionAPI provides an interface to enable mocking the
// rekognition.Rekognition service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Rekognition.
//    func myFunc(svc rekognitioniface.RekognitionAPI) bool {
//        // Make svc.CompareFaces request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rekognition.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRekognitionClient struct {
//        rekognitioniface.RekognitionAPI
//    }
//    func (m *mockRekognitionClient) CompareFaces(input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRekognitionClient{}
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
type RekognitionAPI interface {
	CompareFacesRequest(*rekognition.CompareFacesInput) (*request.Request, *rekognition.CompareFacesOutput)

	CompareFaces(*rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error)

	CreateCollectionRequest(*rekognition.CreateCollectionInput) (*request.Request, *rekognition.CreateCollectionOutput)

	CreateCollection(*rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error)

	DeleteCollectionRequest(*rekognition.DeleteCollectionInput) (*request.Request, *rekognition.DeleteCollectionOutput)

	DeleteCollection(*rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error)

	DeleteFacesRequest(*rekognition.DeleteFacesInput) (*request.Request, *rekognition.DeleteFacesOutput)

	DeleteFaces(*rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error)

	DetectFacesRequest(*rekognition.DetectFacesInput) (*request.Request, *rekognition.DetectFacesOutput)

	DetectFaces(*rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error)

	DetectLabelsRequest(*rekognition.DetectLabelsInput) (*request.Request, *rekognition.DetectLabelsOutput)

	DetectLabels(*rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error)

	IndexFacesRequest(*rekognition.IndexFacesInput) (*request.Request, *rekognition.IndexFacesOutput)

	IndexFaces(*rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error)

	ListCollectionsRequest(*rekognition.ListCollectionsInput) (*request.Request, *rekognition.ListCollectionsOutput)

	ListCollections(*rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error)

	ListCollectionsPages(*rekognition.ListCollectionsInput, func(*rekognition.ListCollectionsOutput, bool) bool) error

	ListFacesRequest(*rekognition.ListFacesInput) (*request.Request, *rekognition.ListFacesOutput)

	ListFaces(*rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error)

	ListFacesPages(*rekognition.ListFacesInput, func(*rekognition.ListFacesOutput, bool) bool) error

	SearchFacesRequest(*rekognition.SearchFacesInput) (*request.Request, *rekognition.SearchFacesOutput)

	SearchFaces(*rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error)

	SearchFacesByImageRequest(*rekognition.SearchFacesByImageInput) (*request.Request, *rekognition.SearchFacesByImageOutput)

	SearchFacesByImage(*rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error)
}

var _ RekognitionAPI = (*rekognition.Rekognition)(nil)
