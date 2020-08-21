package opsworks

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// Instance AWS CloudFormation Resource (AWS::OpsWorks::Instance)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html
type Instance struct {

	// AgentVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-agentversion
	AgentVersion string `json:"AgentVersion,omitempty"`

	// AmiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-amiid
	AmiId string `json:"AmiId,omitempty"`

	// Architecture AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-architecture
	Architecture string `json:"Architecture,omitempty"`

	// AutoScalingType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-autoscalingtype
	AutoScalingType string `json:"AutoScalingType,omitempty"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-availabilityzone
	AvailabilityZone string `json:"AvailabilityZone,omitempty"`

	// BlockDeviceMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-blockdevicemappings
	BlockDeviceMappings []Instance_BlockDeviceMapping `json:"BlockDeviceMappings,omitempty"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized,omitempty"`

	// ElasticIps AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-elasticips
	ElasticIps []string `json:"ElasticIps,omitempty"`

	// Hostname AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-hostname
	Hostname string `json:"Hostname,omitempty"`

	// InstallUpdatesOnBoot AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-installupdatesonboot
	InstallUpdatesOnBoot bool `json:"InstallUpdatesOnBoot,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-instancetype
	InstanceType string `json:"InstanceType,omitempty"`

	// LayerIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-layerids
	LayerIds []string `json:"LayerIds,omitempty"`

	// Os AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-os
	Os string `json:"Os,omitempty"`

	// RootDeviceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-rootdevicetype
	RootDeviceType string `json:"RootDeviceType,omitempty"`

	// SshKeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-sshkeyname
	SshKeyName string `json:"SshKeyName,omitempty"`

	// StackId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-stackid
	StackId string `json:"StackId,omitempty"`

	// SubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-subnetid
	SubnetId string `json:"SubnetId,omitempty"`

	// Tenancy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-tenancy
	Tenancy string `json:"Tenancy,omitempty"`

	// TimeBasedAutoScaling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-timebasedautoscaling
	TimeBasedAutoScaling *Instance_TimeBasedAutoScaling `json:"TimeBasedAutoScaling,omitempty"`

	// VirtualizationType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-virtualizationtype
	VirtualizationType string `json:"VirtualizationType,omitempty"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-volumes
	Volumes []string `json:"Volumes,omitempty"`

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
func (r *Instance) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Instance"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r Instance) MarshalJSON() ([]byte, error) {
	type Properties Instance
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
func (r *Instance) UnmarshalJSON(b []byte) error {
	type P Instance
	props := &Instance{}
	newProps := &struct {
		*P
		AgentVersion       types.StringIsh `json:"AgentVersion,omitempty"`
		AmiId              types.StringIsh `json:"AmiId,omitempty"`
		Architecture       types.StringIsh `json:"Architecture,omitempty"`
		AutoScalingType    types.StringIsh `json:"AutoScalingType,omitempty"`
		AvailabilityZone   types.StringIsh `json:"AvailabilityZone,omitempty"`
		Hostname           types.StringIsh `json:"Hostname,omitempty"`
		InstanceType       types.StringIsh `json:"InstanceType,omitempty"`
		Os                 types.StringIsh `json:"Os,omitempty"`
		RootDeviceType     types.StringIsh `json:"RootDeviceType,omitempty"`
		SshKeyName         types.StringIsh `json:"SshKeyName,omitempty"`
		StackId            types.StringIsh `json:"StackId,omitempty"`
		SubnetId           types.StringIsh `json:"SubnetId,omitempty"`
		Tenancy            types.StringIsh `json:"Tenancy,omitempty"`
		VirtualizationType types.StringIsh `json:"VirtualizationType,omitempty"`
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

		props.AgentVersion = string(newProps.AgentVersion)
		props.AmiId = string(newProps.AmiId)
		props.Architecture = string(newProps.Architecture)
		props.AutoScalingType = string(newProps.AutoScalingType)
		props.AvailabilityZone = string(newProps.AvailabilityZone)
		props.Hostname = string(newProps.Hostname)
		props.InstanceType = string(newProps.InstanceType)
		props.Os = string(newProps.Os)
		props.RootDeviceType = string(newProps.RootDeviceType)
		props.SshKeyName = string(newProps.SshKeyName)
		props.StackId = string(newProps.StackId)
		props.SubnetId = string(newProps.SubnetId)
		props.Tenancy = string(newProps.Tenancy)
		props.VirtualizationType = string(newProps.VirtualizationType)

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
