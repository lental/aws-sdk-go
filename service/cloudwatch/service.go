// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatch

import (
	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/client"
	"github.com/lental/aws-sdk-go/aws/client/metadata"
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/aws/signer/v4"
	"github.com/lental/aws-sdk-go/private/protocol/query"
)

// Amazon CloudWatch monitors your Amazon Web Services (AWS) resources and the
// applications you run on AWS in real-time. You can use CloudWatch to collect
// and track metrics, which are the variables you want to measure for your resources
// and applications.
//
// CloudWatch alarms send notifications or automatically make changes to the
// resources you are monitoring based on rules that you define. For example,
// you can monitor the CPU usage and disk reads and writes of your Amazon Elastic
// Compute Cloud (Amazon EC2) instances and then use this data to determine
// whether you should launch additional instances to handle increased load.
// You can also use this data to stop under-used instances to save money.
//
// In addition to monitoring the built-in metrics that come with AWS, you can
// monitor your own custom metrics. With CloudWatch, you gain system-wide visibility
// into resource utilization, application performance, and operational health.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CloudWatch struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "monitoring"

// New creates a new instance of the CloudWatch client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CloudWatch client from just a session.
//     svc := cloudwatch.New(mySession)
//
//     // Create a CloudWatch client with additional configuration
//     svc := cloudwatch.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CloudWatch {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CloudWatch {
	svc := &CloudWatch{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2010-08-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CloudWatch operation and runs any
// custom request initialization.
func (c *CloudWatch) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
