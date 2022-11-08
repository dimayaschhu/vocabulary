package utils

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

type ObjectMatcher struct {
	placeholderValidator *PlaceholderValidator
}

func NewObjectMatcher(placeholderValidator *PlaceholderValidator) *ObjectMatcher {
	return &ObjectMatcher{
		placeholderValidator: placeholderValidator,
	}
}

func (m *ObjectMatcher) MatchJson(actualJSON []byte, expectedJSON []byte) error {
	var actual, expected map[string]interface{}

	if err := json.Unmarshal(actualJSON, &actual); err != nil {
		return err
	}

	if err := json.Unmarshal(expectedJSON, &expected); err != nil {
		return err
	}

	if err := m.Match(actual, expected); err != nil {
		return err
	}

	return nil
}

func (m *ObjectMatcher) Match(actual map[string]interface{}, expected map[string]interface{}) error {
	filledExpected, err := m.checkAndFillPlaceholders(expected, actual)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(actual, filledExpected) {
		return errors.Errorf(
			"Objects should be equal.\nActual:\n%s\nExpected:\n%s",
			Marshal(actual),
			Marshal(expected),
		)
	}

	return nil
}

func (m *ObjectMatcher) checkAndFillPlaceholders(
	mapWithEmptyFields map[string]interface{},
	mapWithData map[string]interface{},
) (map[string]interface{}, error) {
	for key, value := range mapWithEmptyFields {
		reflectedValue := reflect.TypeOf(value)

		if reflectedValue == nil {
			continue
		}

		switch reflectedValue.Kind() {
		case reflect.Map:
			_, err := m.checkAndFillPlaceholders(
				mapWithEmptyFields[key].(map[string]interface{}),
				mapWithData[key].(map[string]interface{}),
			)
			if err != nil {
				return nil, err
			}
		case reflect.Slice:
			for i, el := range value.([]interface{}) {
				k, ok := el.(map[string]interface{})

				if ok {
					_, err := m.checkAndFillPlaceholders(
						k,
						mapWithData[key].([]interface{})[i].(map[string]interface{}),
					)
					if err != nil {
						return nil, err
					}
				}
			}
		default:
			actualStringValue, isString := mapWithData[key].(string)

			if isString && m.placeholderValidator.IsPlaceholderValue(value.(string)) {
				if err := m.placeholderValidator.Validate(value.(string), actualStringValue); err != nil {
					return nil, err
				}

				mapWithEmptyFields[key] = mapWithData[key]
			}
		}
	}

	return mapWithEmptyFields, nil
}
