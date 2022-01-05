/**
 * MIT License
 *
 * Copyright (c) 2021 Bosch Rexroth AG
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
#pragma once

#include <client.h>

#ifdef __cplusplus
extern "C"
{
#endif

  DLR_RESULT ClientPingASync(DLR_CLIENT client, unsigned long long responseKey);
  DLR_RESULT ClientCreateASync(DLR_CLIENT client, const char *address, DLR_VARIANT variant, const char *token, unsigned long long responseKey);
  DLR_RESULT ClientRemoveASync(DLR_CLIENT client, const char *address, const char *token, unsigned long long responseKey);
  DLR_RESULT ClientBrowseASync(DLR_CLIENT client, const char *address, const char *token, unsigned long long responseKey);
  DLR_RESULT ClientReadASync(DLR_CLIENT client, const char *address, DLR_VARIANT variant, const char *token, unsigned long long responseKey);
  DLR_RESULT ClientWriteASync(DLR_CLIENT client, const char *address, DLR_VARIANT variant, const char *token, unsigned long long responseKey);
  DLR_RESULT ClientMetadataASync(DLR_CLIENT client, const char *address, const char *token, unsigned long long responseKey);

  DLR_RESULT ClientCreateSubscriptionSync(DLR_CLIENT client, DLR_VARIANT ruleset, unsigned long long notifyKey, const char *token);

#ifdef __cplusplus
}
#endif
