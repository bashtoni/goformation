package intrinsics

import (
	"fmt"
	"strings"
)

// FnGetAtt is not implemented, and always returns nil.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html
func FnGetAtt(name string, input interface{}, template interface{}) interface{} {
	// { "Fn::GetAtt" : [ "logicalNameOfResource", "attributeName" ] }
	switch input.(type) {
	case []string:
		getAttInput := input.([]string)
		refRsc := getAttInput[0]
		switch getAttInput[1] {
		case "Arn":
			return generateARN("123456789012", "us-east-1", refRsc, template)
		}
	case string:
		// Perhaps we got the input in logicalNameOfResource.attributeName format?
		if inputStr, ok := input.(string); ok {
			split := strings.Split(inputStr, ".")
			if len(split) == 2 {
				return FnGetAtt(name, split, template)
			}
		}
	case []interface{}:
		inputs := input.([]interface{})
		s := make([]string, len(inputs))
		for i, v := range inputs {
			s[i] = fmt.Sprint(v)
		}
		return FnGetAtt(name, s, template)
	}
	return nil
}
