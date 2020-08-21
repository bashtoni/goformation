package ec2

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// Route AWS CloudFormation Resource (AWS::EC2::Route)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html
type Route struct {

	// DestinationCidrBlock AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationcidrblock
	DestinationCidrBlock string `json:"DestinationCidrBlock,omitempty"`

	// DestinationIpv6CidrBlock AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationipv6cidrblock
	DestinationIpv6CidrBlock string `json:"DestinationIpv6CidrBlock,omitempty"`

	// EgressOnlyInternetGatewayId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-egressonlyinternetgatewayid
	EgressOnlyInternetGatewayId string `json:"EgressOnlyInternetGatewayId,omitempty"`

	// GatewayId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-gatewayid
	GatewayId string `json:"GatewayId,omitempty"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-instanceid
	InstanceId string `json:"InstanceId,omitempty"`

	// NatGatewayId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-natgatewayid
	NatGatewayId string `json:"NatGatewayId,omitempty"`

	// NetworkInterfaceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-networkinterfaceid
	NetworkInterfaceId string `json:"NetworkInterfaceId,omitempty"`

	// RouteTableId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-routetableid
	RouteTableId string `json:"RouteTableId,omitempty"`

	// TransitGatewayId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-transitgatewayid
	TransitGatewayId string `json:"TransitGatewayId,omitempty"`

	// VpcPeeringConnectionId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-vpcpeeringconnectionid
	VpcPeeringConnectionId string `json:"VpcPeeringConnectionId,omitempty"`

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
func (r *Route) AWSCloudFormationType() string {
	return "AWS::EC2::Route"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r Route) MarshalJSON() ([]byte, error) {
	type Properties Route
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
func (r *Route) UnmarshalJSON(b []byte) error {
	type P Route
	props := &Route{}
	newProps := &struct {
		*P
		DestinationCidrBlock        types.StringIsh `json:"DestinationCidrBlock,omitempty"`
		DestinationIpv6CidrBlock    types.StringIsh `json:"DestinationIpv6CidrBlock,omitempty"`
		EgressOnlyInternetGatewayId types.StringIsh `json:"EgressOnlyInternetGatewayId,omitempty"`
		GatewayId                   types.StringIsh `json:"GatewayId,omitempty"`
		InstanceId                  types.StringIsh `json:"InstanceId,omitempty"`
		NatGatewayId                types.StringIsh `json:"NatGatewayId,omitempty"`
		NetworkInterfaceId          types.StringIsh `json:"NetworkInterfaceId,omitempty"`
		RouteTableId                types.StringIsh `json:"RouteTableId,omitempty"`
		TransitGatewayId            types.StringIsh `json:"TransitGatewayId,omitempty"`
		VpcPeeringConnectionId      types.StringIsh `json:"VpcPeeringConnectionId,omitempty"`
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

		props.DestinationCidrBlock = string(newProps.DestinationCidrBlock)
		props.DestinationIpv6CidrBlock = string(newProps.DestinationIpv6CidrBlock)
		props.EgressOnlyInternetGatewayId = string(newProps.EgressOnlyInternetGatewayId)
		props.GatewayId = string(newProps.GatewayId)
		props.InstanceId = string(newProps.InstanceId)
		props.NatGatewayId = string(newProps.NatGatewayId)
		props.NetworkInterfaceId = string(newProps.NetworkInterfaceId)
		props.RouteTableId = string(newProps.RouteTableId)
		props.TransitGatewayId = string(newProps.TransitGatewayId)
		props.VpcPeeringConnectionId = string(newProps.VpcPeeringConnectionId)

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
