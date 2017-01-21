// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package route53iface provides an interface to enable mocking the Amazon Route 53 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package route53iface

import (
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/service/route53"
)

// Route53API provides an interface to enable mocking the
// route53.Route53 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Route 53.
//    func myFunc(svc route53iface.Route53API) bool {
//        // Make svc.AssociateVPCWithHostedZone request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := route53.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRoute53Client struct {
//        route53iface.Route53API
//    }
//    func (m *mockRoute53Client) AssociateVPCWithHostedZone(input *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRoute53Client{}
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
type Route53API interface {
	AssociateVPCWithHostedZoneRequest(*route53.AssociateVPCWithHostedZoneInput) (*request.Request, *route53.AssociateVPCWithHostedZoneOutput)

	AssociateVPCWithHostedZone(*route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error)

	ChangeResourceRecordSetsRequest(*route53.ChangeResourceRecordSetsInput) (*request.Request, *route53.ChangeResourceRecordSetsOutput)

	ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)

	ChangeTagsForResourceRequest(*route53.ChangeTagsForResourceInput) (*request.Request, *route53.ChangeTagsForResourceOutput)

	ChangeTagsForResource(*route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error)

	CreateHealthCheckRequest(*route53.CreateHealthCheckInput) (*request.Request, *route53.CreateHealthCheckOutput)

	CreateHealthCheck(*route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error)

	CreateHostedZoneRequest(*route53.CreateHostedZoneInput) (*request.Request, *route53.CreateHostedZoneOutput)

	CreateHostedZone(*route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)

	CreateReusableDelegationSetRequest(*route53.CreateReusableDelegationSetInput) (*request.Request, *route53.CreateReusableDelegationSetOutput)

	CreateReusableDelegationSet(*route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error)

	CreateTrafficPolicyRequest(*route53.CreateTrafficPolicyInput) (*request.Request, *route53.CreateTrafficPolicyOutput)

	CreateTrafficPolicy(*route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error)

	CreateTrafficPolicyInstanceRequest(*route53.CreateTrafficPolicyInstanceInput) (*request.Request, *route53.CreateTrafficPolicyInstanceOutput)

	CreateTrafficPolicyInstance(*route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error)

	CreateTrafficPolicyVersionRequest(*route53.CreateTrafficPolicyVersionInput) (*request.Request, *route53.CreateTrafficPolicyVersionOutput)

	CreateTrafficPolicyVersion(*route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error)

	CreateVPCAssociationAuthorizationRequest(*route53.CreateVPCAssociationAuthorizationInput) (*request.Request, *route53.CreateVPCAssociationAuthorizationOutput)

	CreateVPCAssociationAuthorization(*route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error)

	DeleteHealthCheckRequest(*route53.DeleteHealthCheckInput) (*request.Request, *route53.DeleteHealthCheckOutput)

	DeleteHealthCheck(*route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error)

	DeleteHostedZoneRequest(*route53.DeleteHostedZoneInput) (*request.Request, *route53.DeleteHostedZoneOutput)

	DeleteHostedZone(*route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error)

	DeleteReusableDelegationSetRequest(*route53.DeleteReusableDelegationSetInput) (*request.Request, *route53.DeleteReusableDelegationSetOutput)

	DeleteReusableDelegationSet(*route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error)

	DeleteTrafficPolicyRequest(*route53.DeleteTrafficPolicyInput) (*request.Request, *route53.DeleteTrafficPolicyOutput)

	DeleteTrafficPolicy(*route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error)

	DeleteTrafficPolicyInstanceRequest(*route53.DeleteTrafficPolicyInstanceInput) (*request.Request, *route53.DeleteTrafficPolicyInstanceOutput)

	DeleteTrafficPolicyInstance(*route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error)

	DeleteVPCAssociationAuthorizationRequest(*route53.DeleteVPCAssociationAuthorizationInput) (*request.Request, *route53.DeleteVPCAssociationAuthorizationOutput)

	DeleteVPCAssociationAuthorization(*route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error)

	DisassociateVPCFromHostedZoneRequest(*route53.DisassociateVPCFromHostedZoneInput) (*request.Request, *route53.DisassociateVPCFromHostedZoneOutput)

	DisassociateVPCFromHostedZone(*route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error)

	GetChangeRequest(*route53.GetChangeInput) (*request.Request, *route53.GetChangeOutput)

	GetChange(*route53.GetChangeInput) (*route53.GetChangeOutput, error)

	GetCheckerIpRangesRequest(*route53.GetCheckerIpRangesInput) (*request.Request, *route53.GetCheckerIpRangesOutput)

	GetCheckerIpRanges(*route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error)

	GetGeoLocationRequest(*route53.GetGeoLocationInput) (*request.Request, *route53.GetGeoLocationOutput)

	GetGeoLocation(*route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error)

	GetHealthCheckRequest(*route53.GetHealthCheckInput) (*request.Request, *route53.GetHealthCheckOutput)

	GetHealthCheck(*route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error)

	GetHealthCheckCountRequest(*route53.GetHealthCheckCountInput) (*request.Request, *route53.GetHealthCheckCountOutput)

	GetHealthCheckCount(*route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error)

	GetHealthCheckLastFailureReasonRequest(*route53.GetHealthCheckLastFailureReasonInput) (*request.Request, *route53.GetHealthCheckLastFailureReasonOutput)

	GetHealthCheckLastFailureReason(*route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error)

	GetHealthCheckStatusRequest(*route53.GetHealthCheckStatusInput) (*request.Request, *route53.GetHealthCheckStatusOutput)

	GetHealthCheckStatus(*route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error)

	GetHostedZoneRequest(*route53.GetHostedZoneInput) (*request.Request, *route53.GetHostedZoneOutput)

	GetHostedZone(*route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error)

	GetHostedZoneCountRequest(*route53.GetHostedZoneCountInput) (*request.Request, *route53.GetHostedZoneCountOutput)

	GetHostedZoneCount(*route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error)

	GetReusableDelegationSetRequest(*route53.GetReusableDelegationSetInput) (*request.Request, *route53.GetReusableDelegationSetOutput)

	GetReusableDelegationSet(*route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error)

	GetTrafficPolicyRequest(*route53.GetTrafficPolicyInput) (*request.Request, *route53.GetTrafficPolicyOutput)

	GetTrafficPolicy(*route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error)

	GetTrafficPolicyInstanceRequest(*route53.GetTrafficPolicyInstanceInput) (*request.Request, *route53.GetTrafficPolicyInstanceOutput)

	GetTrafficPolicyInstance(*route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error)

	GetTrafficPolicyInstanceCountRequest(*route53.GetTrafficPolicyInstanceCountInput) (*request.Request, *route53.GetTrafficPolicyInstanceCountOutput)

	GetTrafficPolicyInstanceCount(*route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error)

	ListGeoLocationsRequest(*route53.ListGeoLocationsInput) (*request.Request, *route53.ListGeoLocationsOutput)

	ListGeoLocations(*route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error)

	ListHealthChecksRequest(*route53.ListHealthChecksInput) (*request.Request, *route53.ListHealthChecksOutput)

	ListHealthChecks(*route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error)

	ListHealthChecksPages(*route53.ListHealthChecksInput, func(*route53.ListHealthChecksOutput, bool) bool) error

	ListHostedZonesRequest(*route53.ListHostedZonesInput) (*request.Request, *route53.ListHostedZonesOutput)

	ListHostedZones(*route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error)

	ListHostedZonesPages(*route53.ListHostedZonesInput, func(*route53.ListHostedZonesOutput, bool) bool) error

	ListHostedZonesByNameRequest(*route53.ListHostedZonesByNameInput) (*request.Request, *route53.ListHostedZonesByNameOutput)

	ListHostedZonesByName(*route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error)

	ListResourceRecordSetsRequest(*route53.ListResourceRecordSetsInput) (*request.Request, *route53.ListResourceRecordSetsOutput)

	ListResourceRecordSets(*route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error)

	ListResourceRecordSetsPages(*route53.ListResourceRecordSetsInput, func(*route53.ListResourceRecordSetsOutput, bool) bool) error

	ListReusableDelegationSetsRequest(*route53.ListReusableDelegationSetsInput) (*request.Request, *route53.ListReusableDelegationSetsOutput)

	ListReusableDelegationSets(*route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error)

	ListTagsForResourceRequest(*route53.ListTagsForResourceInput) (*request.Request, *route53.ListTagsForResourceOutput)

	ListTagsForResource(*route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error)

	ListTagsForResourcesRequest(*route53.ListTagsForResourcesInput) (*request.Request, *route53.ListTagsForResourcesOutput)

	ListTagsForResources(*route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error)

	ListTrafficPoliciesRequest(*route53.ListTrafficPoliciesInput) (*request.Request, *route53.ListTrafficPoliciesOutput)

	ListTrafficPolicies(*route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error)

	ListTrafficPolicyInstancesRequest(*route53.ListTrafficPolicyInstancesInput) (*request.Request, *route53.ListTrafficPolicyInstancesOutput)

	ListTrafficPolicyInstances(*route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error)

	ListTrafficPolicyInstancesByHostedZoneRequest(*route53.ListTrafficPolicyInstancesByHostedZoneInput) (*request.Request, *route53.ListTrafficPolicyInstancesByHostedZoneOutput)

	ListTrafficPolicyInstancesByHostedZone(*route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)

	ListTrafficPolicyInstancesByPolicyRequest(*route53.ListTrafficPolicyInstancesByPolicyInput) (*request.Request, *route53.ListTrafficPolicyInstancesByPolicyOutput)

	ListTrafficPolicyInstancesByPolicy(*route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error)

	ListTrafficPolicyVersionsRequest(*route53.ListTrafficPolicyVersionsInput) (*request.Request, *route53.ListTrafficPolicyVersionsOutput)

	ListTrafficPolicyVersions(*route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error)

	ListVPCAssociationAuthorizationsRequest(*route53.ListVPCAssociationAuthorizationsInput) (*request.Request, *route53.ListVPCAssociationAuthorizationsOutput)

	ListVPCAssociationAuthorizations(*route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error)

	TestDNSAnswerRequest(*route53.TestDNSAnswerInput) (*request.Request, *route53.TestDNSAnswerOutput)

	TestDNSAnswer(*route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error)

	UpdateHealthCheckRequest(*route53.UpdateHealthCheckInput) (*request.Request, *route53.UpdateHealthCheckOutput)

	UpdateHealthCheck(*route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error)

	UpdateHostedZoneCommentRequest(*route53.UpdateHostedZoneCommentInput) (*request.Request, *route53.UpdateHostedZoneCommentOutput)

	UpdateHostedZoneComment(*route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error)

	UpdateTrafficPolicyCommentRequest(*route53.UpdateTrafficPolicyCommentInput) (*request.Request, *route53.UpdateTrafficPolicyCommentOutput)

	UpdateTrafficPolicyComment(*route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error)

	UpdateTrafficPolicyInstanceRequest(*route53.UpdateTrafficPolicyInstanceInput) (*request.Request, *route53.UpdateTrafficPolicyInstanceOutput)

	UpdateTrafficPolicyInstance(*route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error)

	WaitUntilResourceRecordSetsChanged(*route53.GetChangeInput) error
}

var _ Route53API = (*route53.Route53)(nil)
