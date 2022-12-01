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

#include <client.h>
#include <variant.h>

// responseBulkCallback wrapper for go callback
void responseBulkCallback(DLR_VEC_BULK_RESPONSE *response, void *userdata)
{
    responseBulkCallbackGo(response->response, (uint32_t)response->count, (unsigned long long)userdata);
}

// CreateBulkRequest wrapper for DLR_createBulkRequest with responseKey instead of userdata
unsigned long long CreateBulkRequest(uint32_t size)
{
    DLR_VEC_BULK_REQUEST *p = DLR_createBulkRequest(size);
    return (unsigned long long)p;
}

// DeleteBulkRequest wrapper for DLR_deleteBulkRequest with responseKey instead of userdata
void DeleteBulkRequest(unsigned long long p)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)p;
    DLR_deleteBulkRequest(&ptr);
}

// SetBulkRequestAddress helper function
void SetBulkRequestAddress(unsigned long long p, uint32_t index, const char *address)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)p;
    if (index >= ptr->count)
    {
        return;
    }
    ptr->request[index].address = address;
}

// SetBulkRequestData helper function
void SetBulkRequestData(unsigned long long p, uint32_t index, DLR_VARIANT data)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)p;
    if (index >= ptr->count)
    {
        return;
    }
    DLR_variantCopy(ptr->request[index].data, data);
}

// ClientReadBulkASync wrapper for DLR_clientReadBulkASync with responseKey instead of userdata
DLR_RESULT ClientReadBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)request;
    return DLR_clientReadBulkASync(client, ptr, token, responseBulkCallback, (void *)responseKey);
}

// ClientWriteBulkASync wrapper for ClientWriteBulkASync with responseKey instead of userdata
DLR_RESULT ClientWriteBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)request;
    return DLR_clientWriteBulkASync(client, ptr, token, responseBulkCallback, (void *)responseKey);
}

// ClientCreateBulkASync wrapper for DLR_clientCreateBulkASync with responseKey instead of userdata
DLR_RESULT ClientCreateBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)request;
    return DLR_clientCreateBulkASync(client, ptr, token, responseBulkCallback, (void *)responseKey);
}

// ClientDeleteBulkASync wrapper for DLR_clientDeleteBulkASync with responseKey instead of userdata
DLR_RESULT ClientDeleteBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)request;
    return DLR_clientDeleteBulkASync(client, ptr, token, responseBulkCallback, (void *)responseKey);
}

// ClientBrowseBulkASync wrapper for DLR_clientBrowseBulkASync with responseKey instead of userdata
DLR_RESULT ClientBrowseBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)request;
    return DLR_clientBrowseBulkASync(client, ptr, token, responseBulkCallback, (void *)responseKey);
}

// ClientMetadataBulkASync wrapper for DLR_clientMetadataBulkASync with responseKey instead of userdata
DLR_RESULT ClientMetadataBulkASync(DLR_CLIENT client, unsigned long long request, const char *token, unsigned long long responseKey)
{
    DLR_VEC_BULK_REQUEST *ptr = (DLR_VEC_BULK_REQUEST *)request;
    return DLR_clientMetadataBulkASync(client, ptr, token, responseBulkCallback, (void *)responseKey);
}
