package types

import (
	"encoding/json"
	"fmt"
)

type StringIsh string

func (s *StringIsh) UnmarshalJson(b []byte) error {
	var raw interface{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	*s = StringIsh(fmt.Sprintf("%v", raw))
	return nil
}
