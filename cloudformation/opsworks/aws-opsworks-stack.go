package opsworks

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// Stack AWS CloudFormation Resource (AWS::OpsWorks::Stack)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html
type Stack struct {

	// AgentVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-agentversion
	AgentVersion string `json:"AgentVersion,omitempty"`

	// Attributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-attributes
	Attributes map[string]string `json:"Attributes,omitempty"`

	// ChefConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-chefconfiguration
	ChefConfiguration *Stack_ChefConfiguration `json:"ChefConfiguration,omitempty"`

	// CloneAppIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-cloneappids
	CloneAppIds []string `json:"CloneAppIds,omitempty"`

	// ClonePermissions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-clonepermissions
	ClonePermissions bool `json:"ClonePermissions,omitempty"`

	// ConfigurationManager AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-configmanager
	ConfigurationManager *Stack_StackConfigurationManager `json:"ConfigurationManager,omitempty"`

	// CustomCookbooksSource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-custcookbooksource
	CustomCookbooksSource *Stack_Source `json:"CustomCookbooksSource,omitempty"`

	// CustomJson AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-custjson
	CustomJson interface{} `json:"CustomJson,omitempty"`

	// DefaultAvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultaz
	DefaultAvailabilityZone string `json:"DefaultAvailabilityZone,omitempty"`

	// DefaultInstanceProfileArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultinstanceprof
	DefaultInstanceProfileArn string `json:"DefaultInstanceProfileArn,omitempty"`

	// DefaultOs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultos
	DefaultOs string `json:"DefaultOs,omitempty"`

	// DefaultRootDeviceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultrootdevicetype
	DefaultRootDeviceType string `json:"DefaultRootDeviceType,omitempty"`

	// DefaultSshKeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultsshkeyname
	DefaultSshKeyName string `json:"DefaultSshKeyName,omitempty"`

	// DefaultSubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#defaultsubnet
	DefaultSubnetId string `json:"DefaultSubnetId,omitempty"`

	// EcsClusterArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-ecsclusterarn
	EcsClusterArn string `json:"EcsClusterArn,omitempty"`

	// ElasticIps AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-elasticips
	ElasticIps []Stack_ElasticIp `json:"ElasticIps,omitempty"`

	// HostnameTheme AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-hostnametheme
	HostnameTheme string `json:"HostnameTheme,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-name
	Name string `json:"Name,omitempty"`

	// RdsDbInstances AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-rdsdbinstances
	RdsDbInstances []Stack_RdsDbInstance `json:"RdsDbInstances,omitempty"`

	// ServiceRoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-servicerolearn
	ServiceRoleArn string `json:"ServiceRoleArn,omitempty"`

	// SourceStackId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-sourcestackid
	SourceStackId string `json:"SourceStackId,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

	// UseCustomCookbooks AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#usecustcookbooks
	UseCustomCookbooks bool `json:"UseCustomCookbooks,omitempty"`

	// UseOpsworksSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-useopsworkssecuritygroups
	UseOpsworksSecurityGroups bool `json:"UseOpsworksSecurityGroups,omitempty"`

	// VpcId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-vpcid
	VpcId string `json:"VpcId,omitempty"`

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
func (r *Stack) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Stack"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r Stack) MarshalJSON() ([]byte, error) {
	type Properties Stack
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
func (r *Stack) UnmarshalJSON(b []byte) error {
	type P Stack
	props := &Stack{}
	newProps := &struct {
		*P
		AgentVersion              types.StringIsh `json:"AgentVersion,omitempty"`
		DefaultAvailabilityZone   types.StringIsh `json:"DefaultAvailabilityZone,omitempty"`
		DefaultInstanceProfileArn types.StringIsh `json:"DefaultInstanceProfileArn,omitempty"`
		DefaultOs                 types.StringIsh `json:"DefaultOs,omitempty"`
		DefaultRootDeviceType     types.StringIsh `json:"DefaultRootDeviceType,omitempty"`
		DefaultSshKeyName         types.StringIsh `json:"DefaultSshKeyName,omitempty"`
		DefaultSubnetId           types.StringIsh `json:"DefaultSubnetId,omitempty"`
		EcsClusterArn             types.StringIsh `json:"EcsClusterArn,omitempty"`
		HostnameTheme             types.StringIsh `json:"HostnameTheme,omitempty"`
		Name                      types.StringIsh `json:"Name,omitempty"`
		ServiceRoleArn            types.StringIsh `json:"ServiceRoleArn,omitempty"`
		SourceStackId             types.StringIsh `json:"SourceStackId,omitempty"`
		VpcId                     types.StringIsh `json:"VpcId,omitempty"`
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
		props.DefaultAvailabilityZone = string(newProps.DefaultAvailabilityZone)
		props.DefaultInstanceProfileArn = string(newProps.DefaultInstanceProfileArn)
		props.DefaultOs = string(newProps.DefaultOs)
		props.DefaultRootDeviceType = string(newProps.DefaultRootDeviceType)
		props.DefaultSshKeyName = string(newProps.DefaultSshKeyName)
		props.DefaultSubnetId = string(newProps.DefaultSubnetId)
		props.EcsClusterArn = string(newProps.EcsClusterArn)
		props.HostnameTheme = string(newProps.HostnameTheme)
		props.Name = string(newProps.Name)
		props.ServiceRoleArn = string(newProps.ServiceRoleArn)
		props.SourceStackId = string(newProps.SourceStackId)
		props.VpcId = string(newProps.VpcId)

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
