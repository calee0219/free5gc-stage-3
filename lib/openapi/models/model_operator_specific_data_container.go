/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type OperatorSpecificDataContainer struct {
	StringTypeElements  map[string]string  `json:"StringTypeElements,omitempty" bson:"StringTypeElements"`
	IntegerTypeElements map[string]int32   `json:"IntegerTypeElements,omitempty" bson:"IntegerTypeElements"`
	NumberTypeElements  map[string]float32 `json:"NumberTypeElements,omitempty" bson:"NumberTypeElements"`
	BooleanTypeElements map[string]bool    `json:"BooleanTypeElements,omitempty" bson:"BooleanTypeElements"`
}