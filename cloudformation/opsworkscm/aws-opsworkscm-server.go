package opsworkscm

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// Server AWS CloudFormation Resource (AWS::OpsWorksCM::Server)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html
type Server struct {

	// AssociatePublicIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-associatepublicipaddress
	AssociatePublicIpAddress bool `json:"AssociatePublicIpAddress,omitempty"`

	// BackupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-backupid
	BackupId string `json:"BackupId,omitempty"`

	// BackupRetentionCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-backupretentioncount
	BackupRetentionCount int `json:"BackupRetentionCount,omitempty"`

	// CustomCertificate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-customcertificate
	CustomCertificate string `json:"CustomCertificate,omitempty"`

	// CustomDomain AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-customdomain
	CustomDomain string `json:"CustomDomain,omitempty"`

	// CustomPrivateKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-customprivatekey
	CustomPrivateKey string `json:"CustomPrivateKey,omitempty"`

	// DisableAutomatedBackup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-disableautomatedbackup
	DisableAutomatedBackup bool `json:"DisableAutomatedBackup,omitempty"`

	// Engine AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-engine
	Engine string `json:"Engine,omitempty"`

	// EngineAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-engineattributes
	EngineAttributes []Server_EngineAttribute `json:"EngineAttributes,omitempty"`

	// EngineModel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-enginemodel
	EngineModel string `json:"EngineModel,omitempty"`

	// EngineVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-engineversion
	EngineVersion string `json:"EngineVersion,omitempty"`

	// InstanceProfileArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-instanceprofilearn
	InstanceProfileArn string `json:"InstanceProfileArn,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-instancetype
	InstanceType string `json:"InstanceType,omitempty"`

	// KeyPair AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-keypair
	KeyPair string `json:"KeyPair,omitempty"`

	// PreferredBackupWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-preferredbackupwindow
	PreferredBackupWindow string `json:"PreferredBackupWindow,omitempty"`

	// PreferredMaintenanceWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-preferredmaintenancewindow
	PreferredMaintenanceWindow string `json:"PreferredMaintenanceWindow,omitempty"`

	// SecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-securitygroupids
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`

	// ServerName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-servername
	ServerName string `json:"ServerName,omitempty"`

	// ServiceRoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-servicerolearn
	ServiceRoleArn string `json:"ServiceRoleArn,omitempty"`

	// SubnetIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-subnetids
	SubnetIds []string `json:"SubnetIds,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworkscm-server.html#cfn-opsworkscm-server-tags
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
func (r *Server) AWSCloudFormationType() string {
	return "AWS::OpsWorksCM::Server"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r Server) MarshalJSON() ([]byte, error) {
	type Properties Server
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
func (r *Server) UnmarshalJSON(b []byte) error {
	type P Server
	props := &Server{}
	newProps := &struct {
		*P
		BackupId                   types.StringIsh `json:"BackupId,omitempty"`
		CustomCertificate          types.StringIsh `json:"CustomCertificate,omitempty"`
		CustomDomain               types.StringIsh `json:"CustomDomain,omitempty"`
		CustomPrivateKey           types.StringIsh `json:"CustomPrivateKey,omitempty"`
		Engine                     types.StringIsh `json:"Engine,omitempty"`
		EngineModel                types.StringIsh `json:"EngineModel,omitempty"`
		EngineVersion              types.StringIsh `json:"EngineVersion,omitempty"`
		InstanceProfileArn         types.StringIsh `json:"InstanceProfileArn,omitempty"`
		InstanceType               types.StringIsh `json:"InstanceType,omitempty"`
		KeyPair                    types.StringIsh `json:"KeyPair,omitempty"`
		PreferredBackupWindow      types.StringIsh `json:"PreferredBackupWindow,omitempty"`
		PreferredMaintenanceWindow types.StringIsh `json:"PreferredMaintenanceWindow,omitempty"`
		ServerName                 types.StringIsh `json:"ServerName,omitempty"`
		ServiceRoleArn             types.StringIsh `json:"ServiceRoleArn,omitempty"`
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

		props.BackupId = string(newProps.BackupId)
		props.CustomCertificate = string(newProps.CustomCertificate)
		props.CustomDomain = string(newProps.CustomDomain)
		props.CustomPrivateKey = string(newProps.CustomPrivateKey)
		props.Engine = string(newProps.Engine)
		props.EngineModel = string(newProps.EngineModel)
		props.EngineVersion = string(newProps.EngineVersion)
		props.InstanceProfileArn = string(newProps.InstanceProfileArn)
		props.InstanceType = string(newProps.InstanceType)
		props.KeyPair = string(newProps.KeyPair)
		props.PreferredBackupWindow = string(newProps.PreferredBackupWindow)
		props.PreferredMaintenanceWindow = string(newProps.PreferredMaintenanceWindow)
		props.ServerName = string(newProps.ServerName)
		props.ServiceRoleArn = string(newProps.ServiceRoleArn)

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
