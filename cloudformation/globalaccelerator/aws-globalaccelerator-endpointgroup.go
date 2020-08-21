package globalaccelerator

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// EndpointGroup AWS CloudFormation Resource (AWS::GlobalAccelerator::EndpointGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html
type EndpointGroup struct {

	// EndpointConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointconfigurations
	EndpointConfigurations []EndpointGroup_EndpointConfiguration `json:"EndpointConfigurations,omitempty"`

	// EndpointGroupRegion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointgroupregion
	EndpointGroupRegion string `json:"EndpointGroupRegion,omitempty"`

	// HealthCheckIntervalSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckintervalseconds
	HealthCheckIntervalSeconds int `json:"HealthCheckIntervalSeconds,omitempty"`

	// HealthCheckPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckpath
	HealthCheckPath string `json:"HealthCheckPath,omitempty"`

	// HealthCheckPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckport
	HealthCheckPort int `json:"HealthCheckPort,omitempty"`

	// HealthCheckProtocol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckprotocol
	HealthCheckProtocol string `json:"HealthCheckProtocol,omitempty"`

	// ListenerArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-listenerarn
	ListenerArn string `json:"ListenerArn,omitempty"`

	// ThresholdCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-thresholdcount
	ThresholdCount int `json:"ThresholdCount,omitempty"`

	// TrafficDialPercentage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-trafficdialpercentage
	TrafficDialPercentage float64 `json:"TrafficDialPercentage,omitempty"`

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
func (r *EndpointGroup) AWSCloudFormationType() string {
	return "AWS::GlobalAccelerator::EndpointGroup"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r EndpointGroup) MarshalJSON() ([]byte, error) {
	type Properties EndpointGroup
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
func (r *EndpointGroup) UnmarshalJSON(b []byte) error {
	type P EndpointGroup
	props := &EndpointGroup{}
	newProps := &struct {
		*P
		EndpointGroupRegion types.StringIsh `json:"EndpointGroupRegion,omitempty"`
		HealthCheckPath     types.StringIsh `json:"HealthCheckPath,omitempty"`
		HealthCheckProtocol types.StringIsh `json:"HealthCheckProtocol,omitempty"`
		ListenerArn         types.StringIsh `json:"ListenerArn,omitempty"`
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

		props.EndpointGroupRegion = string(newProps.EndpointGroupRegion)
		props.HealthCheckPath = string(newProps.HealthCheckPath)
		props.HealthCheckProtocol = string(newProps.HealthCheckProtocol)
		props.ListenerArn = string(newProps.ListenerArn)

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
