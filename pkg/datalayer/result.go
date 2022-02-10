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

//#include <stdbool.h>
//#include <stdlib.h>
//#include <system.h>
import "C"

// Result ulong
type Result C.DLR_RESULT

const (
	ResultOk     Result = C.DL_OK     //  FUNCTION CALL SUCCEEDED
	ResultFailed Result = C.DL_FAILED //  ONLY ALLOWED FOR TEMPORARY USE - DEFINE MATCHING ERROR CODE

	ResultInvalidAddress       Result = C.DL_INVALID_ADDRESS        //  ADDRESS NOT FOUND, ADDRESS INVALID (BROWSE OF THIS NODE NOT POSSIBLE, WRITE -> ADDRESS NOT VALID)
	ResultUnsupported          Result = C.DL_UNSUPPORTED            //  FUNCTION NOT IMPLEMENTED
	ResultOutOfMemory          Result = C.DL_OUT_OF_MEMORY          //  OUT OF MEMORY OR RESOURCES (RAM, SOCKETS, HANDLES, DISK SPACE ...).
	ResultLimitMin             Result = C.DL_LIMIT_MIN              //  THE MINIMUM OF A LIMITATION IS EXCEEDED.
	ResultLimitMax             Result = C.DL_LIMIT_MAX              //  THE MAXIMUM OF A LIMITATION IS EXCEEDED.
	ResultTypeMismatch         Result = C.DL_TYPE_MISMATCH          //  WRONG FLATBUFFER TYPE, WRONG DATA TYPE
	ResultSizeMismatch         Result = C.DL_SIZE_MISMATCH          //  SIZE MISMATCH, PRESENT SIZE DOESN'T MATCH REQUESTED SIZE.
	ResultInvalidFloatingpoint Result = C.DL_INVALID_FLOATINGPOINT  //  INVALID FLOATING POINT NUMBER.
	ResultInvalidHandle        Result = C.DL_INVALID_HANDLE         //  INVALID HANDLE ARGUMENT OR NULL POINTER ARGUMENT.
	ResultInvalidOperationMode Result = C.DL_INVALID_OPERATION_MODE //  NOT ACCESSIBLE DUE TO INVALID OPERATION MODE (WRITE NOT POSSIBLE)
	ResultInvalidConfiguration Result = C.DL_INVALID_CONFIGURATION  //  MISMATCH OF THIS VALUE WITH OTHER CONFIGURED VALUES
	ResultInvalidValue         Result = C.DL_INVALID_VALUE          //  INVALID VALUE
	ResultSubmoduleFailure     Result = C.DL_SUBMODULE_FAILURE      //  ERROR IN SUBMODULE
	ResultTimeout              Result = C.DL_TIMEOUT                //  REQUEST TIMEOUT
	ResultAlreadyExists        Result = C.DL_ALREADY_EXISTS         //  CREATE: RESOURCE ALREADY EXISTS
	ResultCreationFailed       Result = C.DL_CREATION_FAILED        //  ERROR DURING CREATION
	ResultVersionMismatch      Result = C.DL_VERSION_MISMATCH       //  VERSION CONFLICT
	ResultDeprecated           Result = C.DL_DEPRECATED             //  DEPRECATED - FUNCTION NOT LONGER SUPPORTED
	ResultPermissionDenied     Result = C.DL_PERMISSION_DENIED      //  REQUEST DECLINED DUE TO MISSING PERMISSION RIGHTS
	ResultNotInitialized       Result = C.DL_NOT_INITIALIZED        //  OBJECT NOT INITIALIZED YET

	ResultCommProtocolError Result = C.DL_COMM_PROTOCOL_ERROR //  INTERNAL PROTOCOL ERROR
	ResultCommInvalidHeader Result = C.DL_COMM_INVALID_HEADER //  INTERNAL HEADER MISMATCH

	ResultClientNotConnected Result = C.DL_CLIENT_NOT_CONNECTED //  CLIENT NOT CONNECTED

	ResultRtNotopen          Result = C.DL_RT_NOTOPEN          //  NOT OPEN
	ResultRtInvalidobject    Result = C.DL_RT_INVALIDOBJECT    //  INVALID OBJECT
	ResultRtWrongrevison     Result = C.DL_RT_WRONGREVISON     //  WRONG MEMORY REVISION
	ResultRtNovaliddata      Result = C.DL_RT_NOVALIDDATA      //  NO VALID DATA
	ResultRtMemorylocked     Result = C.DL_RT_MEMORYLOCKED     //  MEMORY ALREADY LOCKED
	ResultRtInvalidmemorymap Result = C.DL_RT_INVALIDMEMORYMAP //  INVALID MEMORY MAP
	ResultRtInvalidRetain    Result = C.DL_RT_INVALID_RETAIN   //  INVALID MEMORY MAP
	ResultRtInternalError    Result = C.DL_RT_INTERNAL_ERROR   //  INTERNAL RT ERROR

	ResultSecNotoken             Result = C.DL_SEC_NOTOKEN             //  NO TOKEN FOUND
	ResultSecInvalidsession      Result = C.DL_SEC_INVALIDSESSION      //  TOKEN NOT VALID (SESSION NOT FOUND)
	ResultSecInvalidtokencontent Result = C.DL_SEC_INVALIDTOKENCONTENT //  TOKEN HAS WRONG CONTENT
	ResultSecUnauthorized        Result = C.DL_SEC_UNAUTHORIZED        //  UNAUTHORIZED
)

