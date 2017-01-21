// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codedeploy

import (
	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/client"
	"github.com/lental/aws-sdk-go/aws/client/metadata"
	"github.com/lental/aws-sdk-go/aws/request"
	"github.com/lental/aws-sdk-go/aws/signer/v4"
	"github.com/lental/aws-sdk-go/private/protocol/jsonrpc"
)

// Overview
//
// This reference guide provides descriptions of the AWS CodeDeploy APIs. For
// more information about AWS CodeDeploy, see the AWS CodeDeploy User Guide
// (http://docs.aws.amazon.com/codedeploy/latest/userguide).
//
// Using the APIs
//
// You can use the AWS CodeDeploy APIs to work with the following:
//
//    * Applications are unique identifiers used by AWS CodeDeploy to ensure
//    the correct combinations of revisions, deployment configurations, and
//    deployment groups are being referenced during deployments.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, list, and update
//    applications.
//
//    * Deployment configurations are sets of deployment rules and success and
//    failure conditions used by AWS CodeDeploy during deployments.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, and list deployment
//    configurations.
//
//    * Deployment groups are groups of instances to which application revisions
//    can be deployed.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, list, and update
//    deployment groups.
//
//    * Instances represent Amazon EC2 instances to which application revisions
//    are deployed. Instances are identified by their Amazon EC2 tags or Auto
//    Scaling group names. Instances belong to deployment groups.
//
// You can use the AWS CodeDeploy APIs to get and list instance.
//
//    * Deployments represent the process of deploying revisions to instances.
//
// You can use the AWS CodeDeploy APIs to create, get, list, and stop deployments.
//
//    * Application revisions are archive files stored in Amazon S3 buckets
//    or GitHub repositories. These revisions contain source content (such as
//    source code, web pages, executable files, and deployment scripts) along
//    with an application specification (AppSpec) file. (The AppSpec file is
//    unique to AWS CodeDeploy; it defines the deployment actions you want AWS
//    CodeDeploy to execute.) For application revisions stored in Amazon S3
//    buckets, an application revision is uniquely identified by its Amazon
//    S3 object key and its ETag, version, or both. For application revisions
//    stored in GitHub repositories, an application revision is uniquely identified
//    by its repository name and commit ID. Application revisions are deployed
//    through deployment groups.
//
// You can use the AWS CodeDeploy APIs to get, list, and register application
//    revisions.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CodeDeploy struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "codedeploy"

// New creates a new instance of the CodeDeploy client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CodeDeploy client from just a session.
//     svc := codedeploy.New(mySession)
//
//     // Create a CodeDeploy client with additional configuration
//     svc := codedeploy.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CodeDeploy {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CodeDeploy {
	svc := &CodeDeploy{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-10-06",
				JSONVersion:   "1.1",
				TargetPrefix:  "CodeDeploy_20141006",
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

// newRequest creates a new request for a CodeDeploy operation and runs any
// custom request initialization.
func (c *CodeDeploy) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
