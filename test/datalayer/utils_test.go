/**
 * MIT License
 *
 * Copyright (c) 2021-2022 Bosch Rexroth AG
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package datalayer_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type globalEnv struct {
	addr             string
	sslport          string
	timeout          string
	alldata_provider string
}

var globalEnvValue globalEnv

func init() {
	globalEnvValue.addr = os.Getenv("CTRLX_ADDRESS")
	globalEnvValue.sslport = os.Getenv("CTRLX_SSL_PORT")
	globalEnvValue.timeout = os.Getenv("CTRLX_TIMEOUT")
	globalEnvValue.alldata_provider = os.Getenv("ALLDATA_PROVIDER")
}

func resetGlobalEnvs() {
	//prevents side effects on other tests
	os.Setenv("CTRLX_ADDRESS", globalEnvValue.addr)
	os.Setenv("CTRLX_SSL_PORT", globalEnvValue.sslport)
	os.Setenv("CTRLX_TIMEOUT", globalEnvValue.timeout)
	os.Setenv("ALLDATA_PROVIDER", globalEnvValue.alldata_provider)
}

// ctrlxAddress - returns the TCP url to connect to the Data Layer of a ctrlX.
// Cannot be used in snap environment!
// There is no need to provide client/provider port (2069/2070) because the
// datalayer component uses automatically the according port number.
func ctrlxAddress() string {
	addr := os.Getenv("CTRLX_ADDRESS")
	if addr == "" || addr == "-" {
		return ""
	}
	url := fmt.Sprintf("tcp://boschrexroth:boschrexroth@%s", addr)

	sslport := os.Getenv("CTRLX_SSL_PORT")
	if sslport == "" || sslport == "-" {
		return url
	}

	return url + "?sslport=" + sslport
}

func ctrlxClientTimeout() uint {
	env := os.Getenv("CTRLX_TIMEOUT")
	if env == "" || env == "-" {
		return 2000
	}

	t, _ := strconv.ParseUint(env, 10, 32)
	return uint(t)
}

func allDataProvider() string {
	env := os.Getenv("ALLDATA_PROVIDER")
	if env == "" || env == "-" {
		return ""
	}
	return env
}

func asyncTestTimeout() time.Duration {
	t := ctrlxClientTimeout()
	if t == uint(2000) {
		return 5 * time.Second
	}
	t = t / 1000
	t *= 2
	return time.Duration(t) * time.Second
}

func TestClientAddressNoTests(t *testing.T) {
	os.Setenv("CTRLX_ADDRESS", "-")
	e := ctrlxAddress()
	assert.Equal(t, e, "")
	resetGlobalEnvs()
}

func TestClientAddressSetEnv(t *testing.T) {
	os.Setenv("CTRLX_ADDRESS", "10.0.2.2")
	e := ctrlxAddress()
	assert.Equal(t, e, "tcp://boschrexroth:boschrexroth@10.0.2.2")
	resetGlobalEnvs()
}

func TestClientAddressSetEnv2(t *testing.T) {
	os.Setenv("CTRLX_ADDRESS", "10.0.2.2")
	os.Setenv("CTRLX_SSL_PORT", "8443")
	e := ctrlxAddress()
	assert.Equal(t, e, "tcp://boschrexroth:boschrexroth@10.0.2.2?sslport=8443")
	resetGlobalEnvs()
}

func TestClientTimeout(t *testing.T) {
	os.Setenv("CTRLX_TIMEOUT", "-")
	e := ctrlxClientTimeout()
	assert.Equal(t, e, uint(2000))
	resetGlobalEnvs()
}

func TestClientTimeoutSetEnv(t *testing.T) {
	os.Setenv("CTRLX_TIMEOUT", "4000")
	e := ctrlxClientTimeout()
	assert.Equal(t, e, uint(4000))
	resetGlobalEnvs()
}

func TestAsyncWaitTimeout(t *testing.T) {
	os.Setenv("CTRLX_TIMEOUT", "-")
	e := asyncTestTimeout()
	assert.Equal(t, e, 5*time.Second)
	resetGlobalEnvs()
}

func TestAsyncWaitTimeoutSetEnv(t *testing.T) {
	os.Setenv("CTRLX_TIMEOUT", "4000")
	e := asyncTestTimeout()
	assert.Equal(t, e, (4000/1000)*2*time.Second)
	resetGlobalEnvs()
}
