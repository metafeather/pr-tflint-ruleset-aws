// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlueDevEndpointInvalidWorkerTypeRule checks the pattern is valid
type AwsGlueDevEndpointInvalidWorkerTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlueDevEndpointInvalidWorkerTypeRule returns new rule with default attributes
func NewAwsGlueDevEndpointInvalidWorkerTypeRule() *AwsGlueDevEndpointInvalidWorkerTypeRule {
	return &AwsGlueDevEndpointInvalidWorkerTypeRule{
		resourceType:  "aws_glue_dev_endpoint",
		attributeName: "worker_type",
		enum: []string{
			"Standard",
			"G.1X",
			"G.2X",
			"G.025X",
			"G.4X",
			"G.8X",
			"Z.2X",
		},
	}
}

// Name returns the rule name
func (r *AwsGlueDevEndpointInvalidWorkerTypeRule) Name() string {
	return "aws_glue_dev_endpoint_invalid_worker_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlueDevEndpointInvalidWorkerTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlueDevEndpointInvalidWorkerTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlueDevEndpointInvalidWorkerTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlueDevEndpointInvalidWorkerTypeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as worker_type`, truncateLongMessage(val)),
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
