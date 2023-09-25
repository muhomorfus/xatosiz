/*
 * API
 *
 * API for traces
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetEventListResponse struct {
	Items []Event `json:"items"`
}

// AssertGetEventListResponseRequired checks if the required fields are not zero-ed
func AssertGetEventListResponseRequired(obj GetEventListResponse) error {
	elements := map[string]interface{}{
		"items": obj.Items,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Items {
		if err := AssertEventRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseGetEventListResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetEventListResponse (e.g. [][]GetEventListResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetEventListResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetEventListResponse, ok := obj.(GetEventListResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetEventListResponseRequired(aGetEventListResponse)
	})
}