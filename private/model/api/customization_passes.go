// +build codegen

package api

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type service struct {
	srcName string
	dstName string

	serviceVersion string
}

var mergeServices = map[string]service{
	"dynamodbstreams": service{
		dstName: "dynamodb",
		srcName: "streams.dynamodb",
	},
	"wafregional": service{
		dstName:        "waf",
		srcName:        "waf-regional",
		serviceVersion: "2015-08-24",
	},
}

// customizationPasses Executes customization logic for the API by package name.
func (a *API) customizationPasses() {
	var svcCustomizations = map[string]func(*API){
		"s3":         s3Customizations,
		"cloudfront": cloudfrontCustomizations,
	}

	for k, _ := range mergeServices {
		svcCustomizations[k] = mergeServicesCustomizations
	}

	if fn := svcCustomizations[a.PackageName()]; fn != nil {
		fn(a)
	}

	blobDocStringCustomizations(a)
}

const base64MarshalDocStr = "// %s is automatically base64 encoded/decoded by the SDK.\n"

func blobDocStringCustomizations(a *API) {
	for _, s := range a.Shapes {
		payloadMemberName := s.Payload

		for refName, ref := range s.MemberRefs {
			if refName == payloadMemberName {
				// Payload members have their own encoding and may
				// be raw bytes or io.Reader
				continue
			}
			if ref.Shape.Type == "blob" {
				docStr := fmt.Sprintf(base64MarshalDocStr, refName)
				if len(strings.TrimSpace(ref.Shape.Documentation)) != 0 {
					ref.Shape.Documentation += "//\n" + docStr
				} else if len(strings.TrimSpace(ref.Documentation)) != 0 {
					ref.Documentation += "//\n" + docStr
				} else {
					ref.Documentation = docStr
				}
			}
		}
	}
}

// s3Customizations customizes the API generation to replace values specific to S3.
func s3Customizations(a *API) {
	var strExpires *Shape

	for name, s := range a.Shapes {
		// Remove ContentMD5 members
		if _, ok := s.MemberRefs["ContentMD5"]; ok {
			delete(s.MemberRefs, "ContentMD5")
		}

		// Expires should be a string not time.Time since the format is not
		// enforced by S3, and any value can be set to this field outside of the SDK.
		if strings.HasSuffix(name, "Output") {
			if ref, ok := s.MemberRefs["Expires"]; ok {
				if strExpires == nil {
					newShape := *ref.Shape
					strExpires = &newShape
					strExpires.Type = "string"
					strExpires.refs = []*ShapeRef{}
				}
				ref.Shape.removeRef(ref)
				ref.Shape = strExpires
				ref.Shape.refs = append(ref.Shape.refs, &s.MemberRef)
			}
		}
	}
}

// cloudfrontCustomizations customized the API generation to replace values
// specific to CloudFront.
func cloudfrontCustomizations(a *API) {
	// MaxItems members should always be integers
	for _, s := range a.Shapes {
		if ref, ok := s.MemberRefs["MaxItems"]; ok {
			ref.ShapeName = "Integer"
			ref.Shape = a.Shapes["Integer"]
		}
	}
}

// mergeServicesCustomizations references any duplicate shapes from DynamoDB
func mergeServicesCustomizations(a *API) {
	info := mergeServices[a.PackageName()]

	p := strings.Replace(a.path, info.srcName, info.dstName, -1)

	if info.serviceVersion != "" {
		index := strings.LastIndex(p, "/")
		files, _ := ioutil.ReadDir(p[:index])
		if len(files) > 1 {
			panic("New version was introduced")
		}
		p = p[:index] + "/" + info.serviceVersion
	}

	file := filepath.Join(p, "api-2.json")

	serviceAPI := API{}
	serviceAPI.Attach(file)
	serviceAPI.Setup()

	for n := range a.Shapes {
		if _, ok := serviceAPI.Shapes[n]; ok {
			a.Shapes[n].resolvePkg = "github.com/lental/aws-sdk-go/service/" + info.dstName
		}
	}
}
