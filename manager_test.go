/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/smartystreets/goconvey/convey"
)

func mockRequest(route string, method string, status int, output string) *httptest.Server {
	r := mux.NewRouter()
	r.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		s := output
		if s == "" {
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			s = buf.String()
		}
		w.WriteHeader(status)
		w.Header().Set("X-Auth-Token", "")
		fmt.Fprint(w, s)
	}).Methods(method)

	return httptest.NewServer(r)
}

func TestForbiddenLogin(t *testing.T) {
	convey.Convey("Given I do a failed login", t, func() {
		server := mockRequest("/session/", "POST", 403, "")
		m := Manager{URL: server.URL}
		body, token, err := m.Login("foo", "bar")
		convey.Convey("Then I should receive an access denied error", func() {
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(body, convey.ShouldEqual, `{"user_name":"foo", "user_password": "bar"}`)
			convey.So(token, convey.ShouldEqual, "")
		})
	})
}

func TestSuccessLogin(t *testing.T) {
	convey.Convey("Given I do a success login", t, func() {
		server := mockRequest("/session/", "POST", 200, ``)
		m := Manager{URL: server.URL}
		body, token, err := m.Login("foo", "bar")
		convey.Convey("Then I should receive a valid token", func() {
			convey.So(err, convey.ShouldBeNil)
			convey.So(body, convey.ShouldEqual, ``)
			convey.So(token, convey.ShouldEqual, `foo`)
		})
	})
}

func TestSuccessLogout(t *testing.T) {
	convey.Convey("Given I do a success logout", t, func() {
		server := mockRequest("/session/", "DELETE", 200, ``)
		m := Manager{URL: server.URL}
		err := m.Logout("foo")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestCreateUser(t *testing.T) {
	t.Skip()
	convey.Convey("Given I create a client", t, func() {
		server := mockRequest("/clients/", "POST", 200, ``)
		m := Manager{URL: server.URL}
		err := m.CreateUser("name", "email", "user", "password", "adminuser", "adminpassword")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

// ********************* Destroy *******************

func TestDestroy(t *testing.T) {
	convey.Convey("Given a service", t, func() {
		server := mockRequest("/services/foo", "DELETE", 200, ``)
		m := Manager{URL: server.URL}
		err := m.Destroy("token", "foo", false)
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

// ********************* Reset *******************

func TestResetService(t *testing.T) {
	convey.Convey("Given a service", t, func() {
		server := mockRequest("/services/foo/reset/", "POST", 200, ``)
		m := Manager{URL: server.URL}
		err := m.ResetService("foo", "token")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

// ********************* Status *******************

func TestServiceStatus(t *testing.T) {
	convey.Convey("Given a service", t, func() {
		server := mockRequest("/services/foo", "GET", 200, ``)
		m := Manager{URL: server.URL}
		_, err := m.ServiceStatus("token", "foo")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestServiceBuildStatus(t *testing.T) {
	convey.Convey("Given a service and a build id", t, func() {
		server := mockRequest("/services/foo/builds/1234567890", "GET", 200, ``)
		m := Manager{URL: server.URL}
		_, err := m.ServiceBuildStatus("token", "foo", "1234567890")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

// ********************* List *********************

func TestListDatacenters(t *testing.T) {
	convey.Convey("Given I get all datacenters", t, func() {
		server := mockRequest("/datacenters/", "GET", 200, ``)
		m := Manager{URL: server.URL}
		_, err := m.ListDatacenters("token")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestListServices(t *testing.T) {
	convey.Convey("Given I get all services", t, func() {
		server := mockRequest("/services/", "GET", 200, ``)
		m := Manager{URL: server.URL}
		_, err := m.ListServices("token")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestListBuilds(t *testing.T) {
	convey.Convey("Given a service", t, func() {
		convey.Convey("Given I get all builds", func() {
			server := mockRequest("/services/foo/builds/", "GET", 200, ``)
			m := Manager{URL: server.URL}
			_, err := m.ListBuilds("foo", "token")
			convey.Convey("Then It does not fail", func() {
				convey.So(err, convey.ShouldBeNil)
			})
		})
	})
}

func TestListUsers(t *testing.T) {
	convey.Convey("Given I get all users", t, func() {
		server := mockRequest("/users/", "GET", 200, ``)
		m := Manager{URL: server.URL}
		_, err := m.ListUsers("token")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestListGroups(t *testing.T) {
	convey.Convey("Given I get all groups", t, func() {
		server := mockRequest("/clients/", "GET", 200, ``)
		m := Manager{URL: server.URL}
		_, err := m.ListGroups("token")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}