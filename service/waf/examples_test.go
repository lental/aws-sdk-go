// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package waf_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/lental/aws-sdk-go/aws"
	"github.com/lental/aws-sdk-go/aws/session"
	"github.com/lental/aws-sdk-go/service/waf"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleWAF_CreateByteMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.CreateByteMatchSetInput{
		ChangeToken: aws.String("ChangeToken"),  // Required
		Name:        aws.String("ResourceName"), // Required
	}
	resp, err := svc.CreateByteMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_CreateIPSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.CreateIPSetInput{
		ChangeToken: aws.String("ChangeToken"),  // Required
		Name:        aws.String("ResourceName"), // Required
	}
	resp, err := svc.CreateIPSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_CreateRule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.CreateRuleInput{
		ChangeToken: aws.String("ChangeToken"),  // Required
		MetricName:  aws.String("MetricName"),   // Required
		Name:        aws.String("ResourceName"), // Required
	}
	resp, err := svc.CreateRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_CreateSizeConstraintSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.CreateSizeConstraintSetInput{
		ChangeToken: aws.String("ChangeToken"),  // Required
		Name:        aws.String("ResourceName"), // Required
	}
	resp, err := svc.CreateSizeConstraintSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_CreateSqlInjectionMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.CreateSqlInjectionMatchSetInput{
		ChangeToken: aws.String("ChangeToken"),  // Required
		Name:        aws.String("ResourceName"), // Required
	}
	resp, err := svc.CreateSqlInjectionMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_CreateWebACL() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.CreateWebACLInput{
		ChangeToken: aws.String("ChangeToken"), // Required
		DefaultAction: &waf.WafAction{ // Required
			Type: aws.String("WafActionType"), // Required
		},
		MetricName: aws.String("MetricName"),   // Required
		Name:       aws.String("ResourceName"), // Required
	}
	resp, err := svc.CreateWebACL(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_CreateXssMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.CreateXssMatchSetInput{
		ChangeToken: aws.String("ChangeToken"),  // Required
		Name:        aws.String("ResourceName"), // Required
	}
	resp, err := svc.CreateXssMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_DeleteByteMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.DeleteByteMatchSetInput{
		ByteMatchSetId: aws.String("ResourceId"),  // Required
		ChangeToken:    aws.String("ChangeToken"), // Required
	}
	resp, err := svc.DeleteByteMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_DeleteIPSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.DeleteIPSetInput{
		ChangeToken: aws.String("ChangeToken"), // Required
		IPSetId:     aws.String("ResourceId"),  // Required
	}
	resp, err := svc.DeleteIPSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_DeleteRule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.DeleteRuleInput{
		ChangeToken: aws.String("ChangeToken"), // Required
		RuleId:      aws.String("ResourceId"),  // Required
	}
	resp, err := svc.DeleteRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_DeleteSizeConstraintSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.DeleteSizeConstraintSetInput{
		ChangeToken:         aws.String("ChangeToken"), // Required
		SizeConstraintSetId: aws.String("ResourceId"),  // Required
	}
	resp, err := svc.DeleteSizeConstraintSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_DeleteSqlInjectionMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.DeleteSqlInjectionMatchSetInput{
		ChangeToken:            aws.String("ChangeToken"), // Required
		SqlInjectionMatchSetId: aws.String("ResourceId"),  // Required
	}
	resp, err := svc.DeleteSqlInjectionMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_DeleteWebACL() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.DeleteWebACLInput{
		ChangeToken: aws.String("ChangeToken"), // Required
		WebACLId:    aws.String("ResourceId"),  // Required
	}
	resp, err := svc.DeleteWebACL(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_DeleteXssMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.DeleteXssMatchSetInput{
		ChangeToken:   aws.String("ChangeToken"), // Required
		XssMatchSetId: aws.String("ResourceId"),  // Required
	}
	resp, err := svc.DeleteXssMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetByteMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetByteMatchSetInput{
		ByteMatchSetId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetByteMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetChangeToken() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	var params *waf.GetChangeTokenInput
	resp, err := svc.GetChangeToken(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetChangeTokenStatus() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetChangeTokenStatusInput{
		ChangeToken: aws.String("ChangeToken"), // Required
	}
	resp, err := svc.GetChangeTokenStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetIPSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetIPSetInput{
		IPSetId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetIPSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetRule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetRuleInput{
		RuleId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetSampledRequests() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetSampledRequestsInput{
		MaxItems: aws.Int64(1),             // Required
		RuleId:   aws.String("ResourceId"), // Required
		TimeWindow: &waf.TimeWindow{ // Required
			EndTime:   aws.Time(time.Now()), // Required
			StartTime: aws.Time(time.Now()), // Required
		},
		WebAclId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetSampledRequests(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetSizeConstraintSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetSizeConstraintSetInput{
		SizeConstraintSetId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetSizeConstraintSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetSqlInjectionMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetSqlInjectionMatchSetInput{
		SqlInjectionMatchSetId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetSqlInjectionMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetWebACL() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetWebACLInput{
		WebACLId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetWebACL(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_GetXssMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.GetXssMatchSetInput{
		XssMatchSetId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.GetXssMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_ListByteMatchSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.ListByteMatchSetsInput{
		Limit:      aws.Int64(1),
		NextMarker: aws.String("NextMarker"),
	}
	resp, err := svc.ListByteMatchSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_ListIPSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.ListIPSetsInput{
		Limit:      aws.Int64(1),
		NextMarker: aws.String("NextMarker"),
	}
	resp, err := svc.ListIPSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_ListRules() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.ListRulesInput{
		Limit:      aws.Int64(1),
		NextMarker: aws.String("NextMarker"),
	}
	resp, err := svc.ListRules(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_ListSizeConstraintSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.ListSizeConstraintSetsInput{
		Limit:      aws.Int64(1),
		NextMarker: aws.String("NextMarker"),
	}
	resp, err := svc.ListSizeConstraintSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_ListSqlInjectionMatchSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.ListSqlInjectionMatchSetsInput{
		Limit:      aws.Int64(1),
		NextMarker: aws.String("NextMarker"),
	}
	resp, err := svc.ListSqlInjectionMatchSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_ListWebACLs() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.ListWebACLsInput{
		Limit:      aws.Int64(1),
		NextMarker: aws.String("NextMarker"),
	}
	resp, err := svc.ListWebACLs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_ListXssMatchSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.ListXssMatchSetsInput{
		Limit:      aws.Int64(1),
		NextMarker: aws.String("NextMarker"),
	}
	resp, err := svc.ListXssMatchSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_UpdateByteMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.UpdateByteMatchSetInput{
		ByteMatchSetId: aws.String("ResourceId"),  // Required
		ChangeToken:    aws.String("ChangeToken"), // Required
		Updates: []*waf.ByteMatchSetUpdate{ // Required
			{ // Required
				Action: aws.String("ChangeAction"), // Required
				ByteMatchTuple: &waf.ByteMatchTuple{ // Required
					FieldToMatch: &waf.FieldToMatch{ // Required
						Type: aws.String("MatchFieldType"), // Required
						Data: aws.String("MatchFieldData"),
					},
					PositionalConstraint: aws.String("PositionalConstraint"), // Required
					TargetString:         []byte("PAYLOAD"),                  // Required
					TextTransformation:   aws.String("TextTransformation"),   // Required
				},
			},
			// More values...
		},
	}
	resp, err := svc.UpdateByteMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_UpdateIPSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.UpdateIPSetInput{
		ChangeToken: aws.String("ChangeToken"), // Required
		IPSetId:     aws.String("ResourceId"),  // Required
		Updates: []*waf.IPSetUpdate{ // Required
			{ // Required
				Action: aws.String("ChangeAction"), // Required
				IPSetDescriptor: &waf.IPSetDescriptor{ // Required
					Type:  aws.String("IPSetDescriptorType"),  // Required
					Value: aws.String("IPSetDescriptorValue"), // Required
				},
			},
			// More values...
		},
	}
	resp, err := svc.UpdateIPSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_UpdateRule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.UpdateRuleInput{
		ChangeToken: aws.String("ChangeToken"), // Required
		RuleId:      aws.String("ResourceId"),  // Required
		Updates: []*waf.RuleUpdate{ // Required
			{ // Required
				Action: aws.String("ChangeAction"), // Required
				Predicate: &waf.Predicate{ // Required
					DataId:  aws.String("ResourceId"),    // Required
					Negated: aws.Bool(true),              // Required
					Type:    aws.String("PredicateType"), // Required
				},
			},
			// More values...
		},
	}
	resp, err := svc.UpdateRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_UpdateSizeConstraintSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.UpdateSizeConstraintSetInput{
		ChangeToken:         aws.String("ChangeToken"), // Required
		SizeConstraintSetId: aws.String("ResourceId"),  // Required
		Updates: []*waf.SizeConstraintSetUpdate{ // Required
			{ // Required
				Action: aws.String("ChangeAction"), // Required
				SizeConstraint: &waf.SizeConstraint{ // Required
					ComparisonOperator: aws.String("ComparisonOperator"), // Required
					FieldToMatch: &waf.FieldToMatch{ // Required
						Type: aws.String("MatchFieldType"), // Required
						Data: aws.String("MatchFieldData"),
					},
					Size:               aws.Int64(1),                     // Required
					TextTransformation: aws.String("TextTransformation"), // Required
				},
			},
			// More values...
		},
	}
	resp, err := svc.UpdateSizeConstraintSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_UpdateSqlInjectionMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.UpdateSqlInjectionMatchSetInput{
		ChangeToken:            aws.String("ChangeToken"), // Required
		SqlInjectionMatchSetId: aws.String("ResourceId"),  // Required
		Updates: []*waf.SqlInjectionMatchSetUpdate{ // Required
			{ // Required
				Action: aws.String("ChangeAction"), // Required
				SqlInjectionMatchTuple: &waf.SqlInjectionMatchTuple{ // Required
					FieldToMatch: &waf.FieldToMatch{ // Required
						Type: aws.String("MatchFieldType"), // Required
						Data: aws.String("MatchFieldData"),
					},
					TextTransformation: aws.String("TextTransformation"), // Required
				},
			},
			// More values...
		},
	}
	resp, err := svc.UpdateSqlInjectionMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_UpdateWebACL() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.UpdateWebACLInput{
		ChangeToken: aws.String("ChangeToken"), // Required
		WebACLId:    aws.String("ResourceId"),  // Required
		DefaultAction: &waf.WafAction{
			Type: aws.String("WafActionType"), // Required
		},
		Updates: []*waf.WebACLUpdate{
			{ // Required
				Action: aws.String("ChangeAction"), // Required
				ActivatedRule: &waf.ActivatedRule{ // Required
					Action: &waf.WafAction{ // Required
						Type: aws.String("WafActionType"), // Required
					},
					Priority: aws.Int64(1),             // Required
					RuleId:   aws.String("ResourceId"), // Required
				},
			},
			// More values...
		},
	}
	resp, err := svc.UpdateWebACL(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleWAF_UpdateXssMatchSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := waf.New(sess)

	params := &waf.UpdateXssMatchSetInput{
		ChangeToken: aws.String("ChangeToken"), // Required
		Updates: []*waf.XssMatchSetUpdate{ // Required
			{ // Required
				Action: aws.String("ChangeAction"), // Required
				XssMatchTuple: &waf.XssMatchTuple{ // Required
					FieldToMatch: &waf.FieldToMatch{ // Required
						Type: aws.String("MatchFieldType"), // Required
						Data: aws.String("MatchFieldData"),
					},
					TextTransformation: aws.String("TextTransformation"), // Required
				},
			},
			// More values...
		},
		XssMatchSetId: aws.String("ResourceId"), // Required
	}
	resp, err := svc.UpdateXssMatchSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
