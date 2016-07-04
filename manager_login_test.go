/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestForbiddenLogin(t *testing.T) {
	convey.Convey("Given I do a failed login", t, func() {
		server := mockRequest("/auth", "POST", 403, "")
		m := Manager{URL: server.URL}
		token, err := m.Login("foo", "bar")
		convey.Convey("Then I should receive an access denied error", func() {
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(token, convey.ShouldEqual, "")
		})
	})
}

func TestSuccessLogin(t *testing.T) {
	convey.Convey("Given I do a success login", t, func() {
		server := mockRequest("/auth", "POST", 200, `{"token": "valid-token"}`)
		m := Manager{URL: server.URL}
		token, err := m.Login("foo", "bar")
		convey.Convey("Then I should receive a valid token", func() {
			convey.So(err, convey.ShouldBeNil)
			convey.So(token, convey.ShouldEqual, "valid-token")
		})
	})
}
