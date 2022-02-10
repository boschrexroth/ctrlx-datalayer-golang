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

package datalayer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcResult = []struct {
	result         Result
	expectedString string
}{
	{
		ResultOk,
		"DL_OK",
	},
	{
		ResultFailed,
		"DL_FAILED",
	},
	{
		ResultInvalidAddress,
		"DL_INVALID_ADDRESS",
	},
	{
		ResultUnsupported,
		"DL_UNSUPPORTED",
	},
	{
		ResultOutOfMemory,
		"DL_OUT_OF_MEMORY",
	},
	{
		ResultLimitMin,
		"DL_LIMIT_MIN",
	},
	{
		ResultLimitMax,
		"DL_LIMIT_MAX",
	},
	{
		ResultTypeMismatch,
		"DL_TYPE_MISMATCH",
	},
	{
		ResultSizeMismatch,
		"DL_SIZE_MISMATCH",
	},
	{
		ResultInvalidFloatingpoint,
		"DL_INVALID_FLOATINGPOINT",
	},
	{
		ResultInvalidHandle,
		"DL_INVALID_HANDLE",
	},
	{
		ResultInvalidOperationMode,
		"DL_INVALID_OPERATION_MODE",
	},
	{
		ResultInvalidConfiguration,
		"DL_INVALID_CONFIGURATION",
	},
	{
		ResultInvalidValue,
		"DL_INVALID_VALUE",
	},
	{
		ResultSubmoduleFailure,
		"DL_SUBMODULE_FAILURE",
	},
	{
		ResultTimeout,
		"DL_TIMEOUT",
	},
	{
		ResultAlreadyExists,
		"DL_ALREADY_EXISTS",
	},
	{
		ResultCreationFailed,
		"DL_CREATION_FAILED",
	},
	{
		ResultVersionMismatch,
		"DL_VERSION_MISMATCH",
	},
	{
		ResultDeprecated,
		"DL_DEPRECATED",
	},
	{
		ResultPermissionDenied,
		"DL_PERMISSION_DENIED",
	},
	{
		ResultNotInitialized,
		"DL_NOT_INITIALIZED",
	},
	{
		ResultCommProtocolError,
		"DL_COMM_PROTOCOL_ERROR",
	},
	{
		ResultCommInvalidHeader,
		"DL_COMM_INVALID_HEADER",
	},
	{
		ResultClientNotConnected,
		"DL_CLIENT_NOT_CONNECTED",
	},
	{
		ResultRtNotopen,
		"DL_RT_NOTOPEN",
	},
	{
		ResultRtInvalidobject,
		"DL_RT_INVALIDOBJECT",
	},
	{
		ResultRtWrongrevison,
		"DL_RT_WRONGREVISON",
	},
	{
		ResultRtNovaliddata,
		"DL_RT_NOVALIDDATA",
	},
	{
		ResultRtMemorylocked,
		"DL_RT_MEMORYLOCKED",
	},
	{
		ResultRtInvalidmemorymap,
		"DL_RT_INVALIDMEMORYMAP",
	},
	{
		ResultRtInvalidRetain,
		"DL_RT_INVALID_RETAIN",
	},
	{
		ResultRtInternalError,
		"DL_RT_INTERNAL_ERROR",
	},
	{
		ResultSecNotoken,
		"DL_SEC_NOTOKEN",
	},
	{
		ResultSecInvalidsession,
		"DL_SEC_INVALIDSESSION",
	},
	{
		ResultSecInvalidtokencontent,
		"DL_SEC_INVALIDTOKENCONTENT",
	},
	{
		ResultSecUnauthorized,
		"DL_SEC_UNAUTHORIZED",
	},
}

func TestResult(t *testing.T) {
	for _, tc := range tcResult {
		assert.Equal(t, tc.expectedString, tc.result.String())
	}
}
