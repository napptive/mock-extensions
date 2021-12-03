# mock-extensions
mock-extensions is a library that includes extensions to be used in the mockups

- `StructMatcher` implements Matcher interface.

## How to use
1. Create a StructMatcher as
```
matcher := NewStructMatcher(map[string]interface{}{"<field_name>": <field_value>...})
```
-`field_name` should be field name in the struct, proto annotation or json annotation. For example, if you have a struct defined as following:
```
type Object struct {
	Field1    string `protobuf:"bytes,1,opt,name=field_proto1,proto3" json:"field_1,omitempty"`
	Field2    int    `protobuf:"bytes,1,opt,name=field_proto2,proto3" json:"field_2,omitempty"`
	...
	Fieldn    string `protobuf:"bytes,1,opt,name=field_proton,proto3" json:"field_n,omitempty"`
}
```
you could use different matches:

1. With the field names
```
matcher := NewStructMatcher(map[string]interface{}{"Field1": <field_value>, "Field2": <field_value>})
```
2. With proto annotations
```
matcher := NewStructMatcher(map[string]interface{}{"field_proto1": <field_value>, "field_proto2": <field_value>})
```
3. With json annotations
```
matcher := NewStructMatcher(map[string]interface{}{"field_1": <field_value>, "field_2": <field_value>})
```
4. Or combining the field names
```
matcher := NewStructMatcher(map[string]interface{}{"Field1": <field_value>, "field_2": <field_value>})
```

## Integration with Github Actions

This template is integrated with GitHub Actions.

![Check changes in the Main branch](https://github.com/napptive/mock-extensions/workflows/Check%20changes%20in%20the%20Main%20branch/badge.svg)

## License

 Copyright 2020 Napptive

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
