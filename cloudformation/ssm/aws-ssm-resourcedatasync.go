package ssm

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// ResourceDataSync AWS CloudFormation Resource (AWS::SSM::ResourceDataSync)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html
type ResourceDataSync struct {

	// BucketName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketname
	BucketName string `json:"BucketName,omitempty"`

	// BucketPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketprefix
	BucketPrefix string `json:"BucketPrefix,omitempty"`

	// BucketRegion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketregion
	BucketRegion string `json:"BucketRegion,omitempty"`

	// KMSKeyArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-kmskeyarn
	KMSKeyArn string `json:"KMSKeyArn,omitempty"`

	// S3Destination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-s3destination
	S3Destination *ResourceDataSync_S3Destination `json:"S3Destination,omitempty"`

	// SyncFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncformat
	SyncFormat string `json:"SyncFormat,omitempty"`

	// SyncName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncname
	SyncName string `json:"SyncName,omitempty"`

	// SyncSource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncsource
	SyncSource *ResourceDataSync_SyncSource `json:"SyncSource,omitempty"`

	// SyncType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-synctype
	SyncType string `json:"SyncType,omitempty"`

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
func (r *ResourceDataSync) AWSCloudFormationType() string {
	return "AWS::SSM::ResourceDataSync"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r ResourceDataSync) MarshalJSON() ([]byte, error) {
	type Properties ResourceDataSync
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
func (r *ResourceDataSync) UnmarshalJSON(b []byte) error {
	type P ResourceDataSync
	props := &ResourceDataSync{}
	newProps := &struct {
		*P
		BucketName   types.StringIsh `json:"BucketName,omitempty"`
		BucketPrefix types.StringIsh `json:"BucketPrefix,omitempty"`
		BucketRegion types.StringIsh `json:"BucketRegion,omitempty"`
		KMSKeyArn    types.StringIsh `json:"KMSKeyArn,omitempty"`
		SyncFormat   types.StringIsh `json:"SyncFormat,omitempty"`
		SyncName     types.StringIsh `json:"SyncName,omitempty"`
		SyncType     types.StringIsh `json:"SyncType,omitempty"`
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

		props.BucketName = string(newProps.BucketName)
		props.BucketPrefix = string(newProps.BucketPrefix)
		props.BucketRegion = string(newProps.BucketRegion)
		props.KMSKeyArn = string(newProps.KMSKeyArn)
		props.SyncFormat = string(newProps.SyncFormat)
		props.SyncName = string(newProps.SyncName)
		props.SyncType = string(newProps.SyncType)

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
