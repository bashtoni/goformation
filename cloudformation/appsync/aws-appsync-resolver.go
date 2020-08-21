package appsync

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// Resolver AWS CloudFormation Resource (AWS::AppSync::Resolver)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html
type Resolver struct {

	// ApiId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-apiid
	ApiId string `json:"ApiId,omitempty"`

	// CachingConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-cachingconfig
	CachingConfig *Resolver_CachingConfig `json:"CachingConfig,omitempty"`

	// DataSourceName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-datasourcename
	DataSourceName string `json:"DataSourceName,omitempty"`

	// FieldName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-fieldname
	FieldName string `json:"FieldName,omitempty"`

	// Kind AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-kind
	Kind string `json:"Kind,omitempty"`

	// PipelineConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-pipelineconfig
	PipelineConfig *Resolver_PipelineConfig `json:"PipelineConfig,omitempty"`

	// RequestMappingTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplate
	RequestMappingTemplate string `json:"RequestMappingTemplate,omitempty"`

	// RequestMappingTemplateS3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplates3location
	RequestMappingTemplateS3Location string `json:"RequestMappingTemplateS3Location,omitempty"`

	// ResponseMappingTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplate
	ResponseMappingTemplate string `json:"ResponseMappingTemplate,omitempty"`

	// ResponseMappingTemplateS3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplates3location
	ResponseMappingTemplateS3Location string `json:"ResponseMappingTemplateS3Location,omitempty"`

	// SyncConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-syncconfig
	SyncConfig *Resolver_SyncConfig `json:"SyncConfig,omitempty"`

	// TypeName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-typename
	TypeName string `json:"TypeName,omitempty"`

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
func (r *Resolver) AWSCloudFormationType() string {
	return "AWS::AppSync::Resolver"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r Resolver) MarshalJSON() ([]byte, error) {
	type Properties Resolver
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
func (r *Resolver) UnmarshalJSON(b []byte) error {
	type P Resolver
	props := &Resolver{}
	newProps := &struct {
		*P
		ApiId                             types.StringIsh `json:"ApiId,omitempty"`
		DataSourceName                    types.StringIsh `json:"DataSourceName,omitempty"`
		FieldName                         types.StringIsh `json:"FieldName,omitempty"`
		Kind                              types.StringIsh `json:"Kind,omitempty"`
		RequestMappingTemplate            types.StringIsh `json:"RequestMappingTemplate,omitempty"`
		RequestMappingTemplateS3Location  types.StringIsh `json:"RequestMappingTemplateS3Location,omitempty"`
		ResponseMappingTemplate           types.StringIsh `json:"ResponseMappingTemplate,omitempty"`
		ResponseMappingTemplateS3Location types.StringIsh `json:"ResponseMappingTemplateS3Location,omitempty"`
		TypeName                          types.StringIsh `json:"TypeName,omitempty"`
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

		props.ApiId = string(newProps.ApiId)
		props.DataSourceName = string(newProps.DataSourceName)
		props.FieldName = string(newProps.FieldName)
		props.Kind = string(newProps.Kind)
		props.RequestMappingTemplate = string(newProps.RequestMappingTemplate)
		props.RequestMappingTemplateS3Location = string(newProps.RequestMappingTemplateS3Location)
		props.ResponseMappingTemplate = string(newProps.ResponseMappingTemplate)
		props.ResponseMappingTemplateS3Location = string(newProps.ResponseMappingTemplateS3Location)
		props.TypeName = string(newProps.TypeName)

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
