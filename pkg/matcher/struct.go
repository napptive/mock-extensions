/**
 * Copyright 2021 Napptive
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package matcher

import (
	"reflect"
	"strings"
)

// StructMatcher defines a generic structure matcher. Implements Matcher interface
type StructMatcher struct {
	// pattern with a map with the name of the field an its value
	pattern map[string]interface{}
}

// NewStructMatcher creates a generic structure matcher with a given pattern
func NewStructMatcher(pattern map[string]interface{}) StructMatcher {
	return StructMatcher{pattern: pattern}
}

// Matches returns whether arg is a match.
func (sm StructMatcher) Matches(arg interface{}) bool {

	for k, v := range sm.pattern {
		// Check if other has the field, and matches type and value.
		// Other has field if
		// 1. Has a field named K
		// 2. Has a field with a json annotation named K
		// 3. Has a field with a proto annotation named K
		// Example:
		// Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
		//    2^^^                                        3^^^                  1^^^^^

		value := reflect.ValueOf(arg)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		found := false
		for i := 0; i < value.NumField() && !found; i++ { // iterates through every struct type field
			name, json, proto := GetFieldName(value.Type().Field(i))
			if k == name || k == json || k == proto {
				found = true
				if value.Field(i).Type() != reflect.ValueOf(v).Type() || value.Field(i).Interface() != v {
					return false
				}
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// String describes what the matcher matches. Not used but necessary to satisfy Matcher interface
func (sm StructMatcher) String() string {
	return "A matcher for arbitrary struct types"
}

// GetFieldName returns the name of the file, the json name and the protobuf name indicated in the annotations
func GetFieldName(t reflect.StructField) (string, string, string) {
	var name, jsonName, protoName string

	name = t.Name

	if jsonTag := t.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
		// check for possible comma as in "json:"username,omitempty""
		var ind int
		if ind = strings.Index(jsonTag, ","); ind < 0 {
			ind = len(jsonTag)
		}
		jsonName = jsonTag[:ind]

	}
	if protoTag := t.Tag.Get("protobuf"); protoTag != "" && protoTag != "-" {
		// get the name from protobuf:"bytes,1,opt,name=username,proto3"
		var ind int
		if ind = strings.Index(protoTag, "name="); ind < 0 {
			ind = len(protoTag)
		}
		fieldName := protoTag[ind:]
		if ind = strings.Index(fieldName, ","); ind < 0 {
			ind = len(fieldName)
		}
		fieldName = fieldName[5:ind]
		protoName = fieldName
	}

	return name, jsonName, protoName
}
