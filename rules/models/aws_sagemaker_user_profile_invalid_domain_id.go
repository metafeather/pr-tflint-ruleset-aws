// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerUserProfileInvalidDomainIDRule checks the pattern is valid
type AwsSagemakerUserProfileInvalidDomainIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerUserProfileInvalidDomainIDRule returns new rule with default attributes
func NewAwsSagemakerUserProfileInvalidDomainIDRule() *AwsSagemakerUserProfileInvalidDomainIDRule {
	return &AwsSagemakerUserProfileInvalidDomainIDRule{
		resourceType:  "aws_sagemaker_user_profile",
		attributeName: "domain_id",
		max:           63,
		pattern:       regexp.MustCompile(`^d-(-*[a-z0-9]){1,61}`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerUserProfileInvalidDomainIDRule) Name() string {
	return "aws_sagemaker_user_profile_invalid_domain_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerUserProfileInvalidDomainIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerUserProfileInvalidDomainIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerUserProfileInvalidDomainIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerUserProfileInvalidDomainIDRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"domain_id must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^d-(-*[a-z0-9]){1,61}`),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
