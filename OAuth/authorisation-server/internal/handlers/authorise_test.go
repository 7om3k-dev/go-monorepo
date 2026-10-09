package handlers_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/7om3k-dev/OAuth/internal/handlers"
)

const clientIdOk, maliciousClientId = "oauth-client-1", "malicious-client-id"
const redirectUriPathOk, maliciousRedirectUriPath = "oauth-redirect", "malicious-redirect"

func createAuthorizeUrl(t *testing.T, values url.Values) url.URL {
	t.Helper()
	return url.URL{
		Scheme:   "http",
		Host:     "localhost:8099",
		Path:     "authorize",
		RawQuery: values.Encode(),
	}
}

func createRedirectUri(t *testing.T, path string) string {
	t.Helper()
	u := url.URL{
		Scheme: "http",
		Host:   "example.com",
		Path:   path,
	}
	return u.String()
}

func TestAuthorise(t *testing.T) {
	type urlValuesAddParam struct {
		key   string
		value string
	}

	clientIdKey, redirectUriPathKey := "client-id", "redirect-uri"

	testCases := []struct {
		name                string
		queryParams         []urlValuesAddParam
		expectedPageContent string
	}{
		{
			"Successful",
			[]urlValuesAddParam{
				{clientIdKey, clientIdOk},
				{redirectUriPathKey, redirectUriPathOk},
			},
			"Sign in",
		},
		{
			"Missing client ID",
			[]urlValuesAddParam{{redirectUriPathKey, redirectUriPathOk}},
			"Error",
		},
		{
			"Missing redirect uri",
			[]urlValuesAddParam{{clientIdKey, clientIdOk}},
			"Error",
		},
		{
			"Invalid client ID",
			[]urlValuesAddParam{
				{clientIdKey, maliciousClientId},
				{redirectUriPathKey, redirectUriPathOk},
			},
			"Error",
		},
		{
			"Invalid redirect uri",
			[]urlValuesAddParam{
				{clientIdKey, clientIdOk},
				{redirectUriPathKey, maliciousRedirectUriPath},
			},
			"Error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			urlValues := url.Values{}

			for _, queryParam := range tc.queryParams {
				if queryParam.key == redirectUriPathKey {
					urlValues.Set(queryParam.key, createRedirectUri(t, queryParam.value))
					continue
				}
				urlValues.Set(queryParam.key, queryParam.value)
			}

			u := createAuthorizeUrl(t, urlValues)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, u.String(), nil)

			handlers.Authorise(w, r)
			resp := w.Result()

			wantContentType := "text/html; charset=utf-8"
			if resp.Header.Get("Content-Type") != wantContentType {
				t.Errorf("Authorise returned wrong content type: got %v want %v", resp.Header.Get("Content-Type"), wantContentType)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Error reading response body")
			}

			if !bytes.Contains(body, []byte(tc.expectedPageContent)) {
				t.Errorf("Authorise returned wrong html page")
				t.Errorf("Body: %s", string(body))
			}
		})
	}
}
