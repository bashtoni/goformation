package config

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// ConformancePack AWS CloudFormation Resource (AWS::Config::ConformancePack)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html
type ConformancePack struct {

	// ConformancePackInputParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-conformancepackinputparameters
	ConformancePackInputParameters []ConformancePack_ConformancePackInputParameter `json:"ConformancePackInputParameters,omitempty"`

	// ConformancePackName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-conformancepackname
	ConformancePackName string `json:"ConformancePackName,omitempty"`

	// DeliveryS3Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-deliverys3bucket
	DeliveryS3Bucket string `json:"DeliveryS3Bucket,omitempty"`

	// DeliveryS3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-deliverys3keyprefix
	DeliveryS3KeyPrefix string `json:"DeliveryS3KeyPrefix,omitempty"`

	// TemplateBody AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-templatebody
	TemplateBody string `json:"TemplateBody,omitempty"`

	// TemplateS3Uri AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-templates3uri
	TemplateS3Uri string `json:"TemplateS3Uri,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *ConformancePack) AWSCloudFormationType() string {
	return "AWS::Config::ConformancePack"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r ConformancePack) MarshalJSON() ([]byte, error) {
	type Properties ConformancePack
	return json.Marshal(&struct {
		Type                string
		Properties          Properties
		DependsOn           []string                     `json:"DependsOn,omitempty"`
		Metadata            map[string]interface{}       `json:"Metadata,omitempty"`
		DeletionPolicy      policies.DeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy policies.UpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Condition           string                       `json:"Condition,omitempty"`
	}{
		Type:                r.AWSCloudFormationType(),
		Properties:          (Properties)(r),
		DependsOn:           r.AWSCloudFormationDependsOn,
		Metadata:            r.AWSCloudFormationMetadata,
		DeletionPolicy:      r.AWSCloudFormationDeletionPolicy,
		UpdateReplacePolicy: r.AWSCloudFormationUpdateReplacePolicy,
		Condition:           r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *ConformancePack) UnmarshalJSON(b []byte) error {
	type P ConformancePack
	props := &ConformancePack{}
	newProps := &struct {
		*P
		ConformancePackName types.StringIsh `json:"ConformancePackName,omitempty"`
		DeliveryS3Bucket    types.StringIsh `json:"DeliveryS3Bucket,omitempty"`
		DeliveryS3KeyPrefix types.StringIsh `json:"DeliveryS3KeyPrefix,omitempty"`
		TemplateBody        types.StringIsh `json:"TemplateBody,omitempty"`
		TemplateS3Uri       types.StringIsh `json:"TemplateS3Uri,omitempty"`
	}{P: (*P)(props)}
	res := &struct {
		Type                string
		Properties          json.RawMessage
		DependsOn           interface{}
		Metadata            map[string]interface{}
		DeletionPolicy      string
		UpdateReplacePolicy string
		Condition           string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	// Unmarshal everything except the properties
	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	if res.Properties != nil {
		// Unmarshal the properties, being forgiving of type mismatches
		if err := json.Unmarshal(res.Properties, newProps); err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return err
		}

		props.ConformancePackName = string(newProps.ConformancePackName)
		props.DeliveryS3Bucket = string(newProps.DeliveryS3Bucket)
		props.DeliveryS3KeyPrefix = string(newProps.DeliveryS3KeyPrefix)
		props.TemplateBody = string(newProps.TemplateBody)
		props.TemplateS3Uri = string(newProps.TemplateS3Uri)

		*r = *props
	}
	if dependsOn, ok := res.DependsOn.(string); ok {
		r.AWSCloudFormationDependsOn = []string{dependsOn}
	}
	if dependsOn, ok := res.DependsOn.([]interface{}); ok {
		var do []string
		for _, d := range dependsOn {
			if dStr, ok := d.(string); ok {
				do = append(do, dStr)
			}
		}
		r.AWSCloudFormationDependsOn = do
	}

	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.UpdateReplacePolicy != "" {
		r.AWSCloudFormationUpdateReplacePolicy = policies.UpdateReplacePolicy(res.UpdateReplacePolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}
