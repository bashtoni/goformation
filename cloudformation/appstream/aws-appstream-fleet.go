package appstream

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// Fleet AWS CloudFormation Resource (AWS::AppStream::Fleet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html
type Fleet struct {

	// ComputeCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-computecapacity
	ComputeCapacity *Fleet_ComputeCapacity `json:"ComputeCapacity,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-description
	Description string `json:"Description,omitempty"`

	// DisconnectTimeoutInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-disconnecttimeoutinseconds
	DisconnectTimeoutInSeconds int `json:"DisconnectTimeoutInSeconds,omitempty"`

	// DisplayName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-displayname
	DisplayName string `json:"DisplayName,omitempty"`

	// DomainJoinInfo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-domainjoininfo
	DomainJoinInfo *Fleet_DomainJoinInfo `json:"DomainJoinInfo,omitempty"`

	// EnableDefaultInternetAccess AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-enabledefaultinternetaccess
	EnableDefaultInternetAccess bool `json:"EnableDefaultInternetAccess,omitempty"`

	// FleetType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-fleettype
	FleetType string `json:"FleetType,omitempty"`

	// IdleDisconnectTimeoutInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-idledisconnecttimeoutinseconds
	IdleDisconnectTimeoutInSeconds int `json:"IdleDisconnectTimeoutInSeconds,omitempty"`

	// ImageArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-imagearn
	ImageArn string `json:"ImageArn,omitempty"`

	// ImageName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-imagename
	ImageName string `json:"ImageName,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-instancetype
	InstanceType string `json:"InstanceType,omitempty"`

	// MaxUserDurationInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-maxuserdurationinseconds
	MaxUserDurationInSeconds int `json:"MaxUserDurationInSeconds,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-name
	Name string `json:"Name,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

	// VpcConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-vpcconfig
	VpcConfig *Fleet_VpcConfig `json:"VpcConfig,omitempty"`

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
func (r *Fleet) AWSCloudFormationType() string {
	return "AWS::AppStream::Fleet"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r Fleet) MarshalJSON() ([]byte, error) {
	type Properties Fleet
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
func (r *Fleet) UnmarshalJSON(b []byte) error {
	type P Fleet
	props := &Fleet{}
	newProps := &struct {
		*P
		Description  types.StringIsh `json:"Description,omitempty"`
		DisplayName  types.StringIsh `json:"DisplayName,omitempty"`
		FleetType    types.StringIsh `json:"FleetType,omitempty"`
		ImageArn     types.StringIsh `json:"ImageArn,omitempty"`
		ImageName    types.StringIsh `json:"ImageName,omitempty"`
		InstanceType types.StringIsh `json:"InstanceType,omitempty"`
		Name         types.StringIsh `json:"Name,omitempty"`
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

		props.Description = string(newProps.Description)
		props.DisplayName = string(newProps.DisplayName)
		props.FleetType = string(newProps.FleetType)
		props.ImageArn = string(newProps.ImageArn)
		props.ImageName = string(newProps.ImageName)
		props.InstanceType = string(newProps.InstanceType)
		props.Name = string(newProps.Name)

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
