package awstesting

import (
	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/client"
	"github.com/lental/aws-sdk-go/aws/client/metadata"
	"github.com/lental/aws-sdk-go/aws/defaults"
)

// NewClient creates and initializes a generic service client for testing.
func NewClient(cfgs ...*aws.Config) *client.Client {
	info := metadata.ClientInfo{
		Endpoint:    "http://endpoint",
		SigningName: "",
	}
	def := defaults.Get()
	def.Config.MergeIn(cfgs...)

	return client.New(*def.Config, info, def.Handlers)
}
