package ecs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// TaskDefinition AWS CloudFormation Resource (AWS::ECS::TaskDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
type TaskDefinition struct {

	// ContainerDefinitions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-containerdefinitions
	ContainerDefinitions []TaskDefinition_ContainerDefinition `json:"ContainerDefinitions,omitempty"`

	// Cpu AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-cpu
	Cpu string `json:"Cpu,omitempty"`

	// ExecutionRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-executionrolearn
	ExecutionRoleArn string `json:"ExecutionRoleArn,omitempty"`

	// Family AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-family
	Family string `json:"Family,omitempty"`

	// InferenceAccelerators AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-inferenceaccelerators
	InferenceAccelerators []TaskDefinition_InferenceAccelerator `json:"InferenceAccelerators,omitempty"`

	// IpcMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-ipcmode
	IpcMode string `json:"IpcMode,omitempty"`

	// Memory AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-memory
	Memory string `json:"Memory,omitempty"`

	// NetworkMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-networkmode
	NetworkMode string `json:"NetworkMode,omitempty"`

	// PidMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-pidmode
	PidMode string `json:"PidMode,omitempty"`

	// PlacementConstraints AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-placementconstraints
	PlacementConstraints []TaskDefinition_TaskDefinitionPlacementConstraint `json:"PlacementConstraints,omitempty"`

	// ProxyConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-proxyconfiguration
	ProxyConfiguration *TaskDefinition_ProxyConfiguration `json:"ProxyConfiguration,omitempty"`

	// RequiresCompatibilities AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-requirescompatibilities
	RequiresCompatibilities []string `json:"RequiresCompatibilities,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

	// TaskDefinitionStatus AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-taskdefinitionstatus
	TaskDefinitionStatus string `json:"TaskDefinitionStatus,omitempty"`

	// TaskRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-taskrolearn
	TaskRoleArn string `json:"TaskRoleArn,omitempty"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-volumes
	Volumes []TaskDefinition_Volume `json:"Volumes,omitempty"`

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
func (r *TaskDefinition) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r TaskDefinition) MarshalJSON() ([]byte, error) {
	type Properties TaskDefinition
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
func (r *TaskDefinition) UnmarshalJSON(b []byte) error {
	type P TaskDefinition
	props := &TaskDefinition{}
	newProps := &struct {
		*P
		Cpu                  types.StringIsh `json:"Cpu,omitempty"`
		ExecutionRoleArn     types.StringIsh `json:"ExecutionRoleArn,omitempty"`
		Family               types.StringIsh `json:"Family,omitempty"`
		IpcMode              types.StringIsh `json:"IpcMode,omitempty"`
		Memory               types.StringIsh `json:"Memory,omitempty"`
		NetworkMode          types.StringIsh `json:"NetworkMode,omitempty"`
		PidMode              types.StringIsh `json:"PidMode,omitempty"`
		TaskDefinitionStatus types.StringIsh `json:"TaskDefinitionStatus,omitempty"`
		TaskRoleArn          types.StringIsh `json:"TaskRoleArn,omitempty"`
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

		props.Cpu = string(newProps.Cpu)
		props.ExecutionRoleArn = string(newProps.ExecutionRoleArn)
		props.Family = string(newProps.Family)
		props.IpcMode = string(newProps.IpcMode)
		props.Memory = string(newProps.Memory)
		props.NetworkMode = string(newProps.NetworkMode)
		props.PidMode = string(newProps.PidMode)
		props.TaskDefinitionStatus = string(newProps.TaskDefinitionStatus)
		props.TaskRoleArn = string(newProps.TaskRoleArn)

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