func (r Result) String() string {
	switch r {
	case ResultOk:
		return "DL_OK"
	case ResultFailed:
		return "DL_FAILED"
	case ResultInvalidAddress:
		return "DL_INVALID_ADDRESS"
	case ResultUnsupported:
		return "DL_UNSUPPORTED"
	case ResultOutOfMemory:
		return "DL_OUT_OF_MEMORY"
	case ResultLimitMin:
		return "DL_LIMIT_MIN"
	case ResultLimitMax:
		return "DL_LIMIT_MAX"
	case ResultTypeMismatch:
		return "DL_TYPE_MISMATCH"
	case ResultSizeMismatch:
		return "DL_SIZE_MISMATCH"
	case ResultInvalidFloatingpoint:
		return "DL_INVALID_FLOATINGPOINT"
	case ResultInvalidHandle:
		return "DL_INVALID_HANDLE"
	case ResultInvalidOperationMode:
		return "DL_INVALID_OPERATION_MODE"
	case ResultInvalidConfiguration:
		return "DL_INVALID_CONFIGURATION"
	case ResultInvalidValue:
		return "DL_INVALID_VALUE"
	case ResultSubmoduleFailure:
		return "DL_SUBMODULE_FAILURE"
	case ResultTimeout:
		return "DL_TIMEOUT"
	case ResultAlreadyExists:
		return "DL_ALREADY_EXISTS"
	case ResultCreationFailed:
		return "DL_CREATION_FAILED"
	case ResultVersionMismatch:
		return "DL_VERSION_MISMATCH"
	case ResultDeprecated:
		return "DL_DEPRECATED"
	case ResultPermissionDenied:
		return "DL_PERMISSION_DENIED"
	case ResultNotInitialized:
		return "DL_NOT_INITIALIZED"
	case ResultCommProtocolError:
		return "DL_COMM_PROTOCOL_ERROR"
	case ResultCommInvalidHeader:
		return "DL_COMM_INVALID_HEADER"
	case ResultClientNotConnected:
		return "DL_CLIENT_NOT_CONNECTED"
	case ResultRtNotopen:
		return "DL_RT_NOTOPEN"
	case ResultRtInvalidobject:
		return "DL_RT_INVALIDOBJECT"
	case ResultRtWrongrevison:
		return "DL_RT_WRONGREVISON"
	case ResultRtNovaliddata:
		return "DL_RT_NOVALIDDATA"
	case ResultRtMemorylocked:
		return "DL_RT_MEMORYLOCKED"
	case ResultRtInvalidmemorymap:
		return "DL_RT_INVALIDMEMORYMAP"
	case ResultRtInvalidRetain:
		return "DL_RT_INVALID_RETAIN"
	case ResultRtInternalError:
		return "DL_RT_INTERNAL_ERROR"
	case ResultSecNotoken:
		return "DL_SEC_NOTOKEN"
	case ResultSecInvalidsession:
		return "DL_SEC_INVALIDSESSION"
	case ResultSecInvalidtokencontent:
		return "DL_SEC_INVALIDTOKENCONTENT"
	case ResultSecUnauthorized:
		return "DL_SEC_UNAUTHORIZED"
	}
	return "Result_Unknown"
}
