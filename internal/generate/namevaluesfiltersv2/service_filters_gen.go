// Code generated by generators/servicefilters/main.go; DO NOT EDIT.

package namevaluesfiltersv2

import ( // nosemgrep:ci.semgrep.aws.multiple-service-imports
	secretsmanagertypes "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

// []*SERVICE.Filter handling

// SecretsmanagerFilters returns secretsmanager service filters.
func (filters NameValuesFilters) SecretsmanagerFilters() []secretsmanagertypes.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]secretsmanagertypes.Filter, 0, len(m))

	for k, v := range m {
		filter := secretsmanagertypes.Filter{
			Key:    secretsmanagertypes.FilterNameStringType(k),
			Values: v,
		}

		result = append(result, filter)
	}

	return result
}