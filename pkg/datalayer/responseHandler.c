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

void responseCallback(DLR_RESULT result, DLR_VARIANT data, void *userdata)
{
    responseCallbackGo(result, data, (unsigned long long)userdata);
}

// Wrapper for DLR_clientPingASync with responseKey instead of userdata
DLR_RESULT ClientPingASync(DLR_CLIENT client, unsigned long long responseKey)
{
    return DLR_clientPingASync(client, responseCallback, (void *)responseKey);
}

// Wrapper for DLR_clientCreateASync with responseKey instead of userdata
DLR_RESULT ClientCreateASync(DLR_CLIENT client, const char *address, DLR_VARIANT variant, const char *token, unsigned long long responseKey)
{
    return DLR_clientCreateASync(client, address, variant, token, responseCallback, (void *)responseKey);
}

// Wrapper for DLR_clientRemoveASync with responseKey instead of userdata
DLR_RESULT ClientRemoveASync(DLR_CLIENT client, const char *address, const char *token, unsigned long long responseKey)
{
    return DLR_clientRemoveASync(client, address, token, responseCallback, (void *)responseKey);
}

// Wrapper for DLR_clientRemoveASync with responseKey instead of userdata
DLR_RESULT ClientBrowseASync(DLR_CLIENT client, const char *address, const char *token, unsigned long long responseKey)
{
    return DLR_clientBrowseASync(client, address, token, responseCallback, (void *)responseKey);
}

// Wrapper for DLR_clientReadASync with responseKey instead of userdata
DLR_RESULT ClientReadASync(DLR_CLIENT client, const char *address, DLR_VARIANT variant, const char *token, unsigned long long responseKey)
{
    return DLR_clientReadASync(client, address, variant, token, responseCallback, (void *)responseKey);
}

// Wrapper for DLR_clientWriteASync with responseKey instead of userdata
DLR_RESULT ClientWriteASync(DLR_CLIENT client, const char *address, DLR_VARIANT variant, const char *token, unsigned long long responseKey)
{
    return DLR_clientReadASync(client, address, variant, token, responseCallback, (void *)responseKey);
}

// Wrapper for DLR_clientMetadataASync with responseKey instead of userdata
DLR_RESULT ClientMetadataASync(DLR_CLIENT client, const char *address, const char *token, unsigned long long responseKey)
{
    return DLR_clientMetadataASync(client, address, token, responseCallback, (void *)responseKey);
}
