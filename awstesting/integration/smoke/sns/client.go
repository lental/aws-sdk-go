// +build integration

//Package sns provides gucumber integration tests support.
package sns

import (
	"github.com/lental/aws-sdk-go/awstesting/integration/smoke"
	"github.com/lental/aws-sdk-go/service/sns"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sns", func() {
		gucumber.World["client"] = sns.New(smoke.Session)
	})
}
