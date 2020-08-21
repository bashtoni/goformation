package sagemaker

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// MonitoringSchedule AWS CloudFormation Resource (AWS::SageMaker::MonitoringSchedule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html
type MonitoringSchedule struct {

	// CreationTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-creationtime
	CreationTime string `json:"CreationTime,omitempty"`

	// EndpointName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-endpointname
	EndpointName string `json:"EndpointName,omitempty"`

	// FailureReason AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-failurereason
	FailureReason string `json:"FailureReason,omitempty"`

	// LastModifiedTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-lastmodifiedtime
	LastModifiedTime string `json:"LastModifiedTime,omitempty"`

	// LastMonitoringExecutionSummary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-lastmonitoringexecutionsummary
	LastMonitoringExecutionSummary *MonitoringSchedule_MonitoringExecutionSummary `json:"LastMonitoringExecutionSummary,omitempty"`

	// MonitoringScheduleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-monitoringschedulearn
	MonitoringScheduleArn string `json:"MonitoringScheduleArn,omitempty"`

	// MonitoringScheduleConfig AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig
	MonitoringScheduleConfig *MonitoringSchedule_MonitoringScheduleConfig `json:"MonitoringScheduleConfig,omitempty"`

	// MonitoringScheduleName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-monitoringschedulename
	MonitoringScheduleName string `json:"MonitoringScheduleName,omitempty"`

	// MonitoringScheduleStatus AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-monitoringschedulestatus
	MonitoringScheduleStatus string `json:"MonitoringScheduleStatus,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

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
func (r *MonitoringSchedule) AWSCloudFormationType() string {
	return "AWS::SageMaker::MonitoringSchedule"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r MonitoringSchedule) MarshalJSON() ([]byte, error) {
	type Properties MonitoringSchedule
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
func (r *MonitoringSchedule) UnmarshalJSON(b []byte) error {
	type P MonitoringSchedule
	props := &MonitoringSchedule{}
	newProps := &struct {
		*P
		CreationTime             types.StringIsh `json:"CreationTime,omitempty"`
		EndpointName             types.StringIsh `json:"EndpointName,omitempty"`
		FailureReason            types.StringIsh `json:"FailureReason,omitempty"`
		LastModifiedTime         types.StringIsh `json:"LastModifiedTime,omitempty"`
		MonitoringScheduleArn    types.StringIsh `json:"MonitoringScheduleArn,omitempty"`
		MonitoringScheduleName   types.StringIsh `json:"MonitoringScheduleName,omitempty"`
		MonitoringScheduleStatus types.StringIsh `json:"MonitoringScheduleStatus,omitempty"`
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

		props.CreationTime = string(newProps.CreationTime)
		props.EndpointName = string(newProps.EndpointName)
		props.FailureReason = string(newProps.FailureReason)
		props.LastModifiedTime = string(newProps.LastModifiedTime)
		props.MonitoringScheduleArn = string(newProps.MonitoringScheduleArn)
		props.MonitoringScheduleName = string(newProps.MonitoringScheduleName)
		props.MonitoringScheduleStatus = string(newProps.MonitoringScheduleStatus)

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
