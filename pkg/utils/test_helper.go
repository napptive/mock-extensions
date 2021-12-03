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
package utils

import (
	"github.com/rs/xid"
	"syreclabs.com/go/faker"
)

// FullExample with a struct to use it in the tests
type FullExample struct {
	UserId    string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Cont      int    `protobuf:"bytes,1,opt,name=cont,proto3" json:"con_tab,omitempty"`
	Suspended bool   `protobuf:"bytes,1,opt,name=sus,proto3" json:"suspend,omitempty"`
}

// NewFullExample returns a filled FullExample
func NewFullExample(userId string, username string, cont int, sus bool) FullExample {
	return FullExample{
		UserId:    userId,
		Username:  username,
		Cont:      cont,
		Suspended: sus,
	}
}

// NakedExample defines a struct without annotations
type NakedExample struct {
	UserId    string
	Username  string
	Cont      int
	Suspended bool
}

// NewNakedExample returns a struct without annotations
func NewNakedExample(userId string, username string, cont int, sus bool) NakedExample {
	return NakedExample{
		UserId:    userId,
		Username:  username,
		Cont:      cont,
		Suspended: sus,
	}
}

// GetUserId generates a random UserId
func GetUserId() string {
	return xid.New().String()
}

// GetUserName generates a random username
func GetUserName() string {
	return faker.Internet().UserName()
}

// JsonExample with a struct with json annotations (no protobuf annotations)
type JsonExample struct {
	UserId    string `json:"user_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Cont      int    `json:"con_tab,omitempty"`
	Suspended bool   `json:"suspend,omitempty"`
}

// NewJsonExample returns a struct with json annotations
func NewJsonExample(userId string, username string, cont int, sus bool) JsonExample {
	return JsonExample{
		UserId:    userId,
		Username:  username,
		Cont:      cont,
		Suspended: sus,
	}
}
