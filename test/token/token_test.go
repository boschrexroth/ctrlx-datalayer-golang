// MIT License
//
// Copyright (c) 2021-2022 Bosch Rexroth AG
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package token_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/token"
	"github.com/stretchr/testify/require"
)

func TestRequstAuthTokenBasic(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("{\"access_token\": \"eyJhbGciOiJIU.xxx.xxx\", \"token_type\":\"Bearer\"}"))
		require.NoError(t, err)
	}))
	defer server.Close()

	tm := &token.TokenManager{
		Connection: &http.Client{},
		Url:        server.URL,
		Username:   "user",
		Password:   "password",
	}

	token, err := tm.RequestAuthToken()
	require.NoError(t, err)

	if token.AccessToken != "eyJhbGciOiJIU.xxx.xxx" {
		t.Errorf("expect access_toke 'eyJhbGciOiJIU.xxx.xxx', but get %v\n", token.AccessToken)
	}
	if token.TokenType != "Bearer" {
		t.Errorf("expect token_type 'Bearer', but get %v\n", token.TokenType)
	}
}

func TestRequestAuthTokenMatrixDriven(t *testing.T) {
	var tests = []struct {
		res   string
		token token.Token
	}{
		{res: "{\"status\":400}", token: token.Token{AccessToken: "", TokenType: ""}},
		{res: "{\"status\":500}", token: token.Token{AccessToken: "", TokenType: ""}},
	}

	for _, test := range tests {
		t.Run(test.res, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte(test.res))
				require.NoError(t, err)
			}))
			defer server.Close()
			tm := &token.TokenManager{
				Connection: &http.Client{},
				Url:        server.URL,
				Username:   "user",
				Password:   "password",
			}
			token, err := tm.RequestAuthToken()
			require.NoError(t, err)
			if token.AccessToken != test.token.AccessToken {
				t.Errorf("test %v: expect access_toke %v, but get %v\n", test.res, test.token.AccessToken, token.AccessToken)
			}
			if token.TokenType != test.token.TokenType {
				t.Errorf("test %v: expect token_type %v, but get %v\n", test.res, test.token.TokenType, token.TokenType)
			}
		})
	}
}

func TestCheckAuthToken(t *testing.T) {
	// Handle request to validate token
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("{\"valid\": true}"))
		require.NoError(t, err)
	}))
	defer server.Close()

	tm := &token.TokenManager{
		Connection: &http.Client{},
		Url:        server.URL,
		Username:   "user",
		Password:   "password",
	}

	tm.Token = token.Token{
		TokenType:   "Bearer",
		AccessToken: "eyJhbGciOiJIU.xxx.xxx",
	}

	v, err := tm.CheckAuthToken()
	require.NoError(t, err)
	require.Equal(t, true, v)
}
