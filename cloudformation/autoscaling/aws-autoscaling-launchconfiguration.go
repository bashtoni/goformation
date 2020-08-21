package autoscaling

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// LaunchConfiguration AWS CloudFormation Resource (AWS::AutoScaling::LaunchConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html
type LaunchConfiguration struct {

	// AssociatePublicIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cf-as-launchconfig-associatepubip
	AssociatePublicIpAddress bool `json:"AssociatePublicIpAddress,omitempty"`

	// BlockDeviceMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-blockdevicemappings
	BlockDeviceMappings []LaunchConfiguration_BlockDeviceMapping `json:"BlockDeviceMappings,omitempty"`

	// ClassicLinkVPCId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcid
	ClassicLinkVPCId string `json:"ClassicLinkVPCId,omitempty"`

	// ClassicLinkVPCSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcsecuritygroups
	ClassicLinkVPCSecurityGroups []string `json:"ClassicLinkVPCSecurityGroups,omitempty"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized,omitempty"`

	// IamInstanceProfile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-iaminstanceprofile
	IamInstanceProfile string `json:"IamInstanceProfile,omitempty"`

	// ImageId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-imageid
	ImageId string `json:"ImageId,omitempty"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instanceid
	InstanceId string `json:"InstanceId,omitempty"`

	// InstanceMonitoring AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancemonitoring
	InstanceMonitoring bool `json:"InstanceMonitoring,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancetype
	InstanceType string `json:"InstanceType,omitempty"`

	// KernelId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-kernelid
	KernelId string `json:"KernelId,omitempty"`

	// KeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-keyname
	KeyName string `json:"KeyName,omitempty"`

	// LaunchConfigurationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-autoscaling-launchconfig-launchconfigurationname
	LaunchConfigurationName string `json:"LaunchConfigurationName,omitempty"`

	// PlacementTenancy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-placementtenancy
	PlacementTenancy string `json:"PlacementTenancy,omitempty"`

	// RamDiskId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ramdiskid
	RamDiskId string `json:"RamDiskId,omitempty"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-securitygroups
	SecurityGroups []string `json:"SecurityGroups,omitempty"`

	// SpotPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-spotprice
	SpotPrice string `json:"SpotPrice,omitempty"`

	// UserData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-userdata
	UserData string `json:"UserData,omitempty"`

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
func (r *LaunchConfiguration) AWSCloudFormationType() string {
	return "AWS::AutoScaling::LaunchConfiguration"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r LaunchConfiguration) MarshalJSON() ([]byte, error) {
	type Properties LaunchConfiguration
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
func (r *LaunchConfiguration) UnmarshalJSON(b []byte) error {
	type P LaunchConfiguration
	props := &LaunchConfiguration{}
	newProps := &struct {
		*P
		ClassicLinkVPCId        types.StringIsh `json:"ClassicLinkVPCId,omitempty"`
		IamInstanceProfile      types.StringIsh `json:"IamInstanceProfile,omitempty"`
		ImageId                 types.StringIsh `json:"ImageId,omitempty"`
		InstanceId              types.StringIsh `json:"InstanceId,omitempty"`
		InstanceType            types.StringIsh `json:"InstanceType,omitempty"`
		KernelId                types.StringIsh `json:"KernelId,omitempty"`
		KeyName                 types.StringIsh `json:"KeyName,omitempty"`
		LaunchConfigurationName types.StringIsh `json:"LaunchConfigurationName,omitempty"`
		PlacementTenancy        types.StringIsh `json:"PlacementTenancy,omitempty"`
		RamDiskId               types.StringIsh `json:"RamDiskId,omitempty"`
		SpotPrice               types.StringIsh `json:"SpotPrice,omitempty"`
		UserData                types.StringIsh `json:"UserData,omitempty"`
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

		props.ClassicLinkVPCId = string(newProps.ClassicLinkVPCId)
		props.IamInstanceProfile = string(newProps.IamInstanceProfile)
		props.ImageId = string(newProps.ImageId)
		props.InstanceId = string(newProps.InstanceId)
		props.InstanceType = string(newProps.InstanceType)
		props.KernelId = string(newProps.KernelId)
		props.KeyName = string(newProps.KeyName)
		props.LaunchConfigurationName = string(newProps.LaunchConfigurationName)
		props.PlacementTenancy = string(newProps.PlacementTenancy)
		props.RamDiskId = string(newProps.RamDiskId)
		props.SpotPrice = string(newProps.SpotPrice)
		props.UserData = string(newProps.UserData)

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
