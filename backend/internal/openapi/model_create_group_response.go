/*
 * API
 *
 * API for traces
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateGroupResponse struct {
	Uuid string `json:"uuid"`
}

// AssertCreateGroupResponseRequired checks if the required fields are not zero-ed
func AssertCreateGroupResponseRequired(obj CreateGroupResponse) error {
	elements := map[string]interface{}{
		"uuid": obj.Uuid,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCreateGroupResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreateGroupResponse (e.g. [][]CreateGroupResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreateGroupResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreateGroupResponse, ok := obj.(CreateGroupResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreateGroupResponseRequired(aCreateGroupResponse)
	})
}
