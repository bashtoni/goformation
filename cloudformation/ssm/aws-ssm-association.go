package ssm

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// Association AWS CloudFormation Resource (AWS::SSM::Association)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html
type Association struct {

	// ApplyOnlyAtCronInterval AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-applyonlyatcroninterval
	ApplyOnlyAtCronInterval bool `json:"ApplyOnlyAtCronInterval,omitempty"`

	// AssociationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-associationname
	AssociationName string `json:"AssociationName,omitempty"`

	// AutomationTargetParameterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-automationtargetparametername
	AutomationTargetParameterName string `json:"AutomationTargetParameterName,omitempty"`

	// ComplianceSeverity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-complianceseverity
	ComplianceSeverity string `json:"ComplianceSeverity,omitempty"`

	// DocumentVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
	DocumentVersion string `json:"DocumentVersion,omitempty"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
	InstanceId string `json:"InstanceId,omitempty"`

	// MaxConcurrency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxconcurrency
	MaxConcurrency string `json:"MaxConcurrency,omitempty"`

	// MaxErrors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxerrors
	MaxErrors string `json:"MaxErrors,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
	Name string `json:"Name,omitempty"`

	// OutputLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-outputlocation
	OutputLocation *Association_InstanceAssociationOutputLocation `json:"OutputLocation,omitempty"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
	Parameters map[string]interface{} `json:"Parameters,omitempty"`

	// ScheduleExpression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
	ScheduleExpression string `json:"ScheduleExpression,omitempty"`

	// SyncCompliance AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-synccompliance
	SyncCompliance string `json:"SyncCompliance,omitempty"`

	// Targets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
	Targets []Association_Target `json:"Targets,omitempty"`

	// WaitForSuccessTimeoutSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-waitforsuccesstimeoutseconds
	WaitForSuccessTimeoutSeconds int `json:"WaitForSuccessTimeoutSeconds,omitempty"`

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
func (r *Association) AWSCloudFormationType() string {
	return "AWS::SSM::Association"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r Association) MarshalJSON() ([]byte, error) {
	type Properties Association
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
func (r *Association) UnmarshalJSON(b []byte) error {
	type P Association
	props := &Association{}
	newProps := &struct {
		*P
		AssociationName               types.StringIsh `json:"AssociationName,omitempty"`
		AutomationTargetParameterName types.StringIsh `json:"AutomationTargetParameterName,omitempty"`
		ComplianceSeverity            types.StringIsh `json:"ComplianceSeverity,omitempty"`
		DocumentVersion               types.StringIsh `json:"DocumentVersion,omitempty"`
		InstanceId                    types.StringIsh `json:"InstanceId,omitempty"`
		MaxConcurrency                types.StringIsh `json:"MaxConcurrency,omitempty"`
		MaxErrors                     types.StringIsh `json:"MaxErrors,omitempty"`
		Name                          types.StringIsh `json:"Name,omitempty"`
		ScheduleExpression            types.StringIsh `json:"ScheduleExpression,omitempty"`
		SyncCompliance                types.StringIsh `json:"SyncCompliance,omitempty"`
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

		props.AssociationName = string(newProps.AssociationName)
		props.AutomationTargetParameterName = string(newProps.AutomationTargetParameterName)
		props.ComplianceSeverity = string(newProps.ComplianceSeverity)
		props.DocumentVersion = string(newProps.DocumentVersion)
		props.InstanceId = string(newProps.InstanceId)
		props.MaxConcurrency = string(newProps.MaxConcurrency)
		props.MaxErrors = string(newProps.MaxErrors)
		props.Name = string(newProps.Name)
		props.ScheduleExpression = string(newProps.ScheduleExpression)
		props.SyncCompliance = string(newProps.SyncCompliance)

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
