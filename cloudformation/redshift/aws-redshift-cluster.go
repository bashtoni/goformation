package redshift

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
	"github.com/awslabs/goformation/v4/cloudformation/types"
)

// Cluster AWS CloudFormation Resource (AWS::Redshift::Cluster)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html
type Cluster struct {

	// AllowVersionUpgrade AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-allowversionupgrade
	AllowVersionUpgrade bool `json:"AllowVersionUpgrade,omitempty"`

	// AutomatedSnapshotRetentionPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-automatedsnapshotretentionperiod
	AutomatedSnapshotRetentionPeriod int `json:"AutomatedSnapshotRetentionPeriod,omitempty"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-availabilityzone
	AvailabilityZone string `json:"AvailabilityZone,omitempty"`

	// ClusterIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clusteridentifier
	ClusterIdentifier string `json:"ClusterIdentifier,omitempty"`

	// ClusterParameterGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clusterparametergroupname
	ClusterParameterGroupName string `json:"ClusterParameterGroupName,omitempty"`

	// ClusterSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustersecuritygroups
	ClusterSecurityGroups []string `json:"ClusterSecurityGroups,omitempty"`

	// ClusterSubnetGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustersubnetgroupname
	ClusterSubnetGroupName string `json:"ClusterSubnetGroupName,omitempty"`

	// ClusterType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustertype
	ClusterType string `json:"ClusterType,omitempty"`

	// ClusterVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clusterversion
	ClusterVersion string `json:"ClusterVersion,omitempty"`

	// DBName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-dbname
	DBName string `json:"DBName,omitempty"`

	// ElasticIp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-elasticip
	ElasticIp string `json:"ElasticIp,omitempty"`

	// Encrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-encrypted
	Encrypted bool `json:"Encrypted,omitempty"`

	// HsmClientCertificateIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-hsmclientcertidentifier
	HsmClientCertificateIdentifier string `json:"HsmClientCertificateIdentifier,omitempty"`

	// HsmConfigurationIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-HsmConfigurationIdentifier
	HsmConfigurationIdentifier string `json:"HsmConfigurationIdentifier,omitempty"`

	// IamRoles AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-iamroles
	IamRoles []string `json:"IamRoles,omitempty"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-kmskeyid
	KmsKeyId string `json:"KmsKeyId,omitempty"`

	// LoggingProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-loggingproperties
	LoggingProperties *Cluster_LoggingProperties `json:"LoggingProperties,omitempty"`

	// MasterUserPassword AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-masteruserpassword
	MasterUserPassword string `json:"MasterUserPassword,omitempty"`

	// MasterUsername AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-masterusername
	MasterUsername string `json:"MasterUsername,omitempty"`

	// NodeType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-nodetype
	NodeType string `json:"NodeType,omitempty"`

	// NumberOfNodes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-nodetype
	NumberOfNodes int `json:"NumberOfNodes,omitempty"`

	// OwnerAccount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-owneraccount
	OwnerAccount string `json:"OwnerAccount,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-port
	Port int `json:"Port,omitempty"`

	// PreferredMaintenanceWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-preferredmaintenancewindow
	PreferredMaintenanceWindow string `json:"PreferredMaintenanceWindow,omitempty"`

	// PubliclyAccessible AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-publiclyaccessible
	PubliclyAccessible bool `json:"PubliclyAccessible,omitempty"`

	// SnapshotClusterIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-snapshotclusteridentifier
	SnapshotClusterIdentifier string `json:"SnapshotClusterIdentifier,omitempty"`

	// SnapshotIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-snapshotidentifier
	SnapshotIdentifier string `json:"SnapshotIdentifier,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

	// VpcSecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-vpcsecuritygroupids
	VpcSecurityGroupIds []string `json:"VpcSecurityGroupIds,omitempty"`

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
func (r *Cluster) AWSCloudFormationType() string {
	return "AWS::Redshift::Cluster"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.'
func (r Cluster) MarshalJSON() ([]byte, error) {
	type Properties Cluster
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
func (r *Cluster) UnmarshalJSON(b []byte) error {
	type P Cluster
	props := &Cluster{}
	newProps := &struct {
		*P
		AvailabilityZone               types.StringIsh `json:"AvailabilityZone,omitempty"`
		ClusterIdentifier              types.StringIsh `json:"ClusterIdentifier,omitempty"`
		ClusterParameterGroupName      types.StringIsh `json:"ClusterParameterGroupName,omitempty"`
		ClusterSubnetGroupName         types.StringIsh `json:"ClusterSubnetGroupName,omitempty"`
		ClusterType                    types.StringIsh `json:"ClusterType,omitempty"`
		ClusterVersion                 types.StringIsh `json:"ClusterVersion,omitempty"`
		DBName                         types.StringIsh `json:"DBName,omitempty"`
		ElasticIp                      types.StringIsh `json:"ElasticIp,omitempty"`
		HsmClientCertificateIdentifier types.StringIsh `json:"HsmClientCertificateIdentifier,omitempty"`
		HsmConfigurationIdentifier     types.StringIsh `json:"HsmConfigurationIdentifier,omitempty"`
		KmsKeyId                       types.StringIsh `json:"KmsKeyId,omitempty"`
		MasterUserPassword             types.StringIsh `json:"MasterUserPassword,omitempty"`
		MasterUsername                 types.StringIsh `json:"MasterUsername,omitempty"`
		NodeType                       types.StringIsh `json:"NodeType,omitempty"`
		OwnerAccount                   types.StringIsh `json:"OwnerAccount,omitempty"`
		PreferredMaintenanceWindow     types.StringIsh `json:"PreferredMaintenanceWindow,omitempty"`
		SnapshotClusterIdentifier      types.StringIsh `json:"SnapshotClusterIdentifier,omitempty"`
		SnapshotIdentifier             types.StringIsh `json:"SnapshotIdentifier,omitempty"`
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

		props.AvailabilityZone = string(newProps.AvailabilityZone)
		props.ClusterIdentifier = string(newProps.ClusterIdentifier)
		props.ClusterParameterGroupName = string(newProps.ClusterParameterGroupName)
		props.ClusterSubnetGroupName = string(newProps.ClusterSubnetGroupName)
		props.ClusterType = string(newProps.ClusterType)
		props.ClusterVersion = string(newProps.ClusterVersion)
		props.DBName = string(newProps.DBName)
		props.ElasticIp = string(newProps.ElasticIp)
		props.HsmClientCertificateIdentifier = string(newProps.HsmClientCertificateIdentifier)
		props.HsmConfigurationIdentifier = string(newProps.HsmConfigurationIdentifier)
		props.KmsKeyId = string(newProps.KmsKeyId)
		props.MasterUserPassword = string(newProps.MasterUserPassword)
		props.MasterUsername = string(newProps.MasterUsername)
		props.NodeType = string(newProps.NodeType)
		props.OwnerAccount = string(newProps.OwnerAccount)
		props.PreferredMaintenanceWindow = string(newProps.PreferredMaintenanceWindow)
		props.SnapshotClusterIdentifier = string(newProps.SnapshotClusterIdentifier)
		props.SnapshotIdentifier = string(newProps.SnapshotIdentifier)

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
