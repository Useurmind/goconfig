package options

import (
	"strconv"
	"reflect"
	_ "os"
	"fmt"
	"strings"
)

type CMDLineOptionsSource struct {
	args []string
}

func NewCMDLineOptionsSource(args []string) CMDLineOptionsSource {
	return CMDLineOptionsSource{
		args: args,
	}
}

func (s *CMDLineOptionsSource) FillOptions(options interface{}) error {
	optValue := reflect.ValueOf(options)
	optType := reflect.TypeOf(optValue.Elem().Interface())
	fmt.Printf("Reflected value name is %s\r\n", optType.Name())

	for i:=1; i<len(s.args); i++ {
		arg := s.args[i]

		isOptionName := strings.HasPrefix(arg, "-")
		if isOptionName {
			optionName := strings.TrimLeft(arg, "-")

			optionField, err := GetFieldByOptionName(optType, optionName)
			if err != nil {
				return err
			}

			if optionField.Type.Kind() == reflect.Bool {
				// for flags we dont need a value, implicit value is true
				optValue.Elem().FieldByIndex(optionField.Index).SetBool(true)
			} else {
				if i + 1 == len(s.args) {
					return fmt.Errorf("Missing value for option %s", arg)
				}

				optionValueString := s.args[i + 1]
	
				switch optionField.Type.Kind() {
				case reflect.Int: 
					intValue, err := strconv.ParseInt(optionValueString, 10, 64)
					if err != nil {
						return fmt.Errorf("Could not convert option value '%s' for option '%s' to integer", optionValueString, optionName)
					}
					optValue.Elem().FieldByIndex(optionField.Index).SetInt(intValue)

				case reflect.String:
					optValue.Elem().FieldByIndex(optionField.Index).SetString(optionValueString)
				default:
					return fmt.Errorf("Option value parsing for type %s not implemented.", optionField.Type.Kind())
				}

			}
		}
	}

	return nil
}

func GetFieldByOptionName(optionType reflect.Type, optionName string) (*reflect.StructField, error) {
	for i:=0; i<optionType.NumField(); i++ {
		field := optionType.Field(i)

		tag, found := field.Tag.Lookup(OptionPropertyTag)
		if !found {
			name := strings.ToLower(field.Name)
			if name == optionName {
				return &field, nil
			}
		} else {
			names := strings.Split(tag, ",")
			for _, name := range names {
				if name == optionName {
					return &field, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("Could not find field for option '%s'", optionName)
}
