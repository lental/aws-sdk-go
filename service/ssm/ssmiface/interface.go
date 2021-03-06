// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ssmiface provides an interface to enable mocking the Amazon Simple Systems Manager (SSM) service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ssmiface

import (
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/service/ssm"
)

// SSMAPI provides an interface to enable mocking the
// ssm.SSM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Simple Systems Manager (SSM).
//    func myFunc(svc ssmiface.SSMAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ssm.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSSMClient struct {
//        ssmiface.SSMAPI
//    }
//    func (m *mockSSMClient) AddTagsToResource(input *ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSSMClient{}
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
type SSMAPI interface {
	AddTagsToResourceRequest(*ssm.AddTagsToResourceInput) (*request.Request, *ssm.AddTagsToResourceOutput)

	AddTagsToResource(*ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error)

	CancelCommandRequest(*ssm.CancelCommandInput) (*request.Request, *ssm.CancelCommandOutput)

	CancelCommand(*ssm.CancelCommandInput) (*ssm.CancelCommandOutput, error)

	CreateActivationRequest(*ssm.CreateActivationInput) (*request.Request, *ssm.CreateActivationOutput)

	CreateActivation(*ssm.CreateActivationInput) (*ssm.CreateActivationOutput, error)

	CreateAssociationRequest(*ssm.CreateAssociationInput) (*request.Request, *ssm.CreateAssociationOutput)

	CreateAssociation(*ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error)

	CreateAssociationBatchRequest(*ssm.CreateAssociationBatchInput) (*request.Request, *ssm.CreateAssociationBatchOutput)

	CreateAssociationBatch(*ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error)

	CreateDocumentRequest(*ssm.CreateDocumentInput) (*request.Request, *ssm.CreateDocumentOutput)

	CreateDocument(*ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error)

	CreateMaintenanceWindowRequest(*ssm.CreateMaintenanceWindowInput) (*request.Request, *ssm.CreateMaintenanceWindowOutput)

	CreateMaintenanceWindow(*ssm.CreateMaintenanceWindowInput) (*ssm.CreateMaintenanceWindowOutput, error)

	DeleteActivationRequest(*ssm.DeleteActivationInput) (*request.Request, *ssm.DeleteActivationOutput)

	DeleteActivation(*ssm.DeleteActivationInput) (*ssm.DeleteActivationOutput, error)

	DeleteAssociationRequest(*ssm.DeleteAssociationInput) (*request.Request, *ssm.DeleteAssociationOutput)

	DeleteAssociation(*ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error)

	DeleteDocumentRequest(*ssm.DeleteDocumentInput) (*request.Request, *ssm.DeleteDocumentOutput)

	DeleteDocument(*ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error)

	DeleteMaintenanceWindowRequest(*ssm.DeleteMaintenanceWindowInput) (*request.Request, *ssm.DeleteMaintenanceWindowOutput)

	DeleteMaintenanceWindow(*ssm.DeleteMaintenanceWindowInput) (*ssm.DeleteMaintenanceWindowOutput, error)

	DeleteParameterRequest(*ssm.DeleteParameterInput) (*request.Request, *ssm.DeleteParameterOutput)

	DeleteParameter(*ssm.DeleteParameterInput) (*ssm.DeleteParameterOutput, error)

	DeregisterManagedInstanceRequest(*ssm.DeregisterManagedInstanceInput) (*request.Request, *ssm.DeregisterManagedInstanceOutput)

	DeregisterManagedInstance(*ssm.DeregisterManagedInstanceInput) (*ssm.DeregisterManagedInstanceOutput, error)

	DeregisterTargetFromMaintenanceWindowRequest(*ssm.DeregisterTargetFromMaintenanceWindowInput) (*request.Request, *ssm.DeregisterTargetFromMaintenanceWindowOutput)

	DeregisterTargetFromMaintenanceWindow(*ssm.DeregisterTargetFromMaintenanceWindowInput) (*ssm.DeregisterTargetFromMaintenanceWindowOutput, error)

	DeregisterTaskFromMaintenanceWindowRequest(*ssm.DeregisterTaskFromMaintenanceWindowInput) (*request.Request, *ssm.DeregisterTaskFromMaintenanceWindowOutput)

	DeregisterTaskFromMaintenanceWindow(*ssm.DeregisterTaskFromMaintenanceWindowInput) (*ssm.DeregisterTaskFromMaintenanceWindowOutput, error)

	DescribeActivationsRequest(*ssm.DescribeActivationsInput) (*request.Request, *ssm.DescribeActivationsOutput)

	DescribeActivations(*ssm.DescribeActivationsInput) (*ssm.DescribeActivationsOutput, error)

	DescribeActivationsPages(*ssm.DescribeActivationsInput, func(*ssm.DescribeActivationsOutput, bool) bool) error

	DescribeAssociationRequest(*ssm.DescribeAssociationInput) (*request.Request, *ssm.DescribeAssociationOutput)

	DescribeAssociation(*ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error)

	DescribeAutomationExecutionsRequest(*ssm.DescribeAutomationExecutionsInput) (*request.Request, *ssm.DescribeAutomationExecutionsOutput)

	DescribeAutomationExecutions(*ssm.DescribeAutomationExecutionsInput) (*ssm.DescribeAutomationExecutionsOutput, error)

	DescribeDocumentRequest(*ssm.DescribeDocumentInput) (*request.Request, *ssm.DescribeDocumentOutput)

	DescribeDocument(*ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error)

	DescribeDocumentPermissionRequest(*ssm.DescribeDocumentPermissionInput) (*request.Request, *ssm.DescribeDocumentPermissionOutput)

	DescribeDocumentPermission(*ssm.DescribeDocumentPermissionInput) (*ssm.DescribeDocumentPermissionOutput, error)

	DescribeEffectiveInstanceAssociationsRequest(*ssm.DescribeEffectiveInstanceAssociationsInput) (*request.Request, *ssm.DescribeEffectiveInstanceAssociationsOutput)

	DescribeEffectiveInstanceAssociations(*ssm.DescribeEffectiveInstanceAssociationsInput) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error)

	DescribeInstanceAssociationsStatusRequest(*ssm.DescribeInstanceAssociationsStatusInput) (*request.Request, *ssm.DescribeInstanceAssociationsStatusOutput)

	DescribeInstanceAssociationsStatus(*ssm.DescribeInstanceAssociationsStatusInput) (*ssm.DescribeInstanceAssociationsStatusOutput, error)

	DescribeInstanceInformationRequest(*ssm.DescribeInstanceInformationInput) (*request.Request, *ssm.DescribeInstanceInformationOutput)

	DescribeInstanceInformation(*ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error)

	DescribeInstanceInformationPages(*ssm.DescribeInstanceInformationInput, func(*ssm.DescribeInstanceInformationOutput, bool) bool) error

	DescribeMaintenanceWindowExecutionTaskInvocationsRequest(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*request.Request, *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput)

	DescribeMaintenanceWindowExecutionTaskInvocations(*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error)

	DescribeMaintenanceWindowExecutionTasksRequest(*ssm.DescribeMaintenanceWindowExecutionTasksInput) (*request.Request, *ssm.DescribeMaintenanceWindowExecutionTasksOutput)

	DescribeMaintenanceWindowExecutionTasks(*ssm.DescribeMaintenanceWindowExecutionTasksInput) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error)

	DescribeMaintenanceWindowExecutionsRequest(*ssm.DescribeMaintenanceWindowExecutionsInput) (*request.Request, *ssm.DescribeMaintenanceWindowExecutionsOutput)

	DescribeMaintenanceWindowExecutions(*ssm.DescribeMaintenanceWindowExecutionsInput) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error)

	DescribeMaintenanceWindowTargetsRequest(*ssm.DescribeMaintenanceWindowTargetsInput) (*request.Request, *ssm.DescribeMaintenanceWindowTargetsOutput)

	DescribeMaintenanceWindowTargets(*ssm.DescribeMaintenanceWindowTargetsInput) (*ssm.DescribeMaintenanceWindowTargetsOutput, error)

	DescribeMaintenanceWindowTasksRequest(*ssm.DescribeMaintenanceWindowTasksInput) (*request.Request, *ssm.DescribeMaintenanceWindowTasksOutput)

	DescribeMaintenanceWindowTasks(*ssm.DescribeMaintenanceWindowTasksInput) (*ssm.DescribeMaintenanceWindowTasksOutput, error)

	DescribeMaintenanceWindowsRequest(*ssm.DescribeMaintenanceWindowsInput) (*request.Request, *ssm.DescribeMaintenanceWindowsOutput)

	DescribeMaintenanceWindows(*ssm.DescribeMaintenanceWindowsInput) (*ssm.DescribeMaintenanceWindowsOutput, error)

	DescribeParametersRequest(*ssm.DescribeParametersInput) (*request.Request, *ssm.DescribeParametersOutput)

	DescribeParameters(*ssm.DescribeParametersInput) (*ssm.DescribeParametersOutput, error)

	GetAutomationExecutionRequest(*ssm.GetAutomationExecutionInput) (*request.Request, *ssm.GetAutomationExecutionOutput)

	GetAutomationExecution(*ssm.GetAutomationExecutionInput) (*ssm.GetAutomationExecutionOutput, error)

	GetCommandInvocationRequest(*ssm.GetCommandInvocationInput) (*request.Request, *ssm.GetCommandInvocationOutput)

	GetCommandInvocation(*ssm.GetCommandInvocationInput) (*ssm.GetCommandInvocationOutput, error)

	GetDocumentRequest(*ssm.GetDocumentInput) (*request.Request, *ssm.GetDocumentOutput)

	GetDocument(*ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error)

	GetInventoryRequest(*ssm.GetInventoryInput) (*request.Request, *ssm.GetInventoryOutput)

	GetInventory(*ssm.GetInventoryInput) (*ssm.GetInventoryOutput, error)

	GetInventorySchemaRequest(*ssm.GetInventorySchemaInput) (*request.Request, *ssm.GetInventorySchemaOutput)

	GetInventorySchema(*ssm.GetInventorySchemaInput) (*ssm.GetInventorySchemaOutput, error)

	GetMaintenanceWindowRequest(*ssm.GetMaintenanceWindowInput) (*request.Request, *ssm.GetMaintenanceWindowOutput)

	GetMaintenanceWindow(*ssm.GetMaintenanceWindowInput) (*ssm.GetMaintenanceWindowOutput, error)

	GetMaintenanceWindowExecutionRequest(*ssm.GetMaintenanceWindowExecutionInput) (*request.Request, *ssm.GetMaintenanceWindowExecutionOutput)

	GetMaintenanceWindowExecution(*ssm.GetMaintenanceWindowExecutionInput) (*ssm.GetMaintenanceWindowExecutionOutput, error)

	GetMaintenanceWindowExecutionTaskRequest(*ssm.GetMaintenanceWindowExecutionTaskInput) (*request.Request, *ssm.GetMaintenanceWindowExecutionTaskOutput)

	GetMaintenanceWindowExecutionTask(*ssm.GetMaintenanceWindowExecutionTaskInput) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error)

	GetParameterHistoryRequest(*ssm.GetParameterHistoryInput) (*request.Request, *ssm.GetParameterHistoryOutput)

	GetParameterHistory(*ssm.GetParameterHistoryInput) (*ssm.GetParameterHistoryOutput, error)

	GetParametersRequest(*ssm.GetParametersInput) (*request.Request, *ssm.GetParametersOutput)

	GetParameters(*ssm.GetParametersInput) (*ssm.GetParametersOutput, error)

	ListAssociationsRequest(*ssm.ListAssociationsInput) (*request.Request, *ssm.ListAssociationsOutput)

	ListAssociations(*ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error)

	ListAssociationsPages(*ssm.ListAssociationsInput, func(*ssm.ListAssociationsOutput, bool) bool) error

	ListCommandInvocationsRequest(*ssm.ListCommandInvocationsInput) (*request.Request, *ssm.ListCommandInvocationsOutput)

	ListCommandInvocations(*ssm.ListCommandInvocationsInput) (*ssm.ListCommandInvocationsOutput, error)

	ListCommandInvocationsPages(*ssm.ListCommandInvocationsInput, func(*ssm.ListCommandInvocationsOutput, bool) bool) error

	ListCommandsRequest(*ssm.ListCommandsInput) (*request.Request, *ssm.ListCommandsOutput)

	ListCommands(*ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error)

	ListCommandsPages(*ssm.ListCommandsInput, func(*ssm.ListCommandsOutput, bool) bool) error

	ListDocumentVersionsRequest(*ssm.ListDocumentVersionsInput) (*request.Request, *ssm.ListDocumentVersionsOutput)

	ListDocumentVersions(*ssm.ListDocumentVersionsInput) (*ssm.ListDocumentVersionsOutput, error)

	ListDocumentsRequest(*ssm.ListDocumentsInput) (*request.Request, *ssm.ListDocumentsOutput)

	ListDocuments(*ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error)

	ListDocumentsPages(*ssm.ListDocumentsInput, func(*ssm.ListDocumentsOutput, bool) bool) error

	ListInventoryEntriesRequest(*ssm.ListInventoryEntriesInput) (*request.Request, *ssm.ListInventoryEntriesOutput)

	ListInventoryEntries(*ssm.ListInventoryEntriesInput) (*ssm.ListInventoryEntriesOutput, error)

	ListTagsForResourceRequest(*ssm.ListTagsForResourceInput) (*request.Request, *ssm.ListTagsForResourceOutput)

	ListTagsForResource(*ssm.ListTagsForResourceInput) (*ssm.ListTagsForResourceOutput, error)

	ModifyDocumentPermissionRequest(*ssm.ModifyDocumentPermissionInput) (*request.Request, *ssm.ModifyDocumentPermissionOutput)

	ModifyDocumentPermission(*ssm.ModifyDocumentPermissionInput) (*ssm.ModifyDocumentPermissionOutput, error)

	PutInventoryRequest(*ssm.PutInventoryInput) (*request.Request, *ssm.PutInventoryOutput)

	PutInventory(*ssm.PutInventoryInput) (*ssm.PutInventoryOutput, error)

	PutParameterRequest(*ssm.PutParameterInput) (*request.Request, *ssm.PutParameterOutput)

	PutParameter(*ssm.PutParameterInput) (*ssm.PutParameterOutput, error)

	RegisterTargetWithMaintenanceWindowRequest(*ssm.RegisterTargetWithMaintenanceWindowInput) (*request.Request, *ssm.RegisterTargetWithMaintenanceWindowOutput)

	RegisterTargetWithMaintenanceWindow(*ssm.RegisterTargetWithMaintenanceWindowInput) (*ssm.RegisterTargetWithMaintenanceWindowOutput, error)

	RegisterTaskWithMaintenanceWindowRequest(*ssm.RegisterTaskWithMaintenanceWindowInput) (*request.Request, *ssm.RegisterTaskWithMaintenanceWindowOutput)

	RegisterTaskWithMaintenanceWindow(*ssm.RegisterTaskWithMaintenanceWindowInput) (*ssm.RegisterTaskWithMaintenanceWindowOutput, error)

	RemoveTagsFromResourceRequest(*ssm.RemoveTagsFromResourceInput) (*request.Request, *ssm.RemoveTagsFromResourceOutput)

	RemoveTagsFromResource(*ssm.RemoveTagsFromResourceInput) (*ssm.RemoveTagsFromResourceOutput, error)

	SendCommandRequest(*ssm.SendCommandInput) (*request.Request, *ssm.SendCommandOutput)

	SendCommand(*ssm.SendCommandInput) (*ssm.SendCommandOutput, error)

	StartAutomationExecutionRequest(*ssm.StartAutomationExecutionInput) (*request.Request, *ssm.StartAutomationExecutionOutput)

	StartAutomationExecution(*ssm.StartAutomationExecutionInput) (*ssm.StartAutomationExecutionOutput, error)

	StopAutomationExecutionRequest(*ssm.StopAutomationExecutionInput) (*request.Request, *ssm.StopAutomationExecutionOutput)

	StopAutomationExecution(*ssm.StopAutomationExecutionInput) (*ssm.StopAutomationExecutionOutput, error)

	UpdateAssociationRequest(*ssm.UpdateAssociationInput) (*request.Request, *ssm.UpdateAssociationOutput)

	UpdateAssociation(*ssm.UpdateAssociationInput) (*ssm.UpdateAssociationOutput, error)

	UpdateAssociationStatusRequest(*ssm.UpdateAssociationStatusInput) (*request.Request, *ssm.UpdateAssociationStatusOutput)

	UpdateAssociationStatus(*ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error)

	UpdateDocumentRequest(*ssm.UpdateDocumentInput) (*request.Request, *ssm.UpdateDocumentOutput)

	UpdateDocument(*ssm.UpdateDocumentInput) (*ssm.UpdateDocumentOutput, error)

	UpdateDocumentDefaultVersionRequest(*ssm.UpdateDocumentDefaultVersionInput) (*request.Request, *ssm.UpdateDocumentDefaultVersionOutput)

	UpdateDocumentDefaultVersion(*ssm.UpdateDocumentDefaultVersionInput) (*ssm.UpdateDocumentDefaultVersionOutput, error)

	UpdateMaintenanceWindowRequest(*ssm.UpdateMaintenanceWindowInput) (*request.Request, *ssm.UpdateMaintenanceWindowOutput)

	UpdateMaintenanceWindow(*ssm.UpdateMaintenanceWindowInput) (*ssm.UpdateMaintenanceWindowOutput, error)

	UpdateManagedInstanceRoleRequest(*ssm.UpdateManagedInstanceRoleInput) (*request.Request, *ssm.UpdateManagedInstanceRoleOutput)

	UpdateManagedInstanceRole(*ssm.UpdateManagedInstanceRoleInput) (*ssm.UpdateManagedInstanceRoleOutput, error)
}

var _ SSMAPI = (*ssm.SSM)(nil)
