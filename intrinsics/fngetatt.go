package intrinsics

// FnGetAtt is not implemented, and always returns nil.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html
func FnGetAtt(name string, input interface{}, template interface{}) interface{} {
	// { "Fn::GetAtt" : [ "logicalNameOfResource", "attributeName" ] }
	if getAttInput, ok := input.([]string); ok {
		refRsc := getAttInput[0]
		switch getAttInput[1] {
		case "Arn":
			// find the referenced resource (refRsc) in the template, and generate an ARN that is valid for its type
			if template, ok := template.(map[string]interface{}); ok {
				// Check there is a resources section
				if uresources, ok := template["Resources"]; ok {
					// Check the resources section is a map
					if resources, ok := uresources.(map[string]interface{}); ok {
						// Check there is a resource with the name we're looking for
						if uresource, ok := resources[refRsc]; ok {
							// Check the resource is a map
							if resource, ok := uresource.(map[string]interface{}); ok {
								// Check the resource has a type
								if uresourceType, ok := resource["Type"]; ok {
									// Return a generated ARN for the resource
									if resourceType, ok := uresourceType.(string); ok {
										return generateARN("123456789012", "us-east-1", resourceType, refRsc)
									}
								}
							}
						}

					}
				}
			}
		}
	}
	return nil
}
