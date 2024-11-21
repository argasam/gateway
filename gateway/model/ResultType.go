package model

import "encoding/json"

type ResultType struct {
	StringValue string
	ObjectValue map[string]interface{}
	IsString    bool // To track the type
}

// MarshalJSON implements custom JSON marshaling
func (r ResultType) MarshalJSON() ([]byte, error) {
	if r.IsString {
		return json.Marshal(r.StringValue)
	}
	return json.Marshal(r.ObjectValue)
}
