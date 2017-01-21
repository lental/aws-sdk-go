// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudhsm

import (
	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/client"
	"github.com/lental/aws-sdk-go/aws/client/metadata"
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/aws/signer/v4"
	"github.com/lental/aws-sdk-go/private/protocol/jsonrpc"
)

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CloudHSM struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "cloudhsm"

// New creates a new instance of the CloudHSM client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CloudHSM client from just a session.
//     svc := cloudhsm.New(mySession)
//
//     // Create a CloudHSM client with additional configuration
//     svc := cloudhsm.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CloudHSM {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CloudHSM {
	svc := &CloudHSM{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-05-30",
				JSONVersion:   "1.1",
				TargetPrefix:  "CloudHsmFrontendService",
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

// newRequest creates a new request for a CloudHSM operation and runs any
// custom request initialization.
func (c *CloudHSM) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
