// +build integration

//Package directconnect provides gucumber integration tests support.
package directconnect

import (
	"github.com/lental/aws-sdk-go/awstesting/integration/smoke"
	"github.com/lental/aws-sdk-go/service/directconnect"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@directconnect", func() {
		gucumber.World["client"] = directconnect.New(smoke.Session)
	})
}
