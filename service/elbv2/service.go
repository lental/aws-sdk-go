// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elbv2

import (
	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/client"
	"github.com/lental/aws-sdk-go/aws/client/metadata"
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/aws/signer/v4"
	"github.com/lental/aws-sdk-go/private/protocol/query"
)

// A load balancer distributes incoming traffic across targets, such as your
// EC2 instances. This enables you to increase the availability of your application.
// The load balancer also monitors the health of its registered targets and
// ensures that it routes traffic only to healthy targets. You configure your
// load balancer to accept incoming traffic by specifying one or more listeners,
// which are configured with a protocol and port number for connections from
// clients to the load balancer. You configure a target group with a protocol
// and port number for connections from the load balancer to the targets, and
// with health check settings to be used when checking the health status of
// the targets.
//
// Elastic Load Balancing supports two types of load balancers: Classic Load
// Balancers and Application Load Balancers. A Classic Load Balancer makes routing
// and load balancing decisions either at the transport layer (TCP/SSL) or the
// application layer (HTTP/HTTPS), and supports either EC2-Classic or a VPC.
// An Application Load Balancer makes routing and load balancing decisions at
// the application layer (HTTP/HTTPS), supports path-based routing, and can
// route requests to one or more ports on each EC2 instance or container instance
// in your virtual private cloud (VPC). For more information, see the Elastic
// Load Balancing User Guide (http://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/).
//
// This reference covers the 2015-12-01 API, which supports Application Load
// Balancers. The 2012-06-01 API supports Classic Load Balancers.
//
// To get started, complete the following tasks:
//
// Create an Application Load Balancer using CreateLoadBalancer.
//
// Create a target group using CreateTargetGroup.
//
// Register targets for the target group using RegisterTargets.
//
// Create one or more listeners for your load balancer using CreateListener.
//
// (Optional) Create one or more rules for content routing based on URL using
// CreateRule.
//
// To delete an Application Load Balancer and its related resources, complete
// the following tasks:
//
// Delete the load balancer using DeleteLoadBalancer.
//
// Delete the target group using DeleteTargetGroup.
//
// All Elastic Load Balancing operations are idempotent, which means that they
// complete at most one time. If you repeat an operation, it succeeds.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ELBV2 struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "elasticloadbalancing"

// New creates a new instance of the ELBV2 client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ELBV2 client from just a session.
//     svc := elbv2.New(mySession)
//
//     // Create a ELBV2 client with additional configuration
//     svc := elbv2.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ELBV2 {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ELBV2 {
	svc := &ELBV2{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-12-01",
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

// newRequest creates a new request for a ELBV2 operation and runs any
// custom request initialization.
func (c *ELBV2) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
