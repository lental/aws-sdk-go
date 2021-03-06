// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticache

import (
	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/client"
	"github.com/lental/aws-sdk-go/aws/client/metadata"
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/aws/signer/v4"
	"github.com/lental/aws-sdk-go/private/protocol/query"
)

// Amazon ElastiCache is a web service that makes it easier to set up, operate,
// and scale a distributed cache in the cloud.
//
// With ElastiCache, customers get all of the benefits of a high-performance,
// in-memory cache with less of the administrative burden involved in launching
// and managing a distributed cache. The service makes setup, scaling, and cluster
// failure handling much simpler than in a self-managed cache deployment.
//
// In addition, through integration with Amazon CloudWatch, customers get enhanced
// visibility into the key performance statistics associated with their cache
// and can receive alarms if a part of their cache runs hot.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ElastiCache struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "elasticache"

// New creates a new instance of the ElastiCache client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ElastiCache client from just a session.
//     svc := elasticache.New(mySession)
//
//     // Create a ElastiCache client with additional configuration
//     svc := elasticache.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ElastiCache {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ElastiCache {
	svc := &ElastiCache{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-02-02",
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

// newRequest creates a new request for a ElastiCache operation and runs any
// custom request initialization.
func (c *ElastiCache) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
