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

//Returns a map of fields/values from an entity.
//Nested objects are flatten ( ["var.subVar" => "foo", "var.anotherOne" => "bar"] )
func GetFlatValuesMap(entity interface{}) map[string]interface{} {
	value := reflect.ValueOf(entity)
	valType := value.Type()
	nbFields := value.NumField()

	valuesMap := make(map[string]interface{})

	for i := 0; i < nbFields; i++ {

		fieldName := valType.Field(i).Name

		if value.Field(i).Kind() != reflect.Struct {
			valuesMap[fieldName] = value.Field(i)
		}

		if value.Field(i).Kind() == reflect.Struct {
			for name, val := range GetFlatValuesMap(value.Field(i).Interface()) {
				valuesMap[fieldName+"."+name] = val
			}
		}
	}

	return valuesMap

}
