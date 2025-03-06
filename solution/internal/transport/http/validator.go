package http

import (
	"fmt"
	"reflect"
	"server/pkg/errors/validate"
	"strings"
)

type CustomValidator struct{}

func (v *CustomValidator) Validate(i interface{}) error {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)
		tags := strings.Split(field.Tag.Get("validate"), ",")
		if len(tags) == 0 {
			continue
		}
		if ContaintTags(tags, "required") {
			switch fieldVal.Kind() {
			case reflect.Ptr:
				if fieldVal.IsNil() {
					return validate.NewValidationError(fmt.Sprintf("field %s is required", field.Name))
				}
				elem := fieldVal.Elem()
				if elem.Kind() == reflect.String && elem.String() == "" {
					return validate.NewValidationError(fmt.Sprintf("field %s is required", field.Name))
				}
			case reflect.Struct:
				if err := v.Validate(fieldVal.Interface()); err != nil {
					return err
				}
			}
		} else if fieldVal.Kind() == reflect.Struct {
			if err := v.Validate(fieldVal.Interface()); err != nil {
				return err
			}
		}
	}
	return nil
}

func ContaintTags(tags []string, tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}
