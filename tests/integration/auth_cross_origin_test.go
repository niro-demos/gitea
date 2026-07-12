// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"testing"

	"gitea.dev/models/unittest"
	user_model "gitea.dev/models/user"
	"gitea.dev/modules/setting"
	"gitea.dev/modules/test"
	"gitea.dev/tests"
)

func crossOriginAuthRequest(t *testing.T, path string, values map[string]string) *RequestWrapper {
	t.Helper()
	req := NewRequestWithValues(t, "POST", path, values)
	req.Header.Set("Origin", "https://evil.example")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	return req
}

func TestCrossOriginAuthPosts(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	defer test.MockVariableValue(&setting.Service.EnableCaptcha, false)()

	t.Run("login", func(t *testing.T) {
		user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2})
		values := map[string]string{"user_name": user.Name, "password": userPassword}

		legitimateSession := emptyTestSession(t)
		legitimateSession.MakeRequest(t, NewRequestWithValues(t, "POST", "/user/login", values), http.StatusSeeOther)
		legitimateSession.MakeRequest(t, NewRequest(t, "GET", "/user/settings"), http.StatusOK)

		attackedSession := emptyTestSession(t)
		attackedSession.MakeRequest(t, crossOriginAuthRequest(t, "/user/login", values), http.StatusForbidden)
		attackedSession.MakeRequest(t, NewRequest(t, "GET", "/user/settings"), http.StatusSeeOther)
	})

	t.Run("sign up", func(t *testing.T) {
		legitimateValues := map[string]string{
			"user_name": "legitimate-signup",
			"email":     "legitimate-signup@example.com",
			"password":  "examplePassword!1",
			"retype":    "examplePassword!1",
		}
		legitimateSession := emptyTestSession(t)
		legitimateSession.MakeRequest(t, NewRequestWithValues(t, "POST", "/user/sign_up", legitimateValues), http.StatusSeeOther)
		legitimateSession.MakeRequest(t, NewRequest(t, "GET", "/user/settings"), http.StatusOK)

		attackedValues := map[string]string{
			"user_name": "cross-origin-signup",
			"email":     "cross-origin-signup@example.com",
			"password":  "examplePassword!1",
			"retype":    "examplePassword!1",
		}
		attackedSession := emptyTestSession(t)
		attackedSession.MakeRequest(t, crossOriginAuthRequest(t, "/user/sign_up", attackedValues), http.StatusForbidden)
		attackedSession.MakeRequest(t, NewRequest(t, "GET", "/user/settings"), http.StatusSeeOther)
		unittest.AssertNotExistsBean(t, &user_model.User{Name: "cross-origin-signup"})
	})
}
