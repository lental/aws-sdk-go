// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package datapipelineiface provides an interface to enable mocking the AWS Data Pipeline service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package datapipelineiface

import (
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/service/datapipeline"
)

// DataPipelineAPI provides an interface to enable mocking the
// datapipeline.DataPipeline service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Data Pipeline.
//    func myFunc(svc datapipelineiface.DataPipelineAPI) bool {
//        // Make svc.ActivatePipeline request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := datapipeline.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDataPipelineClient struct {
//        datapipelineiface.DataPipelineAPI
//    }
//    func (m *mockDataPipelineClient) ActivatePipeline(input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDataPipelineClient{}
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
type DataPipelineAPI interface {
	ActivatePipelineRequest(*datapipeline.ActivatePipelineInput) (*request.Request, *datapipeline.ActivatePipelineOutput)

	ActivatePipeline(*datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error)

	AddTagsRequest(*datapipeline.AddTagsInput) (*request.Request, *datapipeline.AddTagsOutput)

	AddTags(*datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error)

	CreatePipelineRequest(*datapipeline.CreatePipelineInput) (*request.Request, *datapipeline.CreatePipelineOutput)

	CreatePipeline(*datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error)

	DeactivatePipelineRequest(*datapipeline.DeactivatePipelineInput) (*request.Request, *datapipeline.DeactivatePipelineOutput)

	DeactivatePipeline(*datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error)

	DeletePipelineRequest(*datapipeline.DeletePipelineInput) (*request.Request, *datapipeline.DeletePipelineOutput)

	DeletePipeline(*datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error)

	DescribeObjectsRequest(*datapipeline.DescribeObjectsInput) (*request.Request, *datapipeline.DescribeObjectsOutput)

	DescribeObjects(*datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error)

	DescribeObjectsPages(*datapipeline.DescribeObjectsInput, func(*datapipeline.DescribeObjectsOutput, bool) bool) error

	DescribePipelinesRequest(*datapipeline.DescribePipelinesInput) (*request.Request, *datapipeline.DescribePipelinesOutput)

	DescribePipelines(*datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error)

	EvaluateExpressionRequest(*datapipeline.EvaluateExpressionInput) (*request.Request, *datapipeline.EvaluateExpressionOutput)

	EvaluateExpression(*datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error)

	GetPipelineDefinitionRequest(*datapipeline.GetPipelineDefinitionInput) (*request.Request, *datapipeline.GetPipelineDefinitionOutput)

	GetPipelineDefinition(*datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error)

	ListPipelinesRequest(*datapipeline.ListPipelinesInput) (*request.Request, *datapipeline.ListPipelinesOutput)

	ListPipelines(*datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error)

	ListPipelinesPages(*datapipeline.ListPipelinesInput, func(*datapipeline.ListPipelinesOutput, bool) bool) error

	PollForTaskRequest(*datapipeline.PollForTaskInput) (*request.Request, *datapipeline.PollForTaskOutput)

	PollForTask(*datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error)

	PutPipelineDefinitionRequest(*datapipeline.PutPipelineDefinitionInput) (*request.Request, *datapipeline.PutPipelineDefinitionOutput)

	PutPipelineDefinition(*datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error)

	QueryObjectsRequest(*datapipeline.QueryObjectsInput) (*request.Request, *datapipeline.QueryObjectsOutput)

	QueryObjects(*datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error)

	QueryObjectsPages(*datapipeline.QueryObjectsInput, func(*datapipeline.QueryObjectsOutput, bool) bool) error

	RemoveTagsRequest(*datapipeline.RemoveTagsInput) (*request.Request, *datapipeline.RemoveTagsOutput)

	RemoveTags(*datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error)

	ReportTaskProgressRequest(*datapipeline.ReportTaskProgressInput) (*request.Request, *datapipeline.ReportTaskProgressOutput)

	ReportTaskProgress(*datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error)

	ReportTaskRunnerHeartbeatRequest(*datapipeline.ReportTaskRunnerHeartbeatInput) (*request.Request, *datapipeline.ReportTaskRunnerHeartbeatOutput)

	ReportTaskRunnerHeartbeat(*datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error)

	SetStatusRequest(*datapipeline.SetStatusInput) (*request.Request, *datapipeline.SetStatusOutput)

	SetStatus(*datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error)

	SetTaskStatusRequest(*datapipeline.SetTaskStatusInput) (*request.Request, *datapipeline.SetTaskStatusOutput)

	SetTaskStatus(*datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error)

	ValidatePipelineDefinitionRequest(*datapipeline.ValidatePipelineDefinitionInput) (*request.Request, *datapipeline.ValidatePipelineDefinitionOutput)

	ValidatePipelineDefinition(*datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error)
}

var _ DataPipelineAPI = (*datapipeline.DataPipeline)(nil)
