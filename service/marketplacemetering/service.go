// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package marketplacemetering

import (
	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/client"
	"github.com/lental/aws-sdk-go/aws/client/metadata"
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/aws/signer/v4"
	"github.com/lental/aws-sdk-go/private/protocol/jsonrpc"
)

// This reference provides descriptions of the low-level AWS Marketplace Metering
// Service API.
//
// AWS Marketplace sellers can use this API to submit usage data for custom
// usage dimensions.
//
// Submitting Metering Records
//
//    * MeterUsage- Submits the metering record for a Marketplace product. MeterUsage
//    is called from an EC2 instance.
//
//    * BatchMeterUsage- Submits the metering record for a set of customers.
//    BatchMeterUsage is called from a software-as-a-service (SaaS) application.
//
// Accepting New Customers
//
//    * ResolveCustomer- Called by a SaaS application during the registration
//    process. When a buyer visits your website during the registration process,
//    the buyer submits a Registration Token through the browser. The Registration
//    Token is resolved through this API to obtain a CustomerIdentifier and
//    Product Code.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type MarketplaceMetering struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "metering.marketplace"

// New creates a new instance of the MarketplaceMetering client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a MarketplaceMetering client from just a session.
//     svc := marketplacemetering.New(mySession)
//
//     // Create a MarketplaceMetering client with additional configuration
//     svc := marketplacemetering.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *MarketplaceMetering {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *MarketplaceMetering {
	if len(signingName) == 0 {
		signingName = "aws-marketplace"
	}
	svc := &MarketplaceMetering{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-01-14",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSMPMeteringService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a MarketplaceMetering operation and runs any
// custom request initialization.
func (c *MarketplaceMetering) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
