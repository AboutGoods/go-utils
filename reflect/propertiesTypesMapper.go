// Package reflect provides methods based on reflection
package reflect

import (
	"reflect"
	"strings"
)

// PropertiesTypesMapper returns a map giving information of entity properties
// (key: property name (lowercase), value: property type)
func PropertiesTypesMapper(entity interface{}) map[string]string {
	e := reflect.ValueOf(entity).Elem()
	propTypes := map[string]string{}

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		propTypes[strings.ToLower(varName)] = varType.String()
	}

	return propTypes
}
