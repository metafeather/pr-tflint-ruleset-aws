// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewaySmbFileShareInvalidRoleArnRule checks the pattern is valid
type AwsStoragegatewaySmbFileShareInvalidRoleArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewaySmbFileShareInvalidRoleArnRule returns new rule with default attributes
func NewAwsStoragegatewaySmbFileShareInvalidRoleArnRule() *AwsStoragegatewaySmbFileShareInvalidRoleArnRule {
	return &AwsStoragegatewaySmbFileShareInvalidRoleArnRule{
		resourceType:  "aws_storagegateway_smb_file_share",
		attributeName: "role_arn",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^arn:(aws(|-cn|-us-gov|-iso[A-Za-z0-9_-]*)):iam::([0-9]+):role/(\S+)$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Name() string {
	return "aws_storagegateway_smb_file_share_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Check(runner tflint.Runner) error {
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
					"role_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:(aws(|-cn|-us-gov|-iso[A-Za-z0-9_-]*)):iam::([0-9]+):role/(\S+)$`),
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
