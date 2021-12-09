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
	"github.com/napptive/mock-extensions/pkg/utils"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"reflect"
)

var _ = ginkgo.Describe("Testing struct matcher", func() {
	ginkgo.Context("Checking Matches with a Full Example", func() {
		ginkgo.It("Matches using all types (field, proto and json)", func() {

			user := utils.NewFullExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{
				// json and proto name
				"user_id": user.UserId,
				// Field name
				"Username": user.Username,
				// proto name
				"sus": true,
				// json name
				"con_tab": 10,
			}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).Should(gomega.BeTrue())

		})
		ginkgo.It("Matches using all types (field, proto and json) with a pointer", func() {

			user := utils.NewFullExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{
				// json and proto name
				"user_id": user.UserId,
				// Field name
				"Username": user.Username,
				// proto name
				"sus": true,
				// json name
				"con_tab": 10,
			}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(&user)
			gomega.Expect(res).Should(gomega.BeTrue())

		})
		ginkgo.It("no matches when a field is not found", func() {

			user := utils.NewFullExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{"Index": user.UserId}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).ShouldNot(gomega.BeTrue())

		})
		ginkgo.It("no matches when the value is no correct", func() {

			user := utils.NewFullExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{"user_id": utils.GetUserId()}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).ShouldNot(gomega.BeTrue())

		})
	})
	ginkgo.Context("Checking Matches with an example without annotations", func() {
		ginkgo.It("Matches", func() {

			user := utils.NewNakedExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{
				"UserId":    user.UserId,
				"Username":  user.Username,
				"Suspended": true,
				"Cont":      10,
			}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).Should(gomega.BeTrue())

		})
		ginkgo.It("no matches when a field is not found", func() {

			user := utils.NewNakedExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{"Index": user.UserId}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).ShouldNot(gomega.BeTrue())

		})
		ginkgo.It("no matches when the value is no correct", func() {

			user := utils.NewNakedExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{"UserId": utils.GetUserId()}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).ShouldNot(gomega.BeTrue())

		})
	})
	ginkgo.Context("Checking Matches with an example without json annotations", func() {
		ginkgo.It("Matches", func() {

			user := utils.NewJsonExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{
				// json name
				"user_id": user.UserId,
				// Field name
				"Username": user.Username,
				// Field name
				"Suspended": true,
				// json name
				"con_tab": 10,
			}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).Should(gomega.BeTrue())

		})
		ginkgo.It("no matches when a field is not found", func() {

			user := utils.NewJsonExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{"Index": user.UserId}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).ShouldNot(gomega.BeTrue())

		})
		ginkgo.It("no matches when the value is no correct", func() {

			user := utils.NewJsonExample(utils.GetUserId(), utils.GetUserName(), 10, true)
			pattern := map[string]interface{}{"user_id": utils.GetUserId()}

			matcher := NewStructMatcher(pattern)

			res := matcher.Matches(user)
			gomega.Expect(res).ShouldNot(gomega.BeTrue())

		})
	})
	ginkgo.Context("Checking GetFieldName", func() {
		ginkgo.It("all the annotations filled", func() {
			type testStruct struct {
				UserId string `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
			}
			test := testStruct{UserId: utils.GetUserId()}
			name, json, proto := GetFieldName(reflect.ValueOf(test).Type().Field(0))
			gomega.Expect(name).NotTo(gomega.BeEmpty())
			gomega.Expect(name).Should(gomega.Equal("UserId"))
			gomega.Expect(json).NotTo(gomega.BeEmpty())
			gomega.Expect(proto).NotTo(gomega.BeEmpty())
			gomega.Expect(json).Should(gomega.Equal(proto))
		})
		ginkgo.It("proto annotation not filled", func() {
			type testStruct struct {
				UserId string `json:"user_id,omitempty"`
			}
			test := testStruct{UserId: utils.GetUserId()}
			name, json, proto := GetFieldName(reflect.ValueOf(test).Type().Field(0))
			gomega.Expect(name).NotTo(gomega.BeEmpty())
			gomega.Expect(name).Should(gomega.Equal("UserId"))
			gomega.Expect(json).NotTo(gomega.BeEmpty())
			gomega.Expect(json).Should(gomega.Equal("user_id"))
			gomega.Expect(proto).To(gomega.BeEmpty())
		})
		ginkgo.It("json annotation not filled", func() {
			type testStruct struct {
				UserId string `protobuf:"bytes,1,opt,name=user_id,proto3"`
			}
			test := testStruct{UserId: utils.GetUserId()}
			name, json, proto := GetFieldName(reflect.ValueOf(test).Type().Field(0))
			gomega.Expect(name).NotTo(gomega.BeEmpty())
			gomega.Expect(name).Should(gomega.Equal("UserId"))
			gomega.Expect(json).To(gomega.BeEmpty())
			gomega.Expect(proto).NotTo(gomega.BeEmpty())
			gomega.Expect(proto).Should(gomega.Equal("user_id"))
		})
		ginkgo.It("annotation not filled without some data", func() {
			type testStruct struct {
				UserId string `protobuf:"name=user_id,proto3" json:"user_id"`
			}
			test := testStruct{UserId: utils.GetUserId()}
			name, json, proto := GetFieldName(reflect.ValueOf(test).Type().Field(0))
			gomega.Expect(name).NotTo(gomega.BeEmpty())
			gomega.Expect(name).Should(gomega.Equal("UserId"))
			gomega.Expect(json).NotTo(gomega.BeEmpty())
			gomega.Expect(proto).NotTo(gomega.BeEmpty())
			gomega.Expect(json).Should(gomega.Equal(proto))
		})
	})
})
