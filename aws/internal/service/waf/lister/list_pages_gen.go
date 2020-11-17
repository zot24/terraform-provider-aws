// Code generated by "aws/internal/generators/listpages/main.go -function=ListByteMatchSets,ListGeoMatchSets,ListIPSets,ListRateBasedRules,ListRegexMatchSets,ListRegexPatternSets,ListRuleGroups,ListRules,ListSizeConstraintSets,ListSqlInjectionMatchSets,ListWebACLs,ListXssMatchSets -paginator=NextMarker github.com/aws/aws-sdk-go/service/waf"; DO NOT EDIT.

package lister

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/waf"
)

func ListByteMatchSetsPages(conn *waf.WAF, input *waf.ListByteMatchSetsInput, fn func(*waf.ListByteMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListByteMatchSets(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListGeoMatchSetsPages(conn *waf.WAF, input *waf.ListGeoMatchSetsInput, fn func(*waf.ListGeoMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListGeoMatchSets(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListIPSetsPages(conn *waf.WAF, input *waf.ListIPSetsInput, fn func(*waf.ListIPSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListIPSets(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListRateBasedRulesPages(conn *waf.WAF, input *waf.ListRateBasedRulesInput, fn func(*waf.ListRateBasedRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRateBasedRules(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListRegexMatchSetsPages(conn *waf.WAF, input *waf.ListRegexMatchSetsInput, fn func(*waf.ListRegexMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexMatchSets(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListRegexPatternSetsPages(conn *waf.WAF, input *waf.ListRegexPatternSetsInput, fn func(*waf.ListRegexPatternSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListRegexPatternSets(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListRuleGroupsPages(conn *waf.WAF, input *waf.ListRuleGroupsInput, fn func(*waf.ListRuleGroupsOutput, bool) bool) error {
	for {
		output, err := conn.ListRuleGroups(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListRulesPages(conn *waf.WAF, input *waf.ListRulesInput, fn func(*waf.ListRulesOutput, bool) bool) error {
	for {
		output, err := conn.ListRules(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListSizeConstraintSetsPages(conn *waf.WAF, input *waf.ListSizeConstraintSetsInput, fn func(*waf.ListSizeConstraintSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListSizeConstraintSets(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListSqlInjectionMatchSetsPages(conn *waf.WAF, input *waf.ListSqlInjectionMatchSetsInput, fn func(*waf.ListSqlInjectionMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListSqlInjectionMatchSets(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListWebACLsPages(conn *waf.WAF, input *waf.ListWebACLsInput, fn func(*waf.ListWebACLsOutput, bool) bool) error {
	for {
		output, err := conn.ListWebACLs(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}

func ListXssMatchSetsPages(conn *waf.WAF, input *waf.ListXssMatchSetsInput, fn func(*waf.ListXssMatchSetsOutput, bool) bool) error {
	for {
		output, err := conn.ListXssMatchSets(input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextMarker = output.NextMarker
	}
	return nil
}