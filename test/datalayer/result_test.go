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

package datalayer_test

import (
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	"github.com/stretchr/testify/assert"
)

var tcResult = []struct {
	result         datalayer.Result
	expectedString string
}{
	{
		datalayer.ResultOk,
		"DL_OK",
	},
	{
		datalayer.ResultFailed,
		"DL_FAILED",
	},
	{
		datalayer.ResultInvalidAddress,
		"DL_INVALID_ADDRESS",
	},
	{
		datalayer.ResultUnsupported,
		"DL_UNSUPPORTED",
	},
	{
		datalayer.ResultOutOfMemory,
		"DL_OUT_OF_MEMORY",
	},
	{
		datalayer.ResultLimitMin,
		"DL_LIMIT_MIN",
	},
	{
		datalayer.ResultLimitMax,
		"DL_LIMIT_MAX",
	},
	{
		datalayer.ResultTypeMismatch,
		"DL_TYPE_MISMATCH",
	},
	{
		datalayer.ResultSizeMismatch,
		"DL_SIZE_MISMATCH",
	},
	{
		datalayer.ResultInvalidFloatingpoint,
		"DL_INVALID_FLOATINGPOINT",
	},
	{
		datalayer.ResultInvalidHandle,
		"DL_INVALID_HANDLE",
	},
	{
		datalayer.ResultInvalidOperationMode,
		"DL_INVALID_OPERATION_MODE",
	},
	{
		datalayer.ResultInvalidConfiguration,
		"DL_INVALID_CONFIGURATION",
	},
	{
		datalayer.ResultInvalidValue,
		"DL_INVALID_VALUE",
	},
	{
		datalayer.ResultSubmoduleFailure,
		"DL_SUBMODULE_FAILURE",
	},
	{
		datalayer.ResultTimeout,
		"DL_TIMEOUT",
	},
	{
		datalayer.ResultAlreadyExists,
		"DL_ALREADY_EXISTS",
	},
	{
		datalayer.ResultCreationFailed,
		"DL_CREATION_FAILED",
	},
	{
		datalayer.ResultVersionMismatch,
		"DL_VERSION_MISMATCH",
	},
	{
		datalayer.ResultDeprecated,
		"DL_DEPRECATED",
	},
	{
		datalayer.ResultPermissionDenied,
		"DL_PERMISSION_DENIED",
	},
	{
		datalayer.ResultNotInitialized,
		"DL_NOT_INITIALIZED",
	},
	{
		datalayer.ResultCommProtocolError,
		"DL_COMM_PROTOCOL_ERROR",
	},
	{
		datalayer.ResultCommInvalidHeader,
		"DL_COMM_INVALID_HEADER",
	},
	{
		datalayer.ResultClientNotConnected,
		"DL_CLIENT_NOT_CONNECTED",
	},
	{
		datalayer.ResultRtNotopen,
		"DL_RT_NOTOPEN",
	},
	{
		datalayer.ResultRtInvalidobject,
		"DL_RT_INVALIDOBJECT",
	},
	{
		datalayer.ResultRtWrongrevison,
		"DL_RT_WRONGREVISON",
	},
	{
		datalayer.ResultRtNovaliddata,
		"DL_RT_NOVALIDDATA",
	},
	{
		datalayer.ResultRtMemorylocked,
		"DL_RT_MEMORYLOCKED",
	},
	{
		datalayer.ResultRtInvalidmemorymap,
		"DL_RT_INVALIDMEMORYMAP",
	},
	{
		datalayer.ResultRtInvalidRetain,
		"DL_RT_INVALID_RETAIN",
	},
	{
		datalayer.ResultRtInternalError,
		"DL_RT_INTERNAL_ERROR",
	},
	{
		datalayer.ResultSecNotoken,
		"DL_SEC_NOTOKEN",
	},
	{
		datalayer.ResultSecInvalidsession,
		"DL_SEC_INVALIDSESSION",
	},
	{
		datalayer.ResultSecInvalidtokencontent,
		"DL_SEC_INVALIDTOKENCONTENT",
	},
	{
		datalayer.ResultSecUnauthorized,
		"DL_SEC_UNAUTHORIZED",
	},
	{
		datalayer.ResultOkNoContent,
		"DL_OK_NO_CONTENT",
	},
	{
		datalayer.ResultMissingArgument,
		"DL_MISSING_ARGUMENT",
	},
	{
		datalayer.ResultTooManyArguments,
		"DL_TOO_MANY_ARGUMENTS",
	},
	{
		datalayer.ResultResourceUnavailable,
		"DL_RESOURCE_UNAVAILABLE",
	},
	{
		datalayer.ResultCommunicationError,
		"DL_COMMUNICATION_ERROR",
	},
	{
		datalayer.ResultTooManyOperations,
		"DL_TOO_MANY_OPERATIONS",
	},
	{
		datalayer.ResultWouldBlock,
		"DL_WOULD_BLOCK",
	},
	{
		datalayer.ResultPaymentRequired,
		"DL_SEC_PAYMENT_REQUIRED",
	},
	{
		datalayer.Result(1701),
		"Result_Unknown",
	},
}

func TestResult(t *testing.T) {
	for _, tc := range tcResult {
		assert.Equal(t, tc.expectedString, tc.result.String())
	}
}

func TestResultAsError(t *testing.T) {
	result := datalayer.ResultFailed
	assert.Error(t, result)
}
