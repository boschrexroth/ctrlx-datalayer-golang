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

/*
Package token provides implementation of token handling in ctrlX Core.

JSON Web Token is used for authentification in ctrlX Core. Each HTTP request in ctrlX Core should be sent with a JWT token in the request header.
The token will be created after executing sucessful login in ctlrX Core. The token will be expired after a period, therefor it is important
always checking if the existing token is still valid before using it.

Package token implements following structs:
* struct Token
* struct TokenManager

struct Token:

A token in ctrlX Core has type "Bearer", token string is like "eyJhbGciOiJIU...".

String(): combines token type and token string together, which will be inserted into the section "Authorization" in the requst header.

struct TokenManager:

TokenManager implements login in the ctrlX Core and getting the token from the ctrlX Core. Username and Password are needed for login.

RequestAuthToken(): token will be created after login, TokenManager stores the token for latter usage.

CheckAuthToken(): checks the existing token if it is still valid.

So a good example to use TokenManager and Token looks like this:

	v, err := tokenManager.CheckAuthToken()
	if err != nil {
		// Something went wrong during token validation
		fmt.Println(err)
	}
	if !v {
		// Token will be created
		_, err = tokenManager.RequestAuthToken()
		if err != nil {
			return err
		}
	}
	fmt.Println(tokenManager.Token.String())

*/

package token
