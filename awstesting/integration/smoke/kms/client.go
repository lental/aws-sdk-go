// +build integration

//Package kms provides gucumber integration tests support.
package kms

import (
	"github.com/lental/aws-sdk-go/awstesting/integration/smoke"
	"github.com/lental/aws-sdk-go/service/kms"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@kms", func() {
		gucumber.World["client"] = kms.New(smoke.Session)
	})
}
