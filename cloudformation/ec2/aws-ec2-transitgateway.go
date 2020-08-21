package ec2

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// TransitGateway AWS CloudFormation Resource (AWS::EC2::TransitGateway)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html
type TransitGateway struct {

	// AmazonSideAsn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-amazonsideasn
	AmazonSideAsn int `json:"AmazonSideAsn,omitempty"`

	// AutoAcceptSharedAttachments AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-autoacceptsharedattachments
	AutoAcceptSharedAttachments string `json:"AutoAcceptSharedAttachments,omitempty"`

	// DefaultRouteTableAssociation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-defaultroutetableassociation
	DefaultRouteTableAssociation string `json:"DefaultRouteTableAssociation,omitempty"`

	// DefaultRouteTablePropagation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-defaultroutetablepropagation
	DefaultRouteTablePropagation string `json:"DefaultRouteTablePropagation,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-description
	Description string `json:"Description,omitempty"`

	// DnsSupport AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-dnssupport
	DnsSupport string `json:"DnsSupport,omitempty"`

	// MulticastSupport AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-multicastsupport
	MulticastSupport string `json:"MulticastSupport,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

	// VpnEcmpSupport AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-vpnecmpsupport
	VpnEcmpSupport string `json:"VpnEcmpSupport,omitempty"`

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
func (r *TransitGateway) AWSCloudFormationType() string {
	return "AWS::EC2::TransitGateway"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r TransitGateway) MarshalJSON() ([]byte, error) {
	type Properties TransitGateway
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
func (r *TransitGateway) UnmarshalJSON(b []byte) error {
	type P TransitGateway
	props := &TransitGateway{}
	newProps := &struct {
		*P
		AutoAcceptSharedAttachments  types.StringIsh `json:"AutoAcceptSharedAttachments,omitempty"`
		DefaultRouteTableAssociation types.StringIsh `json:"DefaultRouteTableAssociation,omitempty"`
		DefaultRouteTablePropagation types.StringIsh `json:"DefaultRouteTablePropagation,omitempty"`
		Description                  types.StringIsh `json:"Description,omitempty"`
		DnsSupport                   types.StringIsh `json:"DnsSupport,omitempty"`
		MulticastSupport             types.StringIsh `json:"MulticastSupport,omitempty"`
		VpnEcmpSupport               types.StringIsh `json:"VpnEcmpSupport,omitempty"`
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

		props.AutoAcceptSharedAttachments = string(newProps.AutoAcceptSharedAttachments)
		props.DefaultRouteTableAssociation = string(newProps.DefaultRouteTableAssociation)
		props.DefaultRouteTablePropagation = string(newProps.DefaultRouteTablePropagation)
		props.Description = string(newProps.Description)
		props.DnsSupport = string(newProps.DnsSupport)
		props.MulticastSupport = string(newProps.MulticastSupport)
		props.VpnEcmpSupport = string(newProps.VpnEcmpSupport)

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
