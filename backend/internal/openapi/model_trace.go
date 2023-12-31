/*
 * API
 *
 * API for traces
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Trace struct {
	Uuid string `json:"uuid"`

	GroupUuid string `json:"group_uuid"`

	ParentUuid string `json:"parent_uuid,omitempty"`

	Start string `json:"start"`

	End string `json:"end"`

	Title string `json:"title"`

	Component string `json:"component"`

	Finished bool `json:"finished"`

	Children []Trace `json:"children"`

	Events []Event `json:"events"`
}

// AssertTraceRequired checks if the required fields are not zero-ed
func AssertTraceRequired(obj Trace) error {
	elements := map[string]interface{}{
		"uuid":       obj.Uuid,
		"group_uuid": obj.GroupUuid,
		"start":      obj.Start,
		"end":        obj.End,
		"title":      obj.Title,
		"component":  obj.Component,
		"finished":   obj.Finished,
		"children":   obj.Children,
		"events":     obj.Events,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Children {
		if err := AssertTraceRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Events {
		if err := AssertEventRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseTraceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Trace (e.g. [][]Trace), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTraceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTrace, ok := obj.(Trace)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTraceRequired(aTrace)
	})
}
