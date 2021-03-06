// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package codepipelineiface provides an interface to enable mocking the AWS CodePipeline service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codepipelineiface

import (
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/service/codepipeline"
)

// CodePipelineAPI provides an interface to enable mocking the
// codepipeline.CodePipeline service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodePipeline.
//    func myFunc(svc codepipelineiface.CodePipelineAPI) bool {
//        // Make svc.AcknowledgeJob request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codepipeline.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodePipelineClient struct {
//        codepipelineiface.CodePipelineAPI
//    }
//    func (m *mockCodePipelineClient) AcknowledgeJob(input *codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodePipelineClient{}
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
type CodePipelineAPI interface {
	AcknowledgeJobRequest(*codepipeline.AcknowledgeJobInput) (*request.Request, *codepipeline.AcknowledgeJobOutput)

	AcknowledgeJob(*codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error)

	AcknowledgeThirdPartyJobRequest(*codepipeline.AcknowledgeThirdPartyJobInput) (*request.Request, *codepipeline.AcknowledgeThirdPartyJobOutput)

	AcknowledgeThirdPartyJob(*codepipeline.AcknowledgeThirdPartyJobInput) (*codepipeline.AcknowledgeThirdPartyJobOutput, error)

	CreateCustomActionTypeRequest(*codepipeline.CreateCustomActionTypeInput) (*request.Request, *codepipeline.CreateCustomActionTypeOutput)

	CreateCustomActionType(*codepipeline.CreateCustomActionTypeInput) (*codepipeline.CreateCustomActionTypeOutput, error)

	CreatePipelineRequest(*codepipeline.CreatePipelineInput) (*request.Request, *codepipeline.CreatePipelineOutput)

	CreatePipeline(*codepipeline.CreatePipelineInput) (*codepipeline.CreatePipelineOutput, error)

	DeleteCustomActionTypeRequest(*codepipeline.DeleteCustomActionTypeInput) (*request.Request, *codepipeline.DeleteCustomActionTypeOutput)

	DeleteCustomActionType(*codepipeline.DeleteCustomActionTypeInput) (*codepipeline.DeleteCustomActionTypeOutput, error)

	DeletePipelineRequest(*codepipeline.DeletePipelineInput) (*request.Request, *codepipeline.DeletePipelineOutput)

	DeletePipeline(*codepipeline.DeletePipelineInput) (*codepipeline.DeletePipelineOutput, error)

	DisableStageTransitionRequest(*codepipeline.DisableStageTransitionInput) (*request.Request, *codepipeline.DisableStageTransitionOutput)

	DisableStageTransition(*codepipeline.DisableStageTransitionInput) (*codepipeline.DisableStageTransitionOutput, error)

	EnableStageTransitionRequest(*codepipeline.EnableStageTransitionInput) (*request.Request, *codepipeline.EnableStageTransitionOutput)

	EnableStageTransition(*codepipeline.EnableStageTransitionInput) (*codepipeline.EnableStageTransitionOutput, error)

	GetJobDetailsRequest(*codepipeline.GetJobDetailsInput) (*request.Request, *codepipeline.GetJobDetailsOutput)

	GetJobDetails(*codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error)

	GetPipelineRequest(*codepipeline.GetPipelineInput) (*request.Request, *codepipeline.GetPipelineOutput)

	GetPipeline(*codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error)

	GetPipelineExecutionRequest(*codepipeline.GetPipelineExecutionInput) (*request.Request, *codepipeline.GetPipelineExecutionOutput)

	GetPipelineExecution(*codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error)

	GetPipelineStateRequest(*codepipeline.GetPipelineStateInput) (*request.Request, *codepipeline.GetPipelineStateOutput)

	GetPipelineState(*codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error)

	GetThirdPartyJobDetailsRequest(*codepipeline.GetThirdPartyJobDetailsInput) (*request.Request, *codepipeline.GetThirdPartyJobDetailsOutput)

	GetThirdPartyJobDetails(*codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error)

	ListActionTypesRequest(*codepipeline.ListActionTypesInput) (*request.Request, *codepipeline.ListActionTypesOutput)

	ListActionTypes(*codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error)

	ListPipelinesRequest(*codepipeline.ListPipelinesInput) (*request.Request, *codepipeline.ListPipelinesOutput)

	ListPipelines(*codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error)

	PollForJobsRequest(*codepipeline.PollForJobsInput) (*request.Request, *codepipeline.PollForJobsOutput)

	PollForJobs(*codepipeline.PollForJobsInput) (*codepipeline.PollForJobsOutput, error)

	PollForThirdPartyJobsRequest(*codepipeline.PollForThirdPartyJobsInput) (*request.Request, *codepipeline.PollForThirdPartyJobsOutput)

	PollForThirdPartyJobs(*codepipeline.PollForThirdPartyJobsInput) (*codepipeline.PollForThirdPartyJobsOutput, error)

	PutActionRevisionRequest(*codepipeline.PutActionRevisionInput) (*request.Request, *codepipeline.PutActionRevisionOutput)

	PutActionRevision(*codepipeline.PutActionRevisionInput) (*codepipeline.PutActionRevisionOutput, error)

	PutApprovalResultRequest(*codepipeline.PutApprovalResultInput) (*request.Request, *codepipeline.PutApprovalResultOutput)

	PutApprovalResult(*codepipeline.PutApprovalResultInput) (*codepipeline.PutApprovalResultOutput, error)

	PutJobFailureResultRequest(*codepipeline.PutJobFailureResultInput) (*request.Request, *codepipeline.PutJobFailureResultOutput)

	PutJobFailureResult(*codepipeline.PutJobFailureResultInput) (*codepipeline.PutJobFailureResultOutput, error)

	PutJobSuccessResultRequest(*codepipeline.PutJobSuccessResultInput) (*request.Request, *codepipeline.PutJobSuccessResultOutput)

	PutJobSuccessResult(*codepipeline.PutJobSuccessResultInput) (*codepipeline.PutJobSuccessResultOutput, error)

	PutThirdPartyJobFailureResultRequest(*codepipeline.PutThirdPartyJobFailureResultInput) (*request.Request, *codepipeline.PutThirdPartyJobFailureResultOutput)

	PutThirdPartyJobFailureResult(*codepipeline.PutThirdPartyJobFailureResultInput) (*codepipeline.PutThirdPartyJobFailureResultOutput, error)

	PutThirdPartyJobSuccessResultRequest(*codepipeline.PutThirdPartyJobSuccessResultInput) (*request.Request, *codepipeline.PutThirdPartyJobSuccessResultOutput)

	PutThirdPartyJobSuccessResult(*codepipeline.PutThirdPartyJobSuccessResultInput) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error)

	RetryStageExecutionRequest(*codepipeline.RetryStageExecutionInput) (*request.Request, *codepipeline.RetryStageExecutionOutput)

	RetryStageExecution(*codepipeline.RetryStageExecutionInput) (*codepipeline.RetryStageExecutionOutput, error)

	StartPipelineExecutionRequest(*codepipeline.StartPipelineExecutionInput) (*request.Request, *codepipeline.StartPipelineExecutionOutput)

	StartPipelineExecution(*codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error)

	UpdatePipelineRequest(*codepipeline.UpdatePipelineInput) (*request.Request, *codepipeline.UpdatePipelineOutput)

	UpdatePipeline(*codepipeline.UpdatePipelineInput) (*codepipeline.UpdatePipelineOutput, error)
}

var _ CodePipelineAPI = (*codepipeline.CodePipeline)(nil)
