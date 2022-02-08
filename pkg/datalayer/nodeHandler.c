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
#include "_cgo_export.h"

#include <stdio.h>
#include <stdbool.h>

#include <variant.h>
#include <provider_node.h>

typedef DLR_PROVIDER_NODE_CALLBACK TYPE_CB;
typedef DLR_PROVIDER_NODE_CALLBACKDATA TYPE_CBDATA;

void callCallbackC(TYPE_CB cb, TYPE_CBDATA cbdata, DLR_RESULT result, DLR_VARIANT data)
{
	cb(cbdata, result, data);
}

extern void nodeCallbackGo(void *userdata, char *address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata, int action);

void nodeCallbackOnCreate(void *userdata, char *address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata)
{
	nodeCallbackGo(userdata, address, data, cb, cbdata, NODE_ACTION_ON_CREATE);
}

void nodeCallbackOnRemove(void *userdata, char *address, TYPE_CB cb, TYPE_CBDATA cbdata)
{
	nodeCallbackGo(userdata, address, NULL, cb, cbdata, NODE_ACTION_ON_REMOVE);
}

void nodeCallbackOnBrowse(void *userdata, char *address, TYPE_CB cb, TYPE_CBDATA cbdata)
{
	nodeCallbackGo(userdata, address, NULL, cb, cbdata, NODE_ACTION_ON_BROWSE);
}

void nodeCallbackOnRead(void *userdata, char *address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata)
{
	nodeCallbackGo(userdata, address, data, cb, cbdata, NODE_ACTION_ON_READ);
}

void nodeCallbackOnWrite(void *userdata, char *address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata)
{
	nodeCallbackGo(userdata, address, data, cb, cbdata, NODE_ACTION_ON_WRITE);
}

void nodeCallbackOnMetadata(void *userdata, char *address, TYPE_CB cb, TYPE_CBDATA cbdata)
{
	nodeCallbackGo(userdata, address, NULL, cb, cbdata, NODE_ACTION_ON_METADATA);
}