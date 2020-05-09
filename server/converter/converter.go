package converter

import (
	"github.com/onepanelio/core/api"
	v1 "github.com/onepanelio/core/pkg"
)

func APIKeyValueToLabel(apiKeyValues []*api.KeyValue) map[string]string {
	result := make(map[string]string)
	if apiKeyValues == nil {
		return result
	}

	for _, entry := range apiKeyValues {
		result[entry.Key] = entry.Value
	}

	return result
}

func MappingToKeyValue(mapping map[string]string) []*api.KeyValue {
	keyValues := make([]*api.KeyValue, 0)

	for key, value := range mapping {
		keyValues = append(keyValues, &api.KeyValue{
			Key:   key,
			Value: value,
		})
	}

	return keyValues
}

func ParameterOptionToAPI(option v1.ParameterOption) *api.ParameterOption {
	apiOption := &api.ParameterOption{
		Name:  option.Name,
		Value: option.Value,
	}

	return apiOption
}

func ParameterOptionsToAPI(options []*v1.ParameterOption) []*api.ParameterOption {
	result := make([]*api.ParameterOption, len(options))

	for i := range options {
		newItem := ParameterOptionToAPI(*options[i])
		result[i] = newItem
	}

	return result
}

func ParameterToAPI(param v1.Parameter) *api.Parameter {
	apiParam := &api.Parameter{
		Name:     param.Name,
		Type:     param.Type,
		Required: param.Required,
	}

	if param.Value != nil {
		apiParam.Value = *param.Value
	}
	if param.DisplayName != nil {
		apiParam.DisplayName = *param.DisplayName
	}
	if param.Hint != nil {
		apiParam.Hint = *param.Hint
	}

	if param.Options != nil {
		apiParam.Options = ParameterOptionsToAPI(param.Options)
	}

	return apiParam
}

func ParametersToAPI(params []v1.Parameter) []*api.Parameter {
	result := make([]*api.Parameter, len(params))

	for i := range params {
		newItem := ParameterToAPI(params[i])
		result[i] = newItem
	}

	return result
}