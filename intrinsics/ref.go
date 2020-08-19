package intrinsics

import "fmt"

// Ref resolves the 'Ref' AWS CloudFormation intrinsic function.
// Currently, this only resolves against CloudFormation Parameter default values
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html
func Ref(name string, input interface{}, template interface{}) interface{} {

	// Dang son, this has got more nest than a bald eagle
	// Check the input is a string
	if name, ok := input.(string); ok {

		switch name {

		case "AWS::AccountId":
			return "123456789012"
		case "AWS::NotificationARNs": //
			return []string{"arn:aws:sns:us-east-1:123456789012:MyTopic"}
		case "AWS::NoValue":
			return nil
		case "AWS::Region":
			return "us-east-1"
		case "AWS::StackId":
			return "arn:aws:cloudformation:us-east-1:123456789012:stack/MyStack/1c2fa620-982a-11e3-aff7-50e2416294e0"
		case "AWS::StackName":
			return "goformation-stack"

		default:

			// This isn't a pseudo 'Ref' paramater, so we need to look inside the CloudFormation template
			// to see if we can resolve the reference. This implementation just looks at the Parameters section
			// to see if there is a parameter matching the name, and if so, return the default value.

			// Check the template is a map
			if template, ok := template.(map[string]interface{}); ok {
				// Check there is a parameters section
				if uparameters, ok := template["Parameters"]; ok {
					// Check the parameters section is a map
					if parameters, ok := uparameters.(map[string]interface{}); ok {
						// Check there is a parameter with the same name as the Ref
						if uparameter, ok := parameters[name]; ok {
							// Check the parameter is a map
							if parameter, ok := uparameter.(map[string]interface{}); ok {
								// Check the parameter has a default
								if def, ok := parameter["Default"]; ok {
									return def
								}
							}
						}
					}
				}

				// Return the name based on the resource type
				resource := getResource(name, template)
				if resource == nil {
					return nil
				}
				return RefForResource(name, resource, template)
			}
		}
	}

	return nil

}

func RefForResource(name string, resource map[string]interface{}, template interface{}) string {
	cfnType := resourceType(name, resource)
	if uproperties, ok := resource["Properties"]; ok {
		if properties, ok := uproperties.(map[string]interface{}); ok {
			switch cfnType {
			case "AWS::IAM::Role":
				if tRoleName, ok := properties["RoleName"].(string); ok {
					return tRoleName
				}
				// The resource could itself contain a ref..
				if roleRef, ok := properties["RoleName"].([]string); ok {
					if realRoleName, ok := Ref(roleRef[0], roleRef[1], template).(string); ok {
						return realRoleName
					}
				}
				// Ref might appear like this
				if uroleRef, ok := properties["RoleName"].(map[string]interface{}); ok {
					if refTo, ok := uroleRef["Ref"]; ok {
						if realRoleName, ok := Ref("Ref", refTo, template).(string); ok {
							return realRoleName
						}
					}
				}
				return name
			case "AWS::IAM::ManagedPolicy":
				polName := name
				if tPolName, ok := properties["ManagedPolicyName"].(string); ok {
					polName = tPolName
				}
				return fmt.Sprintf("arn:aws:iam::123456789012:policy/%s", polName)
			}
		}
	}
	return name
}
