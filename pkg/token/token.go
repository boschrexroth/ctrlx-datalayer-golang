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

package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Before any communication can be established to the ctrlX device it is necessary to authenticate first.
// After successfull authentication a token will be received, that can be used in all subsequent calls for
// authorization.
// This file is for requesting this token and for checking if it is still valid or might be outdated.
// The implementation is based on the 'ctrlX CORE - Authorization and Authentication API'.
// The API description can be found at:
// https://boschrexroth.github.io/rest-api-description/ctrlx-automation/ctrlx-core/

const (
	validateTokenPath = "/identity-manager/api/v2/auth/token/validity"
	identityPath      = "/identity-manager/api/v2/auth/token"
)

// Token represents the json structure that will be returned from the device after successfull authentication.
// The token can be used for further API calls and needs to be encoded in the 'authorization' field of the
// http header.
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// validToken respresents the json structure that will be returned by the device
// after a request to check if the given token is still valid.
type validToken struct {
	Valid bool `json:"valid"`
}

// String builds the token string which is used in the 'authorization' field of the http header.
func (t *Token) String() string {
	return t.TokenType + " " + t.AccessToken
}

type TokenManager struct {
	Url        string
	Username   string
	Password   string
	Connection *http.Client

	mux   sync.Mutex // mux used to serialize authorization requests
	Token Token
}

// CheckAuthToken checks if the tokenManager has valid authentication token.
func (m *TokenManager) CheckAuthToken() (isValid bool, err error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if m.Token.AccessToken == "" {
		return false, nil
	}

	u := m.Url + validateTokenPath

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request to validate token: %w", err)
	}

	req.Header.Add("Authorization", m.Token.String())

	resp, err := m.Connection.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to execute request to validate token: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read response to validate token: %w", err)
	}
	var vt validToken
	err = json.Unmarshal(body, &vt)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal validate token response: %w", err)
	}
	return vt.Valid, nil
}

// RequestAuthToken gets a valid authentication token from the ctrlX CORE using username and password.
// The token is necessary to create a subscription. An error will be returned, when no token could get aquired.
func (m *TokenManager) RequestAuthToken() (Token, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.Token = Token{}

	tokenURL := m.Url + identityPath
	payload, err := json.Marshal(map[string]string{
		"name":     m.Username,
		"password": m.Password,
	})
	if err != nil {
		return m.Token, fmt.Errorf("failed to create authorization token request: %w", err)
	}
	requestBody := bytes.NewBuffer(payload)

	resp, err := m.Connection.Post(tokenURL, "application/json", requestBody)
	if err != nil {
		r := strings.NewReplacer(identityPath, "/...")
		b := r.Replace(err.Error())
		return m.Token, fmt.Errorf("an error occured during request of authorization token: %q", b)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return m.Token, fmt.Errorf("error while reading the response bytes: %w", err)
	}

	err = json.Unmarshal(body, &m.Token)
	if err != nil {
		return m.Token, fmt.Errorf("an error occured during unmarshal of authorization body: %w", err)
	}

	return m.Token, nil
}
