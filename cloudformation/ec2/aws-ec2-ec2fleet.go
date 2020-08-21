package ec2

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// EC2Fleet AWS CloudFormation Resource (AWS::EC2::EC2Fleet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html
type EC2Fleet struct {

	// ExcessCapacityTerminationPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-excesscapacityterminationpolicy
	ExcessCapacityTerminationPolicy string `json:"ExcessCapacityTerminationPolicy,omitempty"`

	// LaunchTemplateConfigs AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-launchtemplateconfigs
	LaunchTemplateConfigs []EC2Fleet_FleetLaunchTemplateConfigRequest `json:"LaunchTemplateConfigs,omitempty"`

	// OnDemandOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-ondemandoptions
	OnDemandOptions *EC2Fleet_OnDemandOptionsRequest `json:"OnDemandOptions,omitempty"`

	// ReplaceUnhealthyInstances AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-replaceunhealthyinstances
	ReplaceUnhealthyInstances bool `json:"ReplaceUnhealthyInstances,omitempty"`

	// SpotOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-spotoptions
	SpotOptions *EC2Fleet_SpotOptionsRequest `json:"SpotOptions,omitempty"`

	// TagSpecifications AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-tagspecifications
	TagSpecifications []EC2Fleet_TagSpecification `json:"TagSpecifications,omitempty"`

	// TargetCapacitySpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-targetcapacityspecification
	TargetCapacitySpecification *EC2Fleet_TargetCapacitySpecificationRequest `json:"TargetCapacitySpecification,omitempty"`

	// TerminateInstancesWithExpiration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-terminateinstanceswithexpiration
	TerminateInstancesWithExpiration bool `json:"TerminateInstancesWithExpiration,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-type
	Type string `json:"Type,omitempty"`

	// ValidFrom AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validfrom
	ValidFrom string `json:"ValidFrom,omitempty"`

	// ValidUntil AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validuntil
	ValidUntil string `json:"ValidUntil,omitempty"`

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
func (r *EC2Fleet) AWSCloudFormationType() string {
	return "AWS::EC2::EC2Fleet"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r EC2Fleet) MarshalJSON() ([]byte, error) {
	type Properties EC2Fleet
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
func (r *EC2Fleet) UnmarshalJSON(b []byte) error {
	type P EC2Fleet
	props := &EC2Fleet{}
	newProps := &struct {
		*P
		ExcessCapacityTerminationPolicy types.StringIsh `json:"ExcessCapacityTerminationPolicy,omitempty"`
		Type                            types.StringIsh `json:"Type,omitempty"`
		ValidFrom                       types.StringIsh `json:"ValidFrom,omitempty"`
		ValidUntil                      types.StringIsh `json:"ValidUntil,omitempty"`
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

		props.ExcessCapacityTerminationPolicy = string(newProps.ExcessCapacityTerminationPolicy)
		props.Type = string(newProps.Type)
		props.ValidFrom = string(newProps.ValidFrom)
		props.ValidUntil = string(newProps.ValidUntil)

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
